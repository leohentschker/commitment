package main

import (
	"fmt"
	"os/exec"
)

func runFetchFlow() error {
	fmt.Println("🔄 Fetching latest main...")
	
	// Fetch latest main branch
	fetchCmd := exec.Command("git", "fetch", "origin", "main")
	output, err := fetchCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to fetch main: %v\nOutput: %s", err, string(output))
	}

	fmt.Println("📥 Merging main into current branch...")
	
	// Merge origin/main into current branch
	mergeCmd := exec.Command("git", "merge", "origin/main")
	output, err = mergeCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to merge main: %v\nOutput: %s", err, string(output))
	}

	fmt.Println("✅ Successfully merged main into current branch")
	return nil
}
