package cloudcontrolhdl

import (
	"github.com/gin-gonic/gin"
	"github.com/jfonseca85/controlplaneagent/internal/core/service/cloudcontrolsrv"
)

type HTTPHandler struct {
	cloudcontrolService cloudcontrolsrv.Service
}

func NewHTTPHandler(service cloudcontrolsrv.Service) *HTTPHandler {
	return &HTTPHandler{
		cloudcontrolService: service,
	}
}

func (hdl *HTTPHandler) Create(c *gin.Context) {
	body := BodyCreate{}
	c.BindJSON(&body)

	model := BuildRequestCreate(body)

	resource, err := hdl.cloudcontrolService.Create(model)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, BuildResponseCreate(resource))

}
