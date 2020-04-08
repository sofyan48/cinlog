package logger

import (
	"errors"
	"fmt"
	"os"

	"github.com/sofyan48/cinlog/src/app/v1/api/logger/entity"
	"github.com/sofyan48/cinlog/src/app/v1/utility/mongodb"
)

// Logger ..
type Logger struct {
	Driver string
	Mongo  mongodb.MongoDBInterface
}

// LoggerHandler habndler logger
func LoggerHandler() *Logger {
	return &Logger{
		Driver: os.Getenv("LOGGER_STORAGE"),
		Mongo:  mongodb.MongoDBHandler(),
	}
}

// LoggerInterface ..
type LoggerInterface interface {
	Save(origin string, log *entity.LoggerRequest) (*entity.LoggerEventHistory, error)
	Get(uuid, collection string) (*entity.LoggerEventHistory, error)
	All(action string) (interface{}, error)
}

// Save ...
func (logging *Logger) Save(origin string, log *entity.LoggerRequest) (*entity.LoggerEventHistory, error) {
	switch logging.Driver {
	case "mongo":
		return logging.loggerMongoSave(origin, log)
	case "s3":
		fmt.Println("S3")
	default:
		fmt.Println("Default")
	}

	return nil, errors.New("Driver Load Failed")
}

// Get ...
func (logging *Logger) Get(uuid, action string) (*entity.LoggerEventHistory, error) {
	switch logging.Driver {
	case "mongo":
		return logging.loggerGetMongoUUID(uuid, action)
	case "s3":
		fmt.Println("S3")
	default:
		fmt.Println("Default")
	}
	return nil, errors.New("Driver Load Failed")
}

// All ...
func (logging *Logger) All(action string) (interface{}, error) {
	switch logging.Driver {
	case "mongo":
		return logging.loggerGetAll(action)
	case "s3":
		fmt.Println("S3")
	default:
		fmt.Println("Default")
	}
	return nil, errors.New("Driver Load Failed")
}
