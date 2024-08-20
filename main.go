package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/vinit-chauhan/go-consumer/internal/consumer"
	"github.com/vinit-chauhan/go-consumer/internal/generator"
	"github.com/vinit-chauhan/go-consumer/internal/tasks"
	"github.com/vinit-chauhan/go-consumer/internal/types"
	"github.com/vinit-chauhan/go-consumer/utils"
	"go.uber.org/zap"

	"golang.org/x/exp/rand"
)

const MAX_RANDOM int = 5000000

var (
	ctx           context.Context
	config        types.Config
	Log           *zap.Logger
	consumeAmount int
)

func init() {
	config = types.DefaultConfig()
	config.WithGoRoutineCount(runtime.NumCPU())

	var err error
	Log, err = utils.Logger(config)
	if err != nil {
		panic(err)
	}

	Log.Debug("Setting random seed")
	rand.Seed(uint64(time.Now().Unix()))

	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())

	// NOTE: FOR TESTING PURPOSE ONLY
	go func(cencel context.CancelFunc) {
		t := time.NewTicker(3 * 100 * time.Millisecond)
		select {
		case <-t.C:
			Log.Debug("Service timed out")
			cancel()
		}
	}(cancel)

	consumeAmount = 10
}

func main() {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)

	ctx = context.WithValue(ctx, "logger", Log)

	Log.Info("Starting service")

	defer func() {
		Log.Debug("Received termination signal from os")
		cancel()
	}()

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
