package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var test = goyek.Define(goyek.Task{
	Name:  "test",
	Usage: "go test -v ./internal/...",
	Action: func(a *goyek.A) {
		cmd.Exec(a, "go test -v ./internal/...")
	},
})
