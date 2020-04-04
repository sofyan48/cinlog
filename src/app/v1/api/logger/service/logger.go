package service

import (
	"github.com/sofyan48/cinlog/src/app/v1/api/logger/entity"
	"github.com/sofyan48/cinlog/src/app/v1/utility/logger"
)

// LoggerService ...
type LoggerService struct {
	Logger logger.LoggerInterface
}

// LoggerServiceHandler ...
func LoggerServiceHandler() *LoggerService {
	return &LoggerService{
		Logger: logger.LoggerHandler(),
	}
}

// LoggerServiceInterface ...
type LoggerServiceInterface interface {
	CreateLoggerService(log *entity.LoggerRequest) (*entity.LoggerEventHistory, error)
	GetLoggerByUUID(uuid, action string) (*entity.LoggerEventHistory, error)
	GetLoggerAll(action string) (interface{}, error)
}

// CreateLoggerService ...
func (service *LoggerService) CreateLoggerService(log *entity.LoggerRequest) (*entity.LoggerEventHistory, error) {
	return service.Logger.Save(log)
}

// GetLoggerByUUID ...
func (service *LoggerService) GetLoggerByUUID(uuid, action string) (*entity.LoggerEventHistory, error) {
	return service.Logger.Get(uuid, action)
}

// GetLoggerAll ...
func (service *LoggerService) GetLoggerAll(action string) (interface{}, error) {
	return service.Logger.All(action)
}
