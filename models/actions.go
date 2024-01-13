package actions

import (
	"fmt"
    "os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type Actions struct {
	cursor int

	Choices  []string
	Selected map[int]struct{}
}

func (m Actions) Init() tea.Cmd {
	return nil
}

func (m Actions) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.Choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.Selected[m.cursor]
			if ok {
				delete(m.Selected, m.cursor)
			} else {
                // fmt.Println("git", m.Choices[m.cursor])
                cmd := exec.Command("/usr/bin/git", "status")
                stdout, err := cmd.Output()
                fmt.Println(string(stdout))
                fmt.Println(err)
				m.Selected[m.cursor] = struct{}{}
			}
		}
	}

    return m, nil
}

func (m Actions) View() string {
	s := "What git command would you like to execute?\n\n"

	for i, choice := range m.Choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.Selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress <Ctrl>+c or q to quit.\n"

	return s
}
