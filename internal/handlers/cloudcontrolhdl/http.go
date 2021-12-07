package cloudcontrolhdl

import (
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	cloudcontrolService cloudcontrolsrv.service
}

func NewHTTPHandler(service cloudcontrolsrv.service) *HTTPHandler {
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
