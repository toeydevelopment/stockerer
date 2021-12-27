package destination

import (
	"context"

	"github.com/nats-io/nats.go"
)

type Nats struct {
	nc *nats.Conn
}

func (n *Nats) Publish(ctx context.Context, topic, msg string) error {

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	return n.nc.Publish(topic, []byte(msg))
}

func (n *Nats) Close(ctx context.Context) error {

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	n.nc.Close()
	return nil
}
