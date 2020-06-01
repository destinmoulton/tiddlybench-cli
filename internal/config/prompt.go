package config

import (
	"errors"
	"github.com/manifoldco/promptui"
	"tiddly-cli/internal/util"
)

func (c *Config) promptForConfig() {
	url := promptURL()

	username := promptUsername()
	savePassword := selectSavePassword()
	password := ""
	if savePassword == "Yes" {
		password = PromptForPassword()
	}

	// Set the c values
	c.Set("Username", username)
	c.Set("Password", password)
	c.Set("URL", url)
	c.Set("SavePassword", savePassword)
	c.Save()
}

func promptURL() string {

	validate := func(input string) error {

		if !util.IsURL(input) {
			return errors.New("The URL is not valid. Should start with http:// or https://")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Base URL",
		Validate: validate,
	}
	url, err := prompt.Run()

	if err != nil {

	}

	return url

}

func promptUsername() string {

	prompt := promptui.Prompt{
		Label: "Username",
	}
	username, err := prompt.Run()

	if err != nil {

	}

	return username

}

// PromptForPassword uses promptui to get the basic auth password
func PromptForPassword() string {

	prompt := promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}
	password, err := prompt.Run()

	if err != nil {

	}

	return password

}

func selectSavePassword() string {
	prompt := promptui.Select{
		Label: "Save Password?",
		Items: []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()

	if err != nil {
	}

	return result
}
