package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func runStatusFlow() error {
	// Get current branch name
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchOutput, err := branchCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get current branch: %v", err)
	}
	currentBranch := strings.TrimSpace(string(branchOutput))

	fmt.Printf("📌 Current branch: %s\n", currentBranch)

	// Check if branch exists on remote
	remoteExists := checkRemoteBranchExists(currentBranch)
	if remoteExists {
		fmt.Printf("🌐 Remote: exists\n")
	} else {
		fmt.Printf("🌐 Remote: not found\n")
	}

	// Check for existing PR
	prURL := getPRURL(currentBranch)
	if prURL != "" {
		fmt.Printf("🔗 Pull request: %s\n", prURL)
	} else {
		fmt.Printf("🔗 Pull request: none\n")
	}

	return nil
}

func checkRemoteBranchExists(branch string) bool {
	cmd := exec.Command("gh", "api", fmt.Sprintf("repos/:owner/:repo/branches/%s", branch))
	err := cmd.Run()
	return err == nil
}

func getPRURL(branch string) string {
	// Use gh to find PR for the current branch
	cmd := exec.Command("gh", "pr", "view", "--json", "url", "-q", ".url")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}
