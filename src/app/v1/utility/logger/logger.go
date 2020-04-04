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
		Driver: os.Getenv("LOGGER_DRIVER"),
		Mongo:  mongodb.MongoDBHandler(),
	}
}

// LoggerInterface ..
type LoggerInterface interface {
	Save(log *entity.LoggerRequest) (*entity.LoggerEventHistory, error)
	Get(uuid, collection string) (*entity.LoggerEventHistory, error)
}

// Save ...
func (logging *Logger) Save(log *entity.LoggerRequest) (*entity.LoggerEventHistory, error) {
	switch logging.Driver {
	case "mongo":
		return logging.loggerMongoSave(log)
	case "s3":
		fmt.Println("S3")
	default:
		fmt.Println("Default")
	}

	return nil, errors.New("Driver Load Failed")
}

// Get ...
func (logging *Logger) Get(uuid, collection string) (*entity.LoggerEventHistory, error) {
	switch logging.Driver {
	case "mongo":
		return logging.loggerGetMongoUUID(uuid, collection)
	case "s3":
		fmt.Println("S3")
	default:
		fmt.Println("Default")
	}
	return nil, errors.New("Driver Load Failed")
}
