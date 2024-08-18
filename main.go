package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/vinit-chauhan/go-consumer/internal"
	"github.com/vinit-chauhan/go-consumer/utils"

	"golang.org/x/exp/rand"
)

const MAX_RANDOM int = 5000000

func main() {
	done := make(chan bool)
	defer close(done)

	CPUCount := runtime.NumCPU()
	rand.Seed(uint64(time.Now().Unix()))

	rNumFetcher := func() int {
		return rand.Intn(MAX_RANDOM)
	}

	numberStream := internal.Generator(done, rNumFetcher)
	getNumberStream := func() <-chan int {
		return internal.PrimeFinder(done, numberStream)
	}

	primeFinderChannels := utils.FanOut(done, getNumberStream, CPUCount)

	fannedInStream := utils.FanIn(done, primeFinderChannels...)

	for random := range internal.Consume(done, fannedInStream, 10) {
		fmt.Println(random)
	}
}
