package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbletea"
)

func runBranchFlow() error {
	flow := newFlow("branch")
	program := tea.NewProgram(flow)
	
	finalModel, err := program.Run()
	if err != nil {
		return err
	}

	result := finalModel.(flowModel)
	if result.cancelled {
		fmt.Println("Branch creation cancelled")
		return nil
	}

	branchName := fmt.Sprintf("%s/%s", result.selectedType, strings.ReplaceAll(result.description, " ", "-"))
	branchName = strings.ToLower(branchName)

	cmd := exec.Command("git", "checkout", "-b", branchName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create branch: %v\nOutput: %s", err, string(output))
	}

	fmt.Printf("✅ Created and switched to branch: %s\n", branchName)
	return nil
}
