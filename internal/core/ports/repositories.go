package ports

/**
	Centralizarmos a criação das interfaces ports do control agent
	@author: Jorge Luis
	@version: 0.0.1
	@Documentation: https://github.com/aws/aws-sdk-go-v2/tree/main/service/cloudcontrol
**/

/**
	Intterface do port de conexão com o AWS SDK
**/
type CloudControlSDK interface {
	Save(domain.cloudcontrol) error
	Delete(domain.cloudcontrol)
}
