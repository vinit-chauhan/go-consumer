package types

type LogLevel int

const (
	QUITE LogLevel = iota
	ERROR
	INFO
	VERBOSE
	DEBUG
)

func (l LogLevel) String() string {
	return [...]string{"primefinder"}[l]
}

func (l LogLevel) EnumIndex() int {
	return int(l)
}
