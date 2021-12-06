package domain

/**
	Domain API Cloud Control para operações de API
	@author: Jorge Luis
	@version: 0.0.1
	@Documentation: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/what-is-cloudcontrolapi.html
**/
type CloudControl struct {
	//Um identificador exclusivo para garantir a idempotência da solicitação de recurso.
	ClientToken string `json:"clientToken"`

	//Formato de dados estruturados que representa o estado desejado do recurso
	DesiredState string `json:"desiredState"`

	//O Amazon Resource Name (ARN) da função AWS Identity and Access Management (IAM) para a API Cloud Control usar ao executar esta operação de recurso.
	RoleArn string `json:"roleArn"`

	//O nome do tipo de recurso.
	TypeName string `json:"typeName"`
	//Para tipos de recursos privados, a versão do tipo a ser usada nesta operação de recurso.
	TypeVersionId string `json:"typeVersionId"`

	//O identificador do recurso.
	Identifier string `json:"identifier"`
}
