package main

import (
	"account-server/cmd"
)

func main() {
	defer cmd.Clean()
	cmd.Start()
}
