package repository

import (
	"fmt"
	"m2ex-otp-service/internal/util"

	"gopkg.in/Shopify/sarama.v1"
)

type PluginRepository interface {
	EventProducer(topic string, entity any)
}

type pluginRepository struct {
}

func NewPluginRepository() PluginRepository {
	return pluginRepository{}
}

func (r pluginRepository) EventProducer(topic string, event any) {

	config, _ := util.LoadConfig()
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V0_10_2_0
	cfg.Producer.Return.Successes = true

	url := fmt.Sprintf("%v:%v", config.Kafka.Host, config.Kafka.Port)
	kafkaServer := []string{
		url,
	}
	producer, err := sarama.NewSyncProducer(kafkaServer, cfg)
	if err != nil {
		util.Logger.Error(err.Error())
	}

	value := util.CompressToJsonBytes(event)

	message := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}

	_, _, err = producer.SendMessage(&message)
	if err != nil {
		util.Logger.Error(err.Error())
	}

	producer.Close()
}
