package generator

import (
	"context"

	"github.com/vinit-chauhan/go-consumer/utils"
)

func Run[T any](ctx context.Context, fn func() T) <-chan T {

	utils.Logger.Debug("Starting Generator")

	stream := make(chan T)
	go func() {
		defer close(stream)
		for {
			select {
			case <-ctx.Done():
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}
