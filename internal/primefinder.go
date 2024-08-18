package internal

import "context"

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
