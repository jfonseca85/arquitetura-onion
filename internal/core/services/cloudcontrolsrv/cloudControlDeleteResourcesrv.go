//Deleta resources usando o aws sdk go v2 - service cloudcontrol
package cloudcontrolsrv

import (
	"context"
	"log"

	"github.com/jfonseca85/controlplaneagent/internal/core/domain"
	"github.com/jfonseca85/controlplaneagent/internal/types"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

func (srv *Service) Delete(model domain.CloudControlModel) (*domain.ProgressEvent, error) {
	//Metodo que destinado para Deletar os recursos usando o cloudcontrol
	log.Printf("Chamando o cloudControlCreateResourcesrv.Delete>>> ")
	output, err := srv.controlsdk.DeleteResource(context.TODO(), toDeleteResourceInput(model))
	progressEvent := toDeleteResourceProgressEvent(output)
	if err != nil {
		return nil, err
	}
	return progressEvent, nil
}

func toDeleteResourceInput(model domain.CloudControlModel) *cloudcontrol.DeleteResourceInput {
	//Realiza a conversao do domain.CloudControlModel para cloudcontrol.DeleteResourceInput
	log.Printf("Chamando o cloudControlCreateResourcesrv.toDeleteResourceInput>>> ")
	result := cloudcontrol.DeleteResourceInput{
		TypeName:      &(model.TypeName),
		Identifier:    &(model.Identifier),
		ClientToken:   &(model.ClientToken),
		RoleArn:       &(model.RoleArn),
		TypeVersionId: &(model.TypeVersionId),
	}
	return &result
}

func toDeleteResourceProgressEvent(model *cloudcontrol.DeleteResourceOutput) *domain.ProgressEvent {
	//Realiza a conversao do DeleteResourceOutput para domain.ProgressEvent
	log.Printf("Chamando o cloudControlCreateResourcesrv.toDeleteResourceProgressEvent>>> ")
	result := domain.ProgressEvent{
		ErrorCode:       types.HandlerErrorCode(model.ProgressEvent.ErrorCode),
		EventTime:       model.ProgressEvent.EventTime,
		Identifier:      model.ProgressEvent.Identifier,
		Operation:       types.Operation(model.ProgressEvent.Operation),
		OperationStatus: types.OperationStatus(model.ProgressEvent.Operation),
		RequestToken:    model.ProgressEvent.RequestToken,
		ResourceModel:   model.ProgressEvent.ResourceModel,
		RetryAfter:      model.ProgressEvent.RetryAfter,
		StatusMessage:   model.ProgressEvent.StatusMessage,
		TypeName:        model.ProgressEvent.TypeName,
	}
	return &result
}
