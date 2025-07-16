package usecases

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/a10869/api-modules/pnetlabaddon/app/models"

	"gitlab.com/a10869/api-modules/pnetlabaddon/app/queries"
	"gitlab.com/a10869/api-modules/shared/cms_client"
)

const (
	pathOptUNetLab = "/opt/unetlab/labs"
	pathExternal   = "external"
)

type LabCreateInputDTO struct {
	Extra   string `json:"extra"`
	UserPod int    `json:"user_pod"`
}

type LabCreateUC struct {
	UserQueries     *queries.UserQueries
	LabSessionQuery *queries.LabSessionQuery
	CMSClient       *cms_client.CMSClient
	*zerolog.Logger
}

func (u *LabCreateUC) CreateOrUpdateLab(
	userPod int,
	unlFilePathServer string,
	attemptId string,
) error {
	user, err := u.UserQueries.Get(userPod)
	if err != nil {
		return err
	}
	lab, _ := u.LabSessionQuery.GetByAttemptId(attemptId)
	if lab.LabSessionID != 0 {
		lab.AddJoinedUser(userPod)
		u.LabSessionQuery.DB.Save(&lab)
		user.LabSession = &lab.LabSessionID
		u.UserQueries.DB.Save(&user)
		log.Info().Msg(
			fmt.Sprintf(
				"UpdateLab by unlFilePathServer: %s, attemptId:%s and userPod: %d",
				unlFilePathServer,
				attemptId,
				userPod,
			),
		)
		return nil
	}
	lab = models.LabSession{
		LabSessionLID:    attemptId,
		LabSessionPod:    userPod,
		LabSessionJoined: fmt.Sprintf("%d", userPod),
		LabSessionPath:   unlFilePathServer,
	}
	u.LabSessionQuery.DB.Save(&lab)
	user.LabSession = &lab.LabSessionID
	u.UserQueries.DB.Save(&user)
	log.Info().Msg(
		fmt.Sprintf(
			"CreateLab by unlFilePathServer: %s, attemptId:%s and userPod: %d",
			unlFilePathServer,
			attemptId,
			userPod,
		),
	)
	return nil
}

func (u *LabCreateUC) DownloadLab(
	curlRequestID string,
	attemptId string,
) (string, error) {
	fileName := filepath.Clean(fmt.Sprintf("%s.unl", attemptId))
	u64, err := strconv.ParseUint(curlRequestID, 10, 32)
	if err != nil {
		u.Logger.Info().Msg(
			fmt.Sprintf(
				"LabCreateUC: not parsed curlRequestID: %s",
				curlRequestID,
			),
		)
		return "", err
	}
	fileContent, err := u.CMSClient.DownloadUNLFile(
		uint(u64),
		map[string]string{
			"attemptId": attemptId,
		},
	)
	if err != nil {
		return "", err
	}
	if len(fileContent) <= 0 {
		u.Logger.Info().Msg(
			fmt.Sprintf(
				"LabCreateUC: file content downloaded is empty by curlRequestID: %s, attemptId: %s",
				curlRequestID,
				attemptId,
			),
		)
		return "", nil
	}
	if err := os.MkdirAll(filepath.Join(pathOptUNetLab, pathExternal), 0750); err != nil {
		u.Logger.Info().Msg(
			fmt.Sprintf(
				"LabCreateUC: failed to create directory curlRequestID: %s, attemptId: %s",
				curlRequestID,
				attemptId,
			),
		)
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	// Записываем файл (создаем или перезаписываем существующий)
	filePath := filepath.Join(pathOptUNetLab, pathExternal, fileName)
	if !strings.HasPrefix(filePath, filepath.Join(pathOptUNetLab, pathExternal)) {
		return "", fmt.Errorf("invalid path")
	}
	file, err := os.Create(filepath.Clean(filePath))
	if err != nil {
		u.Logger.Info().Msg(
			fmt.Sprintf(
				"LabCreateUC: failed to create file curlRequestID: %s, attemptId: %s",
				curlRequestID,
				attemptId,
			),
		)
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()
	if _, err := file.Write(fileContent); err != nil {
		u.Logger.Info().Msg(
			fmt.Sprintf(
				"LabCreateUC: failed to write file curlRequestID: %s, attemptId: %s",
				curlRequestID,
				attemptId,
			),
		)
		return "", fmt.Errorf("failed to write file: %v", err)
	}
	u.Logger.Info().Msg(
		fmt.Sprintf(
			"LabCreateUC: created file: %s by curlRequestID: %s, attemptId: %s",
			filepath.Join("/", pathExternal, fileName),
			curlRequestID,
			attemptId,
		),
	)
	return filepath.Join("/", pathExternal, fileName), nil
}

func (u *LabCreateUC) Execute(dto LabCreateInputDTO) error {
	extra := cms_client.UnmarshalSSOTokenPublicExtraParams(dto.Extra)
	if extra.AttemptID == "" {
		u.Logger.Info().Msg("LabCreate attemptID is null")
		return nil
	}
	if extra.PNETLabsType == cms_client.PNETLabsTypeDefault && extra.PNETLabsPath != "" {
		u.Logger.Info().Msg(fmt.Sprintf("LabCreate by default path: %s", extra.PNETLabsPath))
		return u.CreateOrUpdateLab(dto.UserPod, extra.PNETLabsPath, extra.AttemptID)
	}
	if extra.PNETLabsType == cms_client.PNETLabsTypeCurl && extra.PNETLabsPath != "" {
		u.Logger.Info().Msg(fmt.Sprintf("LabCreate by default path: %s", extra.PNETLabsPath))
		pnetLabUnlFile, err := u.DownloadLab(extra.PNETLabsPath, extra.AttemptID)
		if err != nil {
			return err
		}
		return u.CreateOrUpdateLab(dto.UserPod, pnetLabUnlFile, extra.AttemptID)
	}
	return nil
}
