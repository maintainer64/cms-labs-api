package di

import (
	"gitlab.com/a10869/api-modules/shared/logs"
)

type DIContainer struct {
	ZeroLogConf *logs.ZeroLoggerConf
}

func (di *DIContainer) Close() {
}

func NewDIContainer(zeroLogConf *logs.ZeroLoggerConf) (*DIContainer, error) {
	di := &DIContainer{
		ZeroLogConf: zeroLogConf,
	}
	return di, nil
}
