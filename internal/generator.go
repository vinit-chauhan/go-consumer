package internal

import "context"

func Generator[T any](ctx context.Context, fn func() T) <-chan T {
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
