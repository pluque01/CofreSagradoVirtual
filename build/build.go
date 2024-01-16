package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var build = goyek.Define(goyek.Task{
	Name:  "build",
	Usage: "go build",
	Action: func(a *goyek.A) {
		cmd.Exec(a, "go build ./cmd/...", cmd.Stdout(a.Output()))
	},
})
