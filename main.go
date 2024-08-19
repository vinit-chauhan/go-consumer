package main

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/vinit-chauhan/go-consumer/internal/consumer"
	"github.com/vinit-chauhan/go-consumer/internal/generator"
	"github.com/vinit-chauhan/go-consumer/internal/tasks"
	"github.com/vinit-chauhan/go-consumer/internal/types"
	"github.com/vinit-chauhan/go-consumer/utils"

	"golang.org/x/exp/rand"
)

const MAX_RANDOM int = 5000000

var (
	ctx           context.Context
	config        types.Config
	consumeAmount int
)

func init() {
	rand.Seed(uint64(time.Now().Unix()))

	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())

	// NOTE: FOR TESTING PURPOSE ONLY
	go func(cencel context.CancelFunc) {
		t := time.NewTicker(30 * 100 * time.Millisecond)
		select {
		case <-t.C:
			cancel()
		}
	}(cancel)

	config = types.DefaultConfig()
	config.WithGoRoutineCount(runtime.NumCPU())

	consumeAmount = 10
}

func main() {
	rNumFetcher := func() int {
		return rand.Intn(MAX_RANDOM)
	}

	numberStream := generator.Run(ctx, rNumFetcher)
	getNumberStream := func() <-chan int {
		return tasks.PrimeFinder(ctx, numberStream)
	}

	primeFinderChannels := utils.FanOut(ctx, getNumberStream, config.GoRoutineCount)

	fannedInStream := utils.FanIn(ctx, primeFinderChannels...)

	for random := range consumer.Run(ctx, fannedInStream, consumeAmount) {
		fmt.Println(random)
	}
}
