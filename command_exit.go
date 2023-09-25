package main

import "os"

func CallbackExit(cfg *config) error {
	os.Exit(0)
	return nil
}
