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
func (p *Prompt) PromptConfigDispatch() {

	// User auth prompts
	connectionDetails := p.promptForConnection()
	shouldSavePassword := p.PromptForSavePassword()
	fmt.Println(connectionDetails)
	fmt.Println(shouldSavePassword)
	/*
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
	*/
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
