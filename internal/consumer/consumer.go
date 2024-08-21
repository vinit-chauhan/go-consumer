package consumer

import (
	"context"

	"github.com/vinit-chauhan/go-consumer/utils"
)

func Run[T any](ctx context.Context, stream <-chan T, count int) <-chan T {
	consume := make(chan T)
	utils.Logger.Debug("[Run] starting consumer")

	go func() {
		defer close(consume)
		for i := 0; i < count; i++ {
			select {
			case <-ctx.Done():
				utils.Logger.Debug("[Run] closing consumer")
				return
			case consume <- <-stream:
			}
		}
	}()

	return consume
}
