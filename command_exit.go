package main

import "os"

func callbackExit(cdf *config) error {
	os.Exit(0)
	return nil
}