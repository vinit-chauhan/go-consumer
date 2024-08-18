package utils

import (
	"context"
	"sync"
)

func FanIn[T int | int64](ctx context.Context, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-ctx.Done():
				return
			case fannedInStream <- i:
			}
		}
	}

	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}

func FanOut[T int | int64](ctx context.Context, getStream func() <-chan T, count int) []<-chan T {
	consumerChans := make([]<-chan T, count)
	for i := 0; i < count; i++ {
		consumerChans[i] = getStream()
	}
	return consumerChans
}
