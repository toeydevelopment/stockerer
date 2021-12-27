package destination

import "context"

type Destinationer interface {
	Publish(ctx context.Context, topic, msg string) error
	Close(ctx context.Context) error
}
