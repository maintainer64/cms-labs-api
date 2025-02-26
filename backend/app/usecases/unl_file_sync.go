package usecases

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/a10869/api-modules/backend/app/models"
	"gitlab.com/a10869/api-modules/backend/app/queries"
	"gitlab.com/a10869/api-modules/backend/app/usecases/response"
	"gitlab.com/a10869/api-modules/shared/utils"
)

type UNLFileSyncUC struct {
	UNLFileQueries *queries.UNLFileQueries
	*zerolog.Logger
}

type UNLFileSyncInputDTO struct {
	Branch     string `json:"branch"`
	Repository string `json:"repository"`
}

type UNLFileSyncOutputDTO struct {
	Count int `json:"count"`
}

type UNLFileSyncResponse = response.Response[UNLFileSyncOutputDTO]

func (u *UNLFileSyncUC) Execute(dto UNLFileSyncInputDTO) (UNLFileSyncOutputDTO, error) {
	output := UNLFileSyncOutputDTO{}
	dir, err := os.MkdirTemp("", "git-repo")
	if err != nil {
		return UNLFileSyncOutputDTO{}, utils.FiberValidationException{
			Status:    fiber.StatusInternalServerError,
			Exception: errors.New("not created tmp directory"),
		}
	}
	defer os.RemoveAll(dir)
	branch := fmt.Sprintf("refs/heads/%s", dto.Branch)
	u.Info().Msg(fmt.Sprintf("git clone %s from %s", dto.Repository, branch))

	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL:           dto.Repository,
		ReferenceName: plumbing.ReferenceName(branch),
		SingleBranch:  true,
		Depth:         1,
	})
	if err != nil {
		return UNLFileSyncOutputDTO{}, utils.FiberValidationException{
			Status:    fiber.StatusInternalServerError,
			Exception: fmt.Errorf("not clone repository. %+v", err),
		}
	}

	syncedFilesID := uuid.New().String()

	log.Info().Msg(fmt.Sprintf("unl file start, synced files: %s", syncedFilesID))

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		// Получение относительного пути
		filePath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		// Пропускаем скрытые папки и директории
		if strings.HasPrefix(filePath, ".") {
			return nil
		}
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}
		// Чтение содержимого файла
		content, err := os.ReadFile(filepath.Clean(path))
		if err != nil {
			return err
		}
		// Получение расширения файла
		ext := strings.ReplaceAll(filepath.Ext(path), ".", "")
		// Обновление файла в базе данных
		var unlFile models.UNLFile
		unlFile.Path = filePath
		unlFile.Content = content
		unlFile.Type = ext
		unlFile.SyncedId = syncedFilesID
		err = u.UNLFileQueries.Upsert(&unlFile)
		output.Count++
		return err
	})

	if err != nil {
		return output, err
	}

	// Шаг 2: Помечаем несинхронизированные файлы как soft_deleted
	err = u.UNLFileQueries.SoftDeleteNotSyncedId(syncedFilesID)
	if err != nil {
		return output, err
	}
	log.Info().Msg(fmt.Sprintf("unl file sync successfully, synced files: %s", syncedFilesID))

	return output, err
}
