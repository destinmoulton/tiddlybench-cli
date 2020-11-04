package prompt

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	te "github.com/muesli/termenv"

	"tiddlybench-cli/internal/util"
)

const focusedTextColor = "205"

var (
	color               = te.ColorProfile().Color
	red                 = color("196")
	focusedPrompt       = te.String("> ").Foreground(color("205")).String()
	promptPrefix        = "> "
	focusedSubmitButton = "[ " + te.String("Submit").Foreground(color("205")).String() + " ]"
	blurredSubmitButton = "[ " + te.String("Submit").Foreground(color("240")).String() + " ]"
)

type connModel struct {
	index         int
	urlInput      textinput.Model
	usernameInput textinput.Model
	passwordInput textinput.Model
	error         string
	submitButton  string
}

// PromptForConnection presents the user with a config selection
func (p *Prompt) promptForConnection() string {

	model := buildInitialConnModel()
	if err := tea.NewProgram(&model).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}

	return model.usernameInput.Value()
}

func buildInitialConnModel() connModel {
	name := textinput.NewModel()
	name.Placeholder = "https://address-to-your-tiddlywiki.com"
	name.Focus()
	name.Prompt = "TiddlyWiki Address " + focusedPrompt
	name.TextColor = focusedTextColor
	name.CharLimit = 255

	email := textinput.NewModel()
	email.Placeholder = "Username"
	email.Prompt = "TiddlyWiki Username " + promptPrefix
	email.CharLimit = 128

	password := textinput.NewModel()
	password.Placeholder = "Password"
	password.Prompt = promptPrefix
	password.EchoMode = textinput.EchoPassword
	password.EchoCharacter = '•'
	password.CharLimit = 128

	error := ""

	return connModel{0, name, email, password, error, blurredSubmitButton}

}
func (m *connModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *connModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		// Cycle between inputs
		case "tab", "shift+tab", "enter", "up", "down":

			inputs := []textinput.Model{
				m.urlInput,
				m.usernameInput,
				m.passwordInput,
			}

			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.index == len(inputs) {
				if vuErr := m.validateURL(); vuErr != nil {
					m.error = vuErr.Error()
				} else if vsErr := m.validateUsername(); vsErr != nil {
					m.error = vsErr.Error()
				} else {

					return m, tea.Quit
				}
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.index--
			} else {
				m.index++
			}

			if m.index > len(inputs) {
				m.index = 0
			} else if m.index < 0 {
				m.index = len(inputs)
			}

			for i := 0; i <= len(inputs)-1; i++ {
				if i == m.index {
					// Set focused state
					inputs[i].Focus()
					inputs[i].Prompt = focusedPrompt
					inputs[i].TextColor = focusedTextColor
					continue
				}
				// Remove focused state
				inputs[i].Blur()
				inputs[i].Prompt = promptPrefix
				inputs[i].TextColor = ""
			}

			m.urlInput = inputs[0]
			m.usernameInput = inputs[1]

			if m.index == len(inputs) {
				m.submitButton = focusedSubmitButton
			} else {
				m.submitButton = blurredSubmitButton
			}

			return m, nil
		}
	}

	// Handle character input and blinks
	m, cmd = updateInputs(msg, m)
	return m, cmd
}

// Pass messages and models through to text input components. Only text inputs
// with Focus() set will respond, so it's safe to simply update all of them
// here without any further logic.
func updateInputs(msg tea.Msg, m *connModel) (*connModel, tea.Cmd) {

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.urlInput, cmd = m.urlInput.Update(msg)
	cmds = append(cmds, cmd)

	m.usernameInput, cmd = m.usernameInput.Update(msg)
	cmds = append(cmds, cmd)

	m.passwordInput, cmd = m.passwordInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *connModel) View() string {
	s := "\n"

	inputs := []string{
		m.urlInput.View(),
		m.usernameInput.View(),
		m.passwordInput.View(),
	}

	for i := 0; i < len(inputs); i++ {
		s += inputs[i]
		if i < len(inputs)-1 {
			s += "\n"
		}
	}

	if m.error != "" {
		s += "\n\n" + te.String(m.error).Foreground(red).String()
	}

	s += "\n\n" + m.submitButton + "\n"
	return s
}

func (m *connModel) validateURL() error {
	url := m.urlInput.Value()

	if url == "" {
		return errors.New("You must include a URL for your TiddlyWiki server")
	}

	if !util.IsURL(url) {
		return errors.New("The URL is invalid")
	}

	if !util.TestURL(url) {
		return errors.New("The TiddlyWiki server is unreachable")
	}

	return nil
}

func (m *connModel) validateUsername() error {
	url := m.usernameInput.Value()

	if url == "" {
		return errors.New("You must include the username for the TiddlyWiki server")
	}

	return nil
}
