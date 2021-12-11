package cloudcontrolhdl

import (
	"github.com/jfonseca85/arquitetura-onion/internal/core/domain"
)

type BodyCreate struct {
	ClientToken   string `json:"clientToken"`
	DesiredState  string `json:"DesiredState"`
	RoleArn       string `json:"roleArn"`
	TypeName      string `json:"typeName"`
	TypeVersionId string `json:"typeVersionId"`
}

type ResponseCreate *domain.ProgressEvent

type RequestCreate *domain.CloudControlModel

func BuildResponseCreate(progressEvent *domain.ProgressEvent) ResponseCreate {
	return ResponseCreate(progressEvent)
}
