package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbletea"
)

func runCommitFlow() error {
	// Add all changes first
	fmt.Println("Adding all changes...")
	addCmd := exec.Command("git", "add", "-A")
	addOutput, err := addCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to add changes: %v\nOutput: %s", err, string(addOutput))
	}

	// Check if there are any staged changes to commit
	statusCmd := exec.Command("git", "diff", "--cached", "--quiet")
	err = statusCmd.Run()
	if err == nil {
		// No staged changes (git diff --cached --quiet returns 0 if no changes)
		fmt.Println("No changes to commit")
		return nil
	}

	// Get current branch and check if it has a type prefix
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchOutput, err := branchCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get current branch: %v", err)
	}
	currentBranch := strings.TrimSpace(string(branchOutput))

	// Check if branch name follows the type/description format
	var detectedType string
	if parts := strings.SplitN(currentBranch, "/", 2); len(parts) == 2 {
		branchType := parts[0]
		// Check if the type is valid
		for _, validType := range commitTypes {
			if branchType == validType {
				detectedType = branchType
				break
			}
		}
	}

	var flow flowModel
	if detectedType != "" {
		fmt.Printf("Auto-detected commit type '%s' from branch name\n", detectedType)
		flow = newFlowWithType("commit", detectedType)
	} else {
		flow = newFlow("commit")
	}
	
	program := tea.NewProgram(flow)
	
	finalModel, err := program.Run()
	if err != nil {
		return err
	}

	result := finalModel.(flowModel)
	if result.cancelled {
		fmt.Println("Commit cancelled")
		return nil
	}

	commitMessage := fmt.Sprintf("%s: %s", result.selectedType, result.description)

	cmd := exec.Command("git", "commit", "-m", commitMessage)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to commit: %v\nOutput: %s", err, string(output))
	}

	fmt.Printf("✅ Committed with message: %s\n", commitMessage)

	// Check if branch exists on remote and push if needed
	if err := checkAndPushBranch(); err != nil {
		fmt.Printf("Warning: %v\n", err)
	}

	return nil
}

func checkAndPushBranch() error {
	// Get current branch name
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchOutput, err := branchCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get current branch: %v", err)
	}
	currentBranch := strings.TrimSpace(string(branchOutput))

	// Check if remote branch exists using gh
	ghCmd := exec.Command("gh", "api", fmt.Sprintf("repos/:owner/:repo/branches/%s", currentBranch))
	err = ghCmd.Run()
	
	if err != nil {
		// Branch doesn't exist on remote, push it
		fmt.Println("Branch not found on remote, pushing...")
		
		pushCmd := exec.Command("git", "push", "--set-upstream", "origin", currentBranch)
		pushOutput, err := pushCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to push branch: %v\nOutput: %s", err, string(pushOutput))
		}
		
		fmt.Printf("✅ Pushed branch %s to remote\n", currentBranch)
	}
	
	return nil
}
