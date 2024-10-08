package types

type TaskType int

const (
	NOOP TaskType = iota
	PrimeFinderTask
)

func (tt TaskType) String() string {
	return [...]string{"primefinder"}[tt]
}

func (tt TaskType) EnumIndex() int {
	return int(tt)
}
