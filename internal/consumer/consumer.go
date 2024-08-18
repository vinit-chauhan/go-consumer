package consumer

import "context"

func Run[T any](ctx context.Context, stream <-chan T, count int) <-chan T {
	consume := make(chan T)

	go func() {
		defer close(consume)
		for i := 0; i < count; i++ {
			select {
			case <-ctx.Done():
				return
			case consume <- <-stream:
			}
		}
	}()

	return consume
}
