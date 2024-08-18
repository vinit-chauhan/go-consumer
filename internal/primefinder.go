package internal

func PrimeFinder[T any, K int | int64](done <-chan T, stream <-chan K) <-chan K {
	isPrime := func(num K) bool {
		for i := num - 1; i > 1; i-- {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan K)

	go func() {
		defer close(primes)

		for {
			select {
			case <-done:
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
