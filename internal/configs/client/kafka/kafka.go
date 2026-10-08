package app_kafka_client

import (
	usecases "authService/internal/useCases"
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaClient struct {
	Config Config
	cl     *kgo.Client
}

func NewKafkaClient(config Config) (KafkaClient, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(config.addr),
	)

	if err != nil {
		return KafkaClient{}, fmt.Errorf("Failed to connet Kafka server: %w", err)
	}

	return KafkaClient{
		Config: config,
		cl:     client,
	}, nil
}

func (kc *KafkaClient) PublishUserRegistered(ctx context.Context, event usecases.RegisterUserEvent) error {
	record := &kgo.Record{Topic: "AuthUser", Value: []byte(event.Email)}

	if err := kc.cl.ProduceSync(ctx, record).FirstErr(); err != nil {
		return fmt.Errorf("record had a produce error while synchronously producing: %v\n", err)
	}

	return nil
}
