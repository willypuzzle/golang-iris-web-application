package log

import (
	"fmt"
	"strings"

	"iris.dom/app/config/utilities"
	"iris.dom/app/config/values/log"
	"iris.dom/app/config/variables"
)

func readEnvironment(key string, def log.LogLevel) log.LogLevel {
	v := utilities.Getenv(key, def.String())
	if v == "" {
		return def
	}

	env := log.LogLevel(strings.ToLower(v))

	for _, value := range log.LEVELS {
		if env == value {
			return env
		}
	}

	panic(fmt.Sprintf("Unknown Log Level (%s)", env.String()))
}

// GetEnvironment returns the environment of the instance
func GetLogLevel() log.LogLevel {
	l := variables.LOG_LEVEL.String()
	return readEnvironment(l, log.LogLevel(variables.DEFAULTS[l]))
}
