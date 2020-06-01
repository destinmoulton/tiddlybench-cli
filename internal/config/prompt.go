package config

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"tiddly-cli/internal/util"
)

func (c *Config) DispatchPrompt() {
	//Check if the URL is already stored
	url := c.Get("URL")

	if url != "" {
	} else {
		c.PromptForConfig()
	}
}

// PromptForConfig asks the user a series of config questions
func (c *Config) PromptForConfig() {
	url := promptURL(c.Get("URL"))

	if !util.TestURL(url) {
		fmt.Println("That URL is unreachable")
	}
	username := promptUsername(c.Get("Username"))
	titletemplate := promptTitleTemplate(c.Get("TitleTemplate"))
	savePassword := promptToSavePassword(c.Get("SavePassword"))
	password := ""
	if savePassword == "y" {
		password = PromptForPassword()
	}

	// Set the c values
	if url != "" && username != "" {
		c.Set("URL", url)
		c.Set("TitleTemplate", titletemplate)
		c.Set("Username", username)
		c.Set("SavePassword", savePassword)
		c.Set("Password", password)
		c.Save()
	}
}

func promptURL(dflt string) string {

	validate := func(input string) error {

		if !util.IsURL(input) {
			return errors.New("The URL is not valid. Should start with http:// or https://")
		}

		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Base URL",
		Validate: validate,
		Default:  dflt,
	}
	url, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return url

}

func promptUsername(dflt string) string {

	prompt := promptui.Prompt{
		Label:   "Username",
		Default: dflt,
	}
	username, err := prompt.Run()

	if err != nil {

	}

	return username

}

func promptTitleTemplate(dflt string) string {

	prompt := promptui.Prompt{
		Label:   "Title Template (can use some TimeFormat options)",
		Default: dflt,
	}
	journal, err := prompt.Run()

	if err != nil {

	}

	return journal

}

func promptToSavePassword(dflt string) string {
	prompt := promptui.Prompt{
		Label:     "Save Password?",
		IsConfirm: true,
		Default:   dflt,
	}

	result, err := prompt.Run()

	if err != nil {
	}

	return result
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
