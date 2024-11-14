package main

import (
	"os"
	"strconv"
)

func port(defaultPort int) int {
	portStr := os.Getenv("PORT")
	var port int

	if portStr != "" {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			port = defaultPort // Use default if conversion fails
		}
	} else {
		port = defaultPort
	}
	return port
}
