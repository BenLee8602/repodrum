package main

import (
	"fmt"
	"os"

	"github.com/BenLee8602/repodrum/internal/cli"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must provide a subcommand")
		os.Exit(1)
	}

	subcommand := os.Args[1]
	args := os.Args[2:]
	switch subcommand {
	case "install":
		cli.Install(args)
		break
	case "uninstall":
		cli.Uninstall()
		break
	case "add":
		cli.Add(args)
		break
	case "remove":
		cli.Remove(args)
		break
	default:
		fmt.Println("Unknown subcommand: ", subcommand)
		os.Exit(1)
	}
}
