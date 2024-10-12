package main

import (
	"math/rand"
	"os"
	"ssl_tool/cmd/commands"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	if err := commands.NewSslToolCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
