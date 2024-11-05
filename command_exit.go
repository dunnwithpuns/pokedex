package main

import "os"

func commandExit(cfg *config, commandParameter string) error {
	os.Exit(0)
	return nil
}
