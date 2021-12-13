//Esse adaptador de driver deve ser capaz de transformar uma solicitação http em uma chamada de serviço.
//Ter todos os componentes separados uns dos outros nos dá a vantagem de implementá-los e
//Testá-los isoladamente ou podemos facilmente paralelizar o trabalho com a ajuda de outros membros da equipe.
package cloudcontrolhdl

import (
	"github.com/jfonseca85/arquitetura-onion/internal/core/services/cloudcontrolservice"
)

type HTTPHandler struct {
	//dependencyInjection
	//TO DO: Passar a interface do port e diminuir o acoplamento
	cloudcontrolService cloudcontrolservice.Service
}

func NewHTTPHandler(service cloudcontrolservice.Service) *HTTPHandler {
	return &HTTPHandler{
		cloudcontrolService: service,
	}
}
