package cloudcontrolsrv

import (
	"github.com/jfonseca85/controlplaneagent/internal/core/domain/cloudcontrol"
	"github.com/jfonseca85/controlplaneagent/internal/core/ports/cloudcontrolapi"
)

type service struct {
	cloudControlSDK cloudcontrolapi.SDK
}

func New(cloudControlSDK cloudcontrolapi.SDK) *service {
	return &service{
		cloudControlSDK: cloudControlSDK,
	}
}

func (srv *service) Create(cloudcontrol.Model) (*cloudcontrol.ProgressEvent, error) {
	progressEvent, err := srv.cloudControlSDK.Save(cloudcontrol.Model{})
	if err != nil {
		return nil, err
	}

	return progressEvent, err
}
