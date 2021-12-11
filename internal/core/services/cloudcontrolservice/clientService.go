package cloudcontrolservice

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

const ServiceID = "CloudControl"

const ServiceAPIVersion = "2021-09-30"

// O Service fornece o API Client para fazer chamadas de operações para o AWS Cloud Control
// API.
type Service struct {
	controlsdk *cloudcontrol.Client
}

// New retorna um cliente inicializado com base nas opções funcionais.
func New(controlsdk *cloudcontrol.Client) Service {
	return Service{
		controlsdk: controlsdk,
	}
}
