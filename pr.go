package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func runPRFlow() error {
	// Get current branch name
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchOutput, err := branchCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get current branch: %v", err)
	}
	currentBranch := strings.TrimSpace(string(branchOutput))

	// Check if we're on main branch
	if currentBranch == "main" || currentBranch == "master" {
		return fmt.Errorf("cannot create PR from main/master branch")
	}

	// Check for existing PR
	existingPR := getPRURL(currentBranch)
	if existingPR != "" {
		fmt.Printf("📖 Opening existing PR: %s\n", existingPR)
		openCmd := exec.Command("gh", "pr", "view", "--web")
		return openCmd.Run()
	}

	// Ensure branch is pushed to remote before creating PR
	if !checkRemoteBranchExists(currentBranch) {
		fmt.Println("🚀 Pushing branch to remote...")
		pushCmd := exec.Command("git", "push", "--set-upstream", "origin", currentBranch)
		pushOutput, err := pushCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to push branch: %v\nOutput: %s", err, string(pushOutput))
		}
	}

	// Create new PR against main
	fmt.Println("📝 Creating new pull request...")
	createCmd := exec.Command("gh", "pr", "create", "--base", "main", "--web")
	err = createCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to create PR: %v", err)
	}

	fmt.Println("✅ Pull request created and opened in browser")
	return nil
}
