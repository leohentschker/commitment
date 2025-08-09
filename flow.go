package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
)

type step int

const (
	selectType step = iota
	enterDescription
	confirm
	done
)

type flowModel struct {
	step         step
	cursor       int
	selectedType string
	description  string
	action       string
	cancelled    bool
}

func newFlow(action string) flowModel {
	return flowModel{
		step:   selectType,
		action: action,
	}
}

func newFlowWithType(action, selectedType string) flowModel {
	return flowModel{
		step:         enterDescription,
		action:       action,
		selectedType: selectedType,
	}
}

func (m flowModel) Init() tea.Cmd {
	return nil
}

func (m flowModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.step {
		case selectType:
			switch msg.String() {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(commitTypes)-1 {
					m.cursor++
				}
			case "enter":
				m.selectedType = commitTypes[m.cursor]
				m.step = enterDescription
			case "ctrl+c", "q":
				m.cancelled = true
				return m, tea.Quit
			}

		case enterDescription:
			switch msg.String() {
			case "ctrl+c":
				m.cancelled = true
				return m, tea.Quit
			case "enter":
				if strings.TrimSpace(m.description) != "" {
					m.step = confirm
				}
			case "backspace":
				if len(m.description) > 0 {
					m.description = m.description[:len(m.description)-1]
				}
			default:
				if len(msg.String()) == 1 && len(m.description) < 40 {
					m.description += msg.String()
				}
			}

		case confirm:
			switch msg.String() {
			case "y", "Y", "enter":
				m.step = done
				return m, tea.Quit
			case "n", "N", "ctrl+c":
				m.cancelled = true
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m flowModel) View() string {
	switch m.step {
	case selectType:
		s := fmt.Sprintf("Select %s type:\n\n", m.action)

		for i, choice := range commitTypes {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			description := commitTypeDescriptions[choice]
			s += fmt.Sprintf("%s %s - %s\n", cursor, choice, description)
		}

		s += "\nPress Enter to select, q to quit"
		return s

	case enterDescription:
		s := fmt.Sprintf("Selected: %s\n\n", m.selectedType)
		s += fmt.Sprintf("Enter %s description:\n", m.action)
		s += fmt.Sprintf("> %s\n", m.description)
		s += fmt.Sprintf("(%d/40 characters)\n\n", len(m.description))
		s += "Press Enter to continue, Ctrl+C to cancel"
		return s

	case confirm:
		var preview string
		if m.action == "branch" {
			branchName := fmt.Sprintf("%s/%s", m.selectedType, strings.ReplaceAll(m.description, " ", "-"))
			branchName = strings.ToLower(branchName)
			preview = fmt.Sprintf("Branch: %s", branchName)
		} else {
			preview = fmt.Sprintf("Commit: %s: %s", m.selectedType, m.description)
		}

		s := fmt.Sprintf("Preview:\n%s\n\n", preview)
		s += "Confirm? (Enter/y for yes, n for no): "
		return s

	default:
		return ""
	}
}
