package usecases

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/queries"
	"gitlab.com/a10869/api-modules/shared/cms_client"
)

type PnetServerPingUC struct {
	LabSessionQuery *queries.LabSessionQuery
	CMSClient       *cms_client.CMSClient
	*zerolog.Logger
}

// ClearOldExternalUnlFiles - очистка старых файлов, у которых нет активной сессии
func (u *PnetServerPingUC) ClearOldExternalUnlFiles(activeFilePaths []string) error {
	// Создаем быстрый поиска разрешенных файлов
	allowedFiles := make(map[string]bool)
	for _, f := range activeFilePaths {
		// Файл состоит из пути .../external/filename
		// Оставляем в allowedFiles только filename
		filePathDB := strings.TrimPrefix(f, pathOptUNetLab)
		prefix := path.Join("/", pathExternal)
		if !strings.HasPrefix(filePathDB, prefix) {
			continue
		}
		allowedFiles[strings.TrimPrefix(filePathDB, prefix)] = true
	}

	files, err := os.ReadDir(path.Join(pathOptUNetLab, pathExternal))
	if err != nil {
		log.Warn().Msg(fmt.Sprintf("ClearOldExternalUnlFiles: failed to read directory: %v", err))
		return fmt.Errorf("failed to read directory: %v", err)
	}
	// 4. Удаляем старые UNL файлы
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if allowedFiles[file.Name()] {
			continue
		}
		if err := os.Remove(path.Join(pathOptUNetLab, pathExternal, file.Name())); err != nil {
			log.Warn().Msg(fmt.Sprintf("ClearOldExternalUnlFiles: failed to remove file: %v", err))
			continue
		}
	}
	return nil
}

// Execute - функция отправки активных сессий лабораторных работ
func (u *PnetServerPingUC) Execute() error {
	labs, err := u.LabSessionQuery.GetRunningLabs()
	if err != nil {
		log.Warn().Msg(fmt.Sprintf("Get running labs failed: %v", err))
		return err
	}
	var attempts []cms_client.PnetServerPingAttemptDTO
	var activeFilePaths []string
	for _, lab := range labs {
		externalUserId, _ := strconv.ParseUint(lab.Name, 10, 32)
		attempt := cms_client.PnetServerPingAttemptDTO{
			AttemptID: lab.LabSessionLID,
			UserEmail: lab.Email,
			UserID:    uint(externalUserId),
		}
		attempts = append(attempts, attempt)
		activeFilePaths = append(activeFilePaths, lab.LabSessionPath)
	}
	response, err := u.CMSClient.PnetServerPing(
		&cms_client.PNETServerPingInputDTO{
			Attempts: attempts,
		},
	)
	if response != nil {
		u.Logger.Info().Msg(fmt.Sprintf("Ping server response count: %d", response.Count))
	}
	_ = u.ClearOldExternalUnlFiles(activeFilePaths)
	return err
}
