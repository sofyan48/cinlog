package service

import (
	"fmt"
	"time"

	"github.com/sofyan48/cinlog/src/app/v1/api/logger/entity"
	"github.com/sofyan48/cinlog/src/app/v1/utility/mongodb"
)

// LoggerService ...
type LoggerService struct {
	mongo mongodb.MongoDBInterface
}

// LoggerServiceHandler ...
func LoggerServiceHandler() *LoggerService {
	return &LoggerService{
		mongo: mongodb.MongoDBHandler(),
	}
}

// LoggerServiceInterface ...
type LoggerServiceInterface interface {
	CreateLoggerService(log *entity.LoggerRequest) (*entity.LoggerEventHistory, error)
	GetLoggerByUUID(uuid, collection string) (*entity.LoggerEventHistory, error)
}

// CreateLoggerService ...
func (service *LoggerService) CreateLoggerService(log *entity.LoggerRequest) (*entity.LoggerEventHistory, error) {
	now := time.Now()
	findLogger := &entity.LoggerSearch{}
	historyLogger := entity.LoggerHistory{}
	historyLoggerSlice := []entity.LoggerHistory{}
	eventLogger := &entity.LoggerEventHistory{}
	findLogger.UUID = log.UUID
	result, err := service.mongo.FindOne(log.Action, findLogger)
	err = result.Decode(eventLogger)
	if err != nil {
		eventLogger.Offset = 0

		eventLogger.Action = log.Action
		eventLogger.CreatedAt = &now
		eventLogger.UpdateAt = &now
		eventLogger.UUID = log.UUID
		eventLogger.Status = log.Status
		historyLogger.Data = log.Data
		historyLoggerSlice = append(historyLoggerSlice, historyLogger)
		eventLogger.History = historyLoggerSlice

		go service.mongo.InsertOne(log.Action, eventLogger)
		return eventLogger, nil
	}
	eventLogger.Offset = eventLogger.Offset + 1
	eventLogger.UpdateAt = &now
	eventLogger.Status = log.Status
	historyLogger.Data = log.Data
	eventLogger.History = append(eventLogger.History, historyLogger)
	fmt.Println(eventLogger)
	go service.mongo.InsertOne(log.Action, eventLogger)
	return eventLogger, nil
}

// GetLoggerByUUID ...
func (service *LoggerService) GetLoggerByUUID(uuid, collection string) (*entity.LoggerEventHistory, error) {
	findLogger := &entity.LoggerSearch{}
	findLogger.UUID = uuid
	eventLogger := &entity.LoggerEventHistory{}
	result, err := service.mongo.GetOne(collection, findLogger)
	err = result.Decode(eventLogger)
	if err != nil {
		return nil, err
	}
	return eventLogger, nil
}
