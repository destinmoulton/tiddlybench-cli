package prompt

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"fmt"
	"os"
	"tiddlybench-cli/internal/config"

	tea "github.com/charmbracelet/bubbletea"
)

// PromptForSavePassword asks the user if they want to save their password
func (p *Prompt) PromptForSavePassword() {
	if err := tea.NewProgram(p.initialModel()).Start(); err != nil {
		fmt.Printf("Could not start Save Password prompt: %s\n", err)
		os.Exit(1)
	}
}

type spmodel struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func (p *Prompt) initialModel() spmodel {
	shouldSave := p.config.Get(config.CKShouldSavePassword)
	cursor := 0
	if shouldSave == "no" {
		cursor = 1
	}

	choices := []string{"Yes", "No"}
	selected := make(map[int]struct{})

	return spmodel{choices, cursor, selected}

}
func (m spmodel) Init() tea.Cmd {
	// no I/O
	return nil
}

func (m spmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m spmodel) View() string {
	// The header
	s := "Do you want to save the password?\n\n"
	s = s + "You may not want to save the password if you are not the administrator or are in a multi user environment.\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
