package main

import (
	"fmt"
	"time"

	"github.com/vinit-chauhan/go-consumer/internal"

	"golang.org/x/exp/rand"
)

const MAX_RANDOM int = 5000000

func main() {
	done := make(chan bool)
	defer close(done)
	rand.Seed(uint64(time.Now().Unix()))

	rNumFetcher := func() int {
		return rand.Intn(MAX_RANDOM)
	}

	numberStream := internal.Generator(done, rNumFetcher)

	primeStream := internal.PrimeFinder(done, numberStream)

	for random := range internal.Consume(done, primeStream, 10) {
		fmt.Println(random)
	}

}
