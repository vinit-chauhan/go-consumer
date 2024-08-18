package internal

func Consume[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	consume := make(chan T)

	go func() {
		defer close(consume)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case consume <- <-stream:
			}
		}
	}()

	return consume
}
