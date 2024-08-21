package tasks

import (
	"context"

	"github.com/vinit-chauhan/go-consumer/utils"
)

func PrimeFinder[T int | int64](ctx context.Context, stream <-chan T) <-chan T {

	isPrime := func(num T) bool {
		for i := num - 1; i > 1; i-- {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan T)

	go func() {
		defer close(primes)

		for {
			select {
			case <-ctx.Done():
				utils.Logger.Debug("[PrimeFinder] closing goroutine for prime finder")
				return
			case num := <-stream:
				if isPrime(num) {
					primes <- num
				}
			}
		}
	}()

	return primes
}
