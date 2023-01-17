package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/labstack/echo/v4"
)

type KafkaConfig struct {
	Brockers []string
	Topics   []string
	Config   *sarama.Config
	Producer sarama.SyncProducer
}

func (KafkaConfig) Init() *KafkaConfig {
	result := &KafkaConfig{Brockers: []string{"kafka:9092"}, Topics: []string{"sarama"}}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	result.Config = config

	producer, err := sarama.NewSyncProducer(result.Brockers, config)
	if err != nil {
		panic("failed to connect kafka - " + err.Error())
	}
	result.Producer = producer

	return result
}

func (config KafkaConfig) KafkaMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		context.Set("kafka", config)
		return next(context)
	}
}
