package main

import "os"

func callbackExit(cdf *config, args ...string) error {
	os.Exit(0)
	return nil
}