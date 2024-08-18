package internal

import "context"

func Consume[T any](ctx context.Context, stream <-chan T, n int) <-chan T {
	consume := make(chan T)

	go func() {
		defer close(consume)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case consume <- <-stream:
			}
		}
	}()

	return consume
}
