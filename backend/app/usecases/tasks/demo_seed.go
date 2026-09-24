package tasks

import (
	"github.com/maintainer64/cms-labs-api/backend/app/models"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// DemoSeedUC creates deterministic records used by the local full-stack demo.
type DemoSeedUC struct { DB *gorm.DB; Logger *zerolog.Logger }
const DemoAttemptID = "550e8400-e29b-41d4-a716-446655440000"

func (u *DemoSeedUC) Execute() error {
	var user models.User
	if err := u.DB.Where("email = ?", "admin@admin.com").First(&user).Error; err != nil { return err }
	var server models.Server
	if err := u.DB.Where("client_id = ?", "demo-kubernetes").First(&server).Error; err != nil {
		server = models.Server{ServerBase: models.ServerBase{Name:"Demo Kubernetes", Url:"http://clabgate:8080", Type:models.ServerTypeKubernetes, IsActive:true}, ServerSecret:models.ServerSecret{ClientID:"demo-kubernetes"}}
		if err := u.DB.Create(&server).Error; err != nil { return err }
	}
	var route models.LTIRouting
	if err := u.DB.Where("name = ?", "Simple Task Demo").First(&route).Error; err != nil {
		route = models.LTIRouting{LTIRoutingBase:models.LTIRoutingBase{Name:"Simple Task Demo"}, LTIRoutingSecret:models.LTIRoutingSecret{LabsType:"default", LabsPath:"task", TestPath:"sdn_lab_5", ServerID:server.ID, IsDefault:true}}
		if err := u.DB.Create(&route).Error; err != nil { return err }
	}
	var attempt models.LTIAttempt
	err := u.DB.Where("attempt_id = ?", DemoAttemptID).First(&attempt).Error
	if err == gorm.ErrRecordNotFound {
		attempt = models.LTIAttempt{LTIAttemptBase:models.LTIAttemptBase{AttemptID:DemoAttemptID, Status:models.AttemptStatusPending, UserID:user.ID, LTIRoutingID:route.ID}, LTIAttemptSecret:models.LTIAttemptSecret{}}
		attempt.ServerID = &server.ID
		if err := u.DB.Create(&attempt).Error; err != nil { return err }
	} else if err != nil { return err }
	u.Logger.Info().Str("attempt_id", DemoAttemptID).Msg("demo records ready")
	return nil
}
