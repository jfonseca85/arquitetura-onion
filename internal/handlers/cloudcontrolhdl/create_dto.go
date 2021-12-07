package cloudcontrolhdl

import (
	"github.com/jfonseca85/controlplaneagent/internal/core/domain/cloudcontrol"
)

type BodyCreate struct {
	ClientToken   string `json:"clientToken"`
	DesiredState  string `json:"DesiredState"`
	RoleArn       string `json:"roleArn"`
	TypeName      string `json:"typeName"`
	TypeVersionId string `json:"typeVersionId"`
}

type ResponseCreate *cloudcontrol.ProgressEvent

type RequestCreate *cloudcontrol.Model

func BuildResponseCreate(progressEvent *cloudcontrol.ProgressEvent) ResponseCreate {
	return ResponseCreate(progressEvent)
}

func BuildRequestCreate(dto BodyCreate) cloudcontrol.Model {
	model := cloudcontrol.Model{
		ClientToken:   &dto.ClientToken,
		DesiredState:  &dto.DesiredState,
		RoleArn:       &dto.RoleArn,
		TypeName:      &dto.TypeName,
		TypeVersionId: &dto.TypeVersionId,
	}
	return model
}
