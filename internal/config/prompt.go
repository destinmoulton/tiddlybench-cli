package config

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"tiddly-cli/internal/util"
)

func PromptForConfig() {
	url := promptURL()

	username := promptUsername()
	savePassword := selectSavePassword()

	if savePassword == "Yes" {

		password := promptPassword()

		fmt.Println(password)
	}
	fmt.Println(url)
	fmt.Println(username)
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

func promptPassword() string {

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
