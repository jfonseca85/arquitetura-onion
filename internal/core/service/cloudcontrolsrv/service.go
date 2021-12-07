package cloudcontrolsrv

import (
	"github.com/jfonseca85/controlplaneagent/internal/core/ports/cloudcontrol/sdk"
)

type service struct {
	cloudControlSDK sdk.CloudControlSDK
}

func New(gamesRepository sdk.CloudControlSDK) *service {
	return &service{
		cloudControlSDK: cloudControlSDK,
	}
}
