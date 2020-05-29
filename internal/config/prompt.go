package config

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"net/url"
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

		if !isUrl(input) {
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
func isUrl(str string) bool {
	// https://stackoverflow.com/questions/31480710/validate-url-with-standard-package-in-go
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
