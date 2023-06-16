package host

import (
	"fmt"
	"strconv"

	"iris.dom/app/config/utilities"
	"iris.dom/app/config/variables"
)

// GetAppPort this function returns the port of the app set by environment variable or by default value
func GetAppPort() uint64 {
	l := variables.APP_PORT.String()
	port, err := strconv.ParseUint(utilities.Getenv(l, variables.DEFAULTS[l]), 10, 32)
	if err != nil {
		panic(fmt.Sprintf("Wrong environment variable %s (%v)", l, err))
	}
	return port
}
