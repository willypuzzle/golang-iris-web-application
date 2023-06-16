package utilities

import "os"

func Getenv(key string, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}
