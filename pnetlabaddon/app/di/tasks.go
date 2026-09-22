package di

import (
	"github.com/maintainer64/cms-labs-api/shared/logs"
)

type PnetServerPingTask struct {
}

func (u *PnetServerPingTask) Execute() error {
	container, err := NewDIContainer(logs.NewZeroLoggerConf(nil).SetName("PnetServerPingTask"))
	if err != nil {
		return err
	}
	defer container.Close()
	return container.PnetServerPingUC().Execute()
}
