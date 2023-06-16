package log

const (
	DISABLE LogLevel = "disable"
	FATAL   LogLevel = "fatal"
	ERROR   LogLevel = "error"
	WARN    LogLevel = "warn"
	INFO    LogLevel = "info"
	DEBUG   LogLevel = "debug"
)

var LEVELS = [6]LogLevel{
	DISABLE,
	FATAL,
	ERROR,
	WARN,
	INFO,
	DEBUG,
}

type LogLevel string

func (e LogLevel) String() string {
	return string(e)
}
