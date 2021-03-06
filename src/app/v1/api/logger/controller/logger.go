package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/cinlog/src/app/v1/api/logger/entity"
	"github.com/sofyan48/cinlog/src/app/v1/api/logger/service"
	"github.com/sofyan48/cinlog/src/app/v1/utility/rest"
)

// LoggerController ...
type LoggerController struct {
	Service service.LoggerServiceInterface
}

// LoggerControllerHandler ...
func LoggerControllerHandler() *LoggerController {
	return &LoggerController{
		Service: service.LoggerServiceHandler(),
	}
}

const origin = "REST"

// LoggerControllerInterface ...
type LoggerControllerInterface interface {
	CreateLogger(context *gin.Context)
	GetLogger(context *gin.Context)
	AllLogger(context *gin.Context)
}

// CreateLogger ...
func (ctrl *LoggerController) CreateLogger(context *gin.Context) {
	payload := &entity.LoggerRequest{}
	context.ShouldBind(payload)
	result, err := ctrl.Service.CreateLoggerService(origin, payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}

// GetLogger ...
func (ctrl *LoggerController) GetLogger(context *gin.Context) {
	payload := &entity.GetLoggerRequest{}
	context.ShouldBind(payload)
	result, err := ctrl.Service.GetLoggerByUUID(payload.UUID, payload.Action)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}

// AllLogger ...
func (ctrl *LoggerController) AllLogger(context *gin.Context) {
	payload := &entity.GetAllLoggerRequest{}
	context.ShouldBind(payload)
	result, err := ctrl.Service.GetLoggerAll(payload.Action)
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, result)
	return
}
