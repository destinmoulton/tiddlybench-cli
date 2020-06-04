package prompt

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"strings"
	"tiddly-cli/internal/config"
	"tiddly-cli/internal/logger"
	"tiddly-cli/internal/util"
	"time"
)

// Prompt is the struct for the Prompt
type Prompt struct {
	config *config.Config
	log    logger.Logger
}

// New returns a new Prompt
func New(config *config.Config, log logger.Logger) *Prompt {
	p := new(Prompt)
	p.log = log
	p.config = config
	return p
}

// PromptForConfig asks the user a series of config questions
func (p *Prompt) PromptForConfig() {
	url := p.promptURL()

	if !util.TestURL(url) {
		fmt.Println("That URL is unreachable")
	}
	username := p.promptUsername()
	titletemplate := p.promptTitleTemplate(c.Get("TitleTemplate"))
	savePassword := p.promptToSavePassword(c.Get("SavePassword"))
	password := ""
	if savePassword == CKYes {
		password = PromptForPassword()
	}

	// Set the c values
	if url != "" && username != "" {
		c.Set(CKURL, url)
		c.Set(TitleTemplate, titletemplate)
		c.Set(Username, username)
		c.Set(ShouldSavePassword, savePassword)
		c.Set(Password, password)
		c.Save()
	}
}

func (p *Prompt) promptURL() string {

	dflt := p.config.Get(p.config.CKURL)
	validate := func(input string) error {

		if !util.IsURL(input) {
			return errors.New("The URL is not valid. Should start with http:// or https://")
		}

		return nil
	}
	prompt := promptui.Prompt{
		Label:    "TiddlyWiki URL",
		Validate: validate,
		Default:  dflt,
	}
	url, err := prompt.Run()

	if err != nil {
		p.log.Fatal("Prompt Error. Unable to get the URL")
	}

	return url

}

func (p *Prompt) promptUsername() string {

	dflt := p.config.Get(p.config.CKUsername)
	prompt := promptui.Prompt{
		Label:   "Username",
		Default: dflt,
	}
	username, err := prompt.Run()

	if err != nil {
		p.log.Fatal("Prompt Error. Unable to get the Username")
	}

	return username
}

func (p *Prompt) promptDestination(dest string) (string, string) {

	destTitle := strings.Title(dest)
	currentTitle := p.config.GetNested(p.config.CKDestinations, dest, p.config.CKTitleTemplate)
	currentTags := p.config.GetNested(p.config.CKDestinations, dest, p.config.CKTags)

	titleprompt := promptui.Prompt{
		Label:   destTitle + " Title Template",
		Default: currentTitle,
	}
	title, terr := titleprompt.Run()

	if terr != nil {
		p.log.Fatal("Prompt Error. Unable to get the Destination title")
	}
	tagsprompt := promptui.Prompt{
		Label:   destTitle + " Tags",
		Default: currentTags,
	}
	tags, gerr := tagsprompt.Run()

	if gerr != nil {
		p.log.Fatal("Prompt Error. Unable to get the Destination tags")
	}

	return title, tags

}

func (p *Prompt) promptToSavePassword() string {
	dflt := p.config.Get(p.config.CKShouldSavePassword)
	prompt := promptui.Prompt{
		Label:     "Save Password?",
		IsConfirm: true,
		Default:   dflt,
	}

	result, err := prompt.Run()

	if result == "y" {
		result = Yes
	} else {
		result = No
	}

	if err != nil {
	}

	return result
}

// PromptForPassword uses promptui to get the basic auth password
func (p *Prompt) PromptForPassword() string {

	prompt := promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}
	password, err := prompt.Run()

	if err != nil {
	}

	return password

}

// PromptTiddlerTitle asks the user for the title of the
// tiddler to add
func (p *Prompt) PromptTiddlerTitle() string {

	title := c.Get("TitleTemplate")
	title = util.ConvertTiddlyTimeToGo(title)
	title = time.Now().Format(title)

	prompt := promptui.Prompt{
		Label:   "Tiddler Title",
		Default: title,
	}

	finaltitle, err := prompt.Run()
	if err != nil {
		c.log.Fatal("The prompt failed to process the title")
	}

	return finaltitle
}

// PromptTiddlerText gets the text for the tiddler
func (p *Prompt) PromptTiddlerText() string {

	prompt := promptui.Prompt{
		Label:   "New Tiddler Text",
		Default: "",
	}

	text, err := prompt.Run()
	if err != nil {
		c.log.Fatal("The prompt failed to process the text")
	}

	return text
}
