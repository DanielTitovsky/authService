package app_kafka_client

import (
	usecases "authService/internal/useCases"
	"context"
	"encoding/json"
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

	payload, err := json.Marshal(event)

	if err != nil {
		return fmt.Errorf("failed to parse event: %w", err)
	}

	record := &kgo.Record{Topic: "AuthUser", Value: payload}

	if err := kc.cl.ProduceSync(ctx, record).FirstErr(); err != nil {
		return fmt.Errorf("record had a produce error while synchronously producing: %v\n", err)
	}

	return nil
}
