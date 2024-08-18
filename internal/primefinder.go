package internal

func PrimeFinder(done <-chan bool, randIntStream <-chan int) <-chan int {
	isPrime := func(num int) bool {
		for i := num - 1; i > 1; i-- {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)

	go func() {
		defer close(primes)

		for {
			select {
			case <-done:
				return
			case num := <-randIntStream:
				if isPrime(num) {
					primes <- num
				}
			}
		}
	}()

	return primes
}
