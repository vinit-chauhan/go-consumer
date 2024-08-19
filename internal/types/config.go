package types

import "runtime"

type Config struct {
	GoRoutineCount int      // number of go routine
	LogLevel       LogLevel // log level for the service
	Task           TaskType // which task to run
}

func DefaultConfig() Config {
	return Config{
		GoRoutineCount: runtime.NumCPU(),
		LogLevel:       INFO,
		Task:           NOOP,
	}
}

func (c *Config) WithGoRoutineCount(count int) *Config {
	c.GoRoutineCount = count
	return c
}

func (c *Config) WithLogLevel(logLevel LogLevel) *Config {
	c.LogLevel = logLevel
	return c
}

func (c *Config) WithTaskType(task TaskType) *Config {
	c.Task = task
	return c
}
