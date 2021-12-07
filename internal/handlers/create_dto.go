package gamehdl

import "github.com/jfonseca85/controlplaneagent/internal/core/domain/cloudcontrol"

type BodyCreate struct {
	Name  string `json:"name"`
	Size  uint   `json:"size"`
	Bombs uint   `json:"bombs"`
}

type ResponseCreate cloudcontrol.Model

func BuildResponseCreate(model cloudcontrol.Model) ResponseCreate {
	return ResponseCreate(model)
}
