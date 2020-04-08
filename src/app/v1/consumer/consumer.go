package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/sofyan48/cinlog/src/app/v1/utility/kafka"
	"github.com/sofyan48/cinlog/src/app/v1/utility/logger"
)

// V1ConsumerEvents ...
type V1ConsumerEvents struct {
	Kafka  kafka.KafkaLibraryInterface
	Logger logger.LoggerInterface
}

// V1ConsumerEventsHandler ...
func V1ConsumerEventsHandler() *V1ConsumerEvents {
	return &V1ConsumerEvents{
		Kafka:  kafka.KafkaLibraryHandler(),
		Logger: logger.LoggerHandler(),
	}
}

// V1ConsumerEventsInterface ...
type V1ConsumerEventsInterface interface {
	Consume(topics []string, signals chan os.Signal)
}

// Consume ...
func (consumer *V1ConsumerEvents) Consume(topics []string, signals chan os.Signal) {
	// StateFullData := consumer.Kafka.GetStateFull()
	chanMessage := make(chan *sarama.ConsumerMessage, 256)
	csm, err := consumer.Kafka.InitConsumer()
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	for _, topic := range topics {
		partitionList, err := csm.Partitions(topic)
		if err != nil {
			log.Println("Unable to get partition got error ", err)
			continue
		}
		for _, partition := range partitionList {
			go consumeMessage(csm, topic, partition, chanMessage)
		}
	}
	log.Println("Event is Started....")

ConsumerLoop:
	for {
		select {
		case msg := <-chanMessage:
			eventData := &kafka.StateFullFormat{}
			json.Unmarshal(msg.Value, eventData)
			switch eventData.Function {
			case "save":
				consumer.eventSave(eventData)
			case "delete":
				consumer.eventDelete(eventData.Payload.UUID)
			default:
				log.Println("DEFAULT")

			}
		case sig := <-signals:
			if sig == os.Interrupt {
				break ConsumerLoop
			}
		}
	}
}

func consumeMessage(consumer sarama.Consumer, topic string, partition int32, c chan *sarama.ConsumerMessage) {
	msg, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		log.Println("Unable to consume partition got error ", partition, err)
		return
	}
	defer func() {
		if err := msg.Close(); err != nil {
			log.Println("Unable to close partition : ", partition, err)
		}
	}()
	for {
		msg := <-msg.Messages()
		c <- msg
	}

}

func (consumer *V1ConsumerEvents) eventSave(data *kafka.StateFullFormat) {
	const eventOrigin = "BROKER"
	go consumer.Logger.Save(eventOrigin, data.Payload)
}

func (consumer *V1ConsumerEvents) eventDelete(uuid string) {
	fmt.Println("DELETED EVENT")
}
