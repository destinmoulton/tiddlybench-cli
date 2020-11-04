package prompt

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"fmt"
	"os"
	"tiddlybench-cli/internal/config"

	tea "github.com/charmbracelet/bubbletea"
	te "github.com/muesli/termenv"
)

type spmodel struct {
	choices  []string // items on the to-do list
	cursor   int      // which to-do list item our cursor is pointing at
	selected int      // which to-do items are selected
}

var (
	spContinueButtonIndex   = 2
	spContinueButtonBlurred = "[ " + te.String("Continue").Foreground(color("205")).String() + " ]"
	spContinueButtonFocused = "[ " + te.String("Continue").Foreground(color("240")).String() + " ]"
)

// PromptForSavePassword asks the user if they want to save their password
func (p *Prompt) PromptForSavePassword() string {
	model := p.initialModel()
	if err := tea.NewProgram(&model).Start(); err != nil {
		fmt.Printf("Could not start Save Password prompt: %s\n", err)
		os.Exit(1)
	}

	return model.getSelected()

}

func (m *spmodel) getSelected() string {
	if m.selected == 0 {
		return config.CKNo
	}
	return config.CKYes
}

func (p *Prompt) initialModel() spmodel {
	shouldSave := p.config.Get(config.CKShouldSavePassword)
	cursor := 0
	selected := 0
	if shouldSave == config.CKNo {
		selected = 1
	}

	choices := []string{"No", "Yes"}

	return spmodel{choices, cursor, selected}

}
func (m *spmodel) Init() tea.Cmd {
	// no I/O
	return nil
}

func (m *spmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			// Submit button adds one, so not `len(m.choices)-1`
			if m.cursor < len(m.choices) {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			if m.cursor < len(m.choices)-1 {

				m.selected = m.cursor
			} else if m.cursor == spContinueButtonIndex {
				return m, tea.Quit
			}
		}
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m *spmodel) View() string {
	// The header
	s := "Do you want to save the password?\n\n"
	s = s + "You may not want to save the password\nif you are not the administrator or are in a multi user environment.\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if i == m.selected {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	if spContinueButtonIndex == m.cursor {
		s += fmt.Sprintf("%s %s\n", ">", spContinueButtonFocused)
	} else {
		s += fmt.Sprintf("%s %s\n", " ", spContinueButtonBlurred)
	}
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
