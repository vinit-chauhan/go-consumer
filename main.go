package main

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/vinit-chauhan/go-consumer/internal"
	"github.com/vinit-chauhan/go-consumer/utils"

	"golang.org/x/exp/rand"
)

const MAX_RANDOM int = 5000000

func main() {
	ctx := context.Background()
	defer ctx.Done()

	CPUCount := runtime.NumCPU()
	rand.Seed(uint64(time.Now().Unix()))

	rNumFetcher := func() int {
		return rand.Intn(MAX_RANDOM)
	}

	numberStream := internal.Generator(ctx, rNumFetcher)
	getNumberStream := func() <-chan int {
		return internal.PrimeFinder(ctx, numberStream)
	}

	primeFinderChannels := utils.FanOut(ctx, getNumberStream, CPUCount)

	fannedInStream := utils.FanIn(ctx, primeFinderChannels...)

	for random := range internal.Consume(ctx, fannedInStream, 10) {
		fmt.Println(random)
	}
}
