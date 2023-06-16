package environment

import (
	"strings"

	"iris.dom/app/config/utilities"
	"iris.dom/app/config/values/environment"
	"iris.dom/app/config/variables"
)

func readEnvironment(key string, def environment.EnvironmentValue) environment.EnvironmentValue {
	v := utilities.Getenv(key, def.String())
	if v == "" {
		return def
	}

	env := environment.EnvironmentValue(strings.ToLower(v))
	switch env {
	case environment.PROD, environment.DEV:
	default:
		panic("unexpected environment " + v)
	}

	return env
}

// GetEnvironment returns the environment of the instance
func GetEnvironment() environment.EnvironmentValue {
	l := variables.APP_ENVIRONMENT.String()
	return readEnvironment(l, environment.EnvironmentValue(variables.DEFAULTS[l]))
}
