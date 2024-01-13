package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	actions "github.com/haydenrou/gti/models"
)

func initialActions() actions.Actions {
	return actions.Actions{
        Choices: []string{"status"},
        Selected: make(map[int]struct{}),
    }
}

func main() {
    p := tea.NewProgram(initialActions())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

