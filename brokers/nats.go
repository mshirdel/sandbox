package brokers

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func Run(msg string) error {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return fmt.Errorf("connecting to nats: %w", err)
	}
	defer nc.Drain()
	
	return nc.Publish("foo", []byte(msg))
}
