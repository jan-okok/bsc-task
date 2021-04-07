package main

import (
	"github/bsc-task/cli/command"
	"github/bsc-task/log"
)

func main() {
	log.Init()
	command.Execute()
}
