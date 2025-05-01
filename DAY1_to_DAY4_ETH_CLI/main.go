package main

import (
	"ethcli/cmd"
	"os"
)

func main() {
	switch os.Args[1] {
	case "queryblock":
		cmd.QueryBlock()
	case "querytx":
		cmd.QueryAddress()
	case "sendtx":
		cmd.SendEth()
	default:
		cmd.Execute()
	}
}
