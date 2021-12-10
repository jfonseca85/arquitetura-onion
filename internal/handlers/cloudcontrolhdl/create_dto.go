package cloudcontrolhdl

import (
	"github.com/jfonseca85/controlplaneagent/internal/core/domain/cloudcontrolmdl"
)

type BodyCreate struct {
	ClientToken   string `json:"clientToken"`
	DesiredState  string `json:"DesiredState"`
	RoleArn       string `json:"roleArn"`
	TypeName      string `json:"typeName"`
	TypeVersionId string `json:"typeVersionId"`
}

type ResponseCreate *cloudcontrolmdl.ProgressEvent

type RequestCreate *cloudcontrolmdl.Model

func BuildResponseCreate(progressEvent *cloudcontrolmdl.ProgressEvent) ResponseCreate {
	return ResponseCreate(progressEvent)
}

func BuildRequestCreate(dto BodyCreate) cloudcontrolmdl.Model {
	model := cloudcontrolmdl.Model{
		//ClientToken:   dto.ClientToken,
		DesiredState: dto.DesiredState,
		//RoleArn:       dto.RoleArn,
		TypeName: dto.TypeName,
		//TypeVersionId: dto.TypeVersionId,
	}
	return model
}
