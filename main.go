package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: committing [br|cm|status|pr|fe]")
		fmt.Println("  br     - create new branch")
		fmt.Println("  cm     - create commit")
		fmt.Println("  status - show branch status")
		fmt.Println("  pr     - create or open pull request")
		fmt.Println("  fe     - fetch and merge main into current branch")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "br":
		if err := runBranchFlow(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "cm":
		if err := runCommitFlow(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "status":
		if err := runStatusFlow(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "pr":
		if err := runPRFlow(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "fe":
		if err := runFetchFlow(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: br, cm, status, pr, fe")
		os.Exit(1)
	}
}
