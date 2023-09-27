package main

import "os"

func CallbackExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
