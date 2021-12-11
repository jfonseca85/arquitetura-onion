package builder

import "github.com/jfonseca85/arquitetura-onion/internal/core/domain"

type newCloudControlModel struct {
	model domain.CloudControlModel
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

func (c *newCloudControlModel) Agora() domain.CloudControlModel {
	return c.model
}
