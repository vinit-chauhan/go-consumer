package utils

import (
	"sync"
)

func FanIn[T any, K int | int64](done <-chan T, channels ...<-chan K) <-chan K {
	var wg sync.WaitGroup
	fannedInStream := make(chan K)

	transfer := func(c <-chan K) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
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

func FanOut[T any, K int | int64](done <-chan T, getStream func() <-chan K, count int) []<-chan K {
	consumerChans := make([]<-chan K, count)
	for i := 0; i < count; i++ {
		consumerChans[i] = getStream()
	}
	return consumerChans
}
