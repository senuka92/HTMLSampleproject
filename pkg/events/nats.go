package events

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	js nats.JetStreamContext
}

func NewPublisher(url string) (*Publisher, error) {
	if url == "" {
		return nil, nil
	}
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("connect nats: %w", err)
	}
	js, err := conn.JetStream()
	if err != nil {
		return nil, fmt.Errorf("jetstream: %w", err)
	}
	return &Publisher{js: js}, nil
}

func (p *Publisher) Publish(ctx context.Context, subject string, payload interface{}) error {
	if p == nil || p.js == nil {
		return nil
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal event: %w", err)
	}
	_, err = p.js.PublishAsync(subject, data)
	return err
}
