package variables

import (
	"iris.dom/app/config/values/environment"
	"iris.dom/app/config/values/log"
)

const (
	APP_ENVIRONMENT variable = "APP_ENVIRONMENT"
	APP_PORT        variable = "APP_PORT"
	LOG_LEVEL       variable = "LOG_LEVEL"
)

var DEFAULTS = map[string]string{
	APP_ENVIRONMENT.String(): environment.PROD.String(),
	APP_PORT.String():        "8000",
	LOG_LEVEL.String():       log.DISABLE.String(),
}

type variable string

func (v variable) String() string {
	return string(v)
}
