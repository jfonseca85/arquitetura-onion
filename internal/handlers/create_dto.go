package gamehdl

import "github.com/jfonseca85/controlplaneagent/internal/core/domain/cloudcontrolmdl"

type BodyCreate struct {
	Name  string `json:"name"`
	Size  uint   `json:"size"`
	Bombs uint   `json:"bombs"`
}

type ResponseCreate cloudcontrolmdl.Model

func BuildResponseCreate(model cloudcontrolmdl.Model) ResponseCreate {
	return ResponseCreate(model)
}
