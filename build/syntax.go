package main

import (
	"io"
	"strings"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var lint = goyek.Define(goyek.Task{
	Name:  "lint",
	Usage: "gofmt -l",
	Action: func(a *goyek.A) {
		sb := &strings.Builder{}
		out := io.MultiWriter(a.Output(), sb)
		cmd.Exec(a, "gofmt -l ./internal", cmd.Stdout(out))
		if sb.Len() > 0 {
			a.Error("gofmt -l returned output")
		}
	},
})
