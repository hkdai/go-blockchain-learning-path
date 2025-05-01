package main

import (
	"os"
	"ethcli/cmd"
)
func main() {
	switch os.Args[1] {
	case "queryblock":
		cmd.QueryBlock()
	default:
		cmd.Execute()
	}
}
