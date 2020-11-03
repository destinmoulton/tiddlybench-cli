package prompt

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"strings"
	"tiddlybench-cli/internal/config"
	"tiddlybench-cli/internal/logger"
	"tiddlybench-cli/internal/util"
	"time"
)

// Prompt is the struct for the Prompt
type Prompt struct {
	config *config.Config
	log    logger.Logger
}

// New returns a new Prompt
func New(log logger.Logger, config *config.Config) *Prompt {
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

	// User auth prompts
	username := p.promptUsername()
	savePassword := p.promptToSavePassword()
	password := ""
	if savePassword == config.CKYes {
		password = p.PromptForPassword()
	}

	// destination prompts
	inboxTitle, inboxTags := p.promptDestination(config.CKInbox)
	journalTitle, journalTags := p.promptDestination(config.CKJournal)
	defaultDestination := p.promptDefaultDestination()

	// Set the c values
	if url != "" && username != "" {
		p.config.Set(config.CKURL, url)
		p.config.Set(config.CKUsername, username)
		p.config.Set(config.CKShouldSavePassword, savePassword)
		p.config.Set(config.CKPassword, password)
		p.config.SetNested([]string{config.CKDestinations, config.CKInbox, config.CKTitleTemplate}, inboxTitle)
		p.config.SetNested([]string{config.CKDestinations, config.CKInbox, config.CKTags}, inboxTags)
		p.config.SetNested([]string{config.CKDestinations, config.CKJournal, config.CKTitleTemplate}, journalTitle)
		p.config.SetNested([]string{config.CKDestinations, config.CKJournal, config.CKTags}, journalTags)
		p.config.Set(config.CKDestinations+"."+config.CKDefaultDestination, defaultDestination)
		p.config.Save()
	}
}

func (p *Prompt) promptURL() string {

	dflt := p.config.Get(config.CKURL)
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

	dflt := p.config.Get(config.CKUsername)
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

func (p *Prompt) promptDefaultDestination() string {

	prompt := promptui.Select{
		Label: "Select default destination",
		Items: []string{"Inbox", "Journal"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		p.log.Fatal("Prompt Error. Unable to get the default destination")
	}

	return strings.ToLower(result)
}

func (p *Prompt) promptDestination(dest string) (string, string) {

	destTitle := strings.Title(dest)
	currentTitle := p.config.GetNested(config.CKDestinations, dest, config.CKTitleTemplate)
	currentTags := p.config.GetNested(config.CKDestinations, dest, config.CKTags)

	titleprompt := promptui.Prompt{
		Label:   destTitle + " Tiddler Title",
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
	dflt := p.config.Get(config.CKShouldSavePassword)
	prompt := promptui.Prompt{
		Label:     "Save Password?",
		IsConfirm: true,
		Default:   dflt,
	}

	result, err := prompt.Run()

	if result == "y" {
		result = config.CKYes
	} else {
		result = config.CKNo
	}

	if err != nil {
		p.log.Fatal("Prompt Error. Unable to get the Save Password option")
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
func (p *Prompt) PromptTiddlerTitle(currentTitle string) string {

	title := currentTitle
	title = util.ConvertTiddlyTimeToGo(title)
	title = time.Now().Format(title)

	prompt := promptui.Prompt{
		Label:   "Tiddler Title",
		Default: title,
	}

	finaltitle, err := prompt.Run()
	if err != nil {
		p.log.Fatal("Prompt Error. The prompt failed to process the tiddler title")
	}

	return finaltitle
}

// PromptTiddlerText gets the text for the tiddler
func (p *Prompt) PromptTiddlerText() string {

	prompt := promptui.Prompt{
		Label:   "Tiddler Text to Add",
		Default: "",
	}

	text, err := prompt.Run()
	if err != nil {
		p.log.Fatal("Prompt Error. The prompt failed to process the tiddler text")
	}

	return text
}
