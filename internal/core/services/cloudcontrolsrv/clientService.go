package cloudcontrolsrv

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

type Service struct {
	controlsdk *cloudcontrol.Client
}

func New(controlsdk *cloudcontrol.Client) Service {
	return Service{
		controlsdk: controlsdk,
	}
}
