package builder

import "github.com/jfonseca85/controlplaneagent/internal/core/domain/cloudcontrolmdl"

type newCloudControlModel struct {
	model cloudcontrolmdl.Model
}

func UmCloudControlModel() *newCloudControlModel {
	return &newCloudControlModel{}
}

func (c *newCloudControlModel) ComDesiredState(desiredState string) *newCloudControlModel {
	c.model.DesiredState = desiredState
	return c
}

func (c *newCloudControlModel) ComTypeName(typeName string) *newCloudControlModel {
	c.model.TypeName = typeName
	return c
}

func (c *newCloudControlModel) ComClientToken(clientToken string) *newCloudControlModel {
	c.model.ClientToken = clientToken
	return c
}

func (c *newCloudControlModel) ComRoleArn(roleArn string) *newCloudControlModel {
	c.model.RoleArn = roleArn
	return c
}

func (c *newCloudControlModel) ComTypeVersionId(typeVersionId string) *newCloudControlModel {
	c.model.TypeVersionId = typeVersionId
	return c
}

func (c *newCloudControlModel) Agora() cloudcontrolmdl.Model {
	return c.model
}
