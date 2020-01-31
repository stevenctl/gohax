package main

import (
	"flag"
	"fmt"
	"github.com/stevenctl/gohax/cmd"
)

func main() {
	command := flag.String("command", "none", "the command to run")
	host := flag.String("host", "", "hostname to scan ports on")
	flag.Parse()

	if command == nil {
		fmt.Println("please specify a command")
		return
	}

	switch *command {
	case "portscan":
		if host == nil {
			fmt.Println("--host must be specified")
			return
		}
		cmd.PortScan(*host)
	default:
		fmt.Printf("invalid command: %s\n", *command)
	}
}