package kafka

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/sofyan48/cinlog/src/app/v1/api/logger/entity"
)

// KafkaLibrary ...
type KafkaLibrary struct{}

// StateFullFormat ...
type StateFullFormat struct {
	Payload  *entity.LoggerRequest `json:"payload"`
	Function string                `json:"__function"`
}

// KafkaLibraryHandler ...
func KafkaLibraryHandler() *KafkaLibrary {
	return &KafkaLibrary{}
}

// KafkaLibraryInterface ...
type KafkaLibraryInterface interface {
	GetStateFull() *StateFullFormat
	InitConsumer() (sarama.Consumer, error)
}

// GetStateFull ...
func (kafka *KafkaLibrary) GetStateFull() *StateFullFormat {
	return &StateFullFormat{}
}

// Init ...
func (kafka *KafkaLibrary) init(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}
