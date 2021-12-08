package cloudcontrolsrv

import (
	"github.com/jfonseca85/controlplaneagent/internal/core/domain/cloudcontrolmdl"
	"github.com/jfonseca85/controlplaneagent/internal/core/ports/cloudcontrolapi"

	"github.com/aws/aws-sdk-go-v2/config"
)

type service struct {
	cloudControlSDK cloudcontrolapi.SDK
	cfg             config.Config
}

func New() *service {
	return &service{}
}

func (srv *service) Create(model cloudcontrolmdl.Model) (*cloudcontrolmdl.ProgressEvent, error) {

	instanceAWS := singleton.getInstance()
	progressEvent, err := srv.cloudControlSDK.Save(instanceAWS.ctx, model)
	if err != nil {
		return nil, err
	}

	return progressEvent, err
}
