package main

import (
	"log"
	"os"
	"os/user"
)

func antivm() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	// yes i know blocking usernames is retarded, ill fix it when i feel like it
	blockedusers := map[string]bool{
		"admin":              true,
		"administrator":      true,
		"WDAGUtilityAccount": true,
	}

	if blockedusers[user.Username] {
		os.Exit(0)
	}
}
