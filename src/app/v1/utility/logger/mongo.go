package logger

import (
	"log"
	"time"

	"github.com/sofyan48/cinlog/src/app/v1/api/logger/entity"
)

// LoggerMongoSave ..
func (logging *Logger) loggerMongoSave(log *entity.LoggerRequest) (*entity.LoggerEventHistory, error) {
	now := time.Now()
	findLogger := &entity.LoggerSearch{}
	historyLogger := entity.LoggerHistory{}
	historyLoggerSlice := []entity.LoggerHistory{}
	eventLogger := &entity.LoggerEventHistory{}
	findLogger.UUID = log.UUID
	result, err := logging.Mongo.FindOne(log.Action, findLogger)
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
		go logging.Mongo.InsertOne(log.Action, eventLogger)
		return eventLogger, nil
	}
	eventLogger.Offset = eventLogger.Offset + 1
	eventLogger.UpdateAt = &now
	eventLogger.Status = log.Status
	historyLogger.Data = log.Data
	eventLogger.History = append(eventLogger.History, historyLogger)
	go logging.Mongo.InsertOne(log.Action, eventLogger)
	return eventLogger, nil
}

// LoggerGetMongoUUID ...
func (logging *Logger) loggerGetMongoUUID(uuid, collection string) (*entity.LoggerEventHistory, error) {
	findLogger := &entity.LoggerSearch{}
	findLogger.UUID = uuid
	eventLogger := &entity.LoggerEventHistory{}
	result, err := logging.Mongo.GetOne(collection, findLogger)
	err = result.Decode(eventLogger)
	if err != nil {
		return nil, err
	}
	return eventLogger, nil
}

func (logging *Logger) loggerGetAll(action string) ([]entity.LoggerEventHistory, error) {

	cur, ctx, err := logging.Mongo.Find(action)
	if err != nil {
		return nil, err
	}
	result := []entity.LoggerEventHistory{}
	for cur.Next(ctx) {
		data := entity.LoggerEventHistory{}
		err = cur.Decode(&data)
		if err != nil {
			log.Println("Error on Decoding the document", err)
		}
		result = append(result, data)
	}
	return result, nil
}
