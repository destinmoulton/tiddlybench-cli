package dispatch

import (
	"fmt"
	"os"
	"tiddly-cli/internal/apicall"
	"tiddly-cli/internal/cliflags"
	"tiddly-cli/internal/clipboard"
	"tiddly-cli/internal/config"
	"tiddly-cli/internal/logger"
	"tiddly-cli/internal/piper"
	prompter "tiddly-cli/internal/prompt"
)

var (
	cfg    *config.Config
	pipe   *piper.Pipe
	prompt *prompter.Prompt
)

// Dispatch the right activity for the app
func Dispatch(log logger.Logger) {

	cfg = config.New(log)
	pipe = piper.New(log)
	prompt = prompter.New(cfg, log)

	cliflags.Setup()

	if pipe.IsPipeSet() {
		// Piping breaks the ability to use the prompt
		dispatchForPipe()
		return
	}

	if cliflags.ShouldPromptForConfig() || !cfg.IsConfigFileSet() {
		// Prompt to configure the username/password
		prompt.PromptForConfig()
		os.Exit(0)
	}

	if cfg.IsConfigFileSet() {

		savepassword := cfg.Get(config.CKShouldSavePassword)

		if savepassword == config.CKNo {
			// Password is not saved
			password := prompt.PromptForPassword()
			cfg.Set(config.CKPassword, password)

			// DO NOT cfg.Save after this point as we don't want to
			// write the password to the file per user request
		}

		tiddlerTitle := getTiddlerTitleFromFlags()
		if tiddlerTitle == "" {
			tiddlerTitle = prompt.PromptTiddlerTitle(tiddlerTitle)
		}
		api := apicall.New(log, cfg)
		block := cliflags.GetSelectedBlock()

		currentTiddler := api.GetTiddlerByName(tiddlerTitle)

		tidtext := ""
		if cliflags.ShouldPaste() {
			// Use the clipboard contents for the tiddler
			tidtext = clipboard.Paste(log)
		} else {
			// Prompt the user for the tiddler
			tidtext = prompt.PromptTiddlerText()
		}

		//Setup the block
		beginBlock := cfg.GetNested(config.CKBlocks, block, config.CKBegin)
		endBlock := cfg.GetNested(config.CKBlocks, block, config.CKEnd)
		tidtext = beginBlock + tidtext + endBlock
		if currentTiddler.Title != "" {
			fulltext := currentTiddler.Text + "\n" + tidtext
			api.UpdateTiddler(currentTiddler.Title, fulltext)
		} else {
			creator := cfg.Get(config.CKUsername)
			api.AddNewTiddler(tiddlerTitle, creator, tidtext)
		}

	}
}

func dispatchForPipe() {
	// Pipe is set, so can't use
	// any of the prompt methods

	// Must be configured
	requireConfigFile()

	// Must have password
	requirePasswordFlag()

	// Must have Inbox, Journal, or -t flag
	requireTiddlerTitleFlag()
}

func getTiddlerTitleFromFlags() string {
	tiddlerTitle := cliflags.GetTiddlerTitle()
	if tiddlerTitle != "" {
		return tiddlerTitle
	}

	sendTo := cliflags.GetSendTo()
	if sendTo != "" {
		return cfg.GetNested(config.CKDestinations, sendTo, config.CKTitleTemplate)
	}
	return ""
}

func requireConfigFile() {

	if !cfg.IsConfigFileSet() {
		fmt.Println("Config file has not been set.")
		fmt.Println("Run tikli with -c option to configure")
		os.Exit(1)
	}
}

func requirePasswordFlag() {
	if !cfg.IsPasswordSaved() {
		fmt.Println("Password is required, but it is not saved in the config file.")
		fmt.Println("Add the password to the command line arguments: tikli --password 'YourPass'")
		os.Exit(1)
	}
}

func requireTiddlerTitleFlag() {
	hasTiddlerTitle := cliflags.GetTiddlerTitle() != ""
	hasSendTo := cliflags.GetSendTo() != ""
	if hasTiddlerTitle && hasSendTo {
		fmt.Println("You have set too many destination tiddlers.")
		fmt.Println("Include just one of -i, -j, or -t.")
		os.Exit(1)
	}
	if !hasTiddlerTitle && !hasSendTo {
		fmt.Println("You must include a destination tiddler.")
		fmt.Println("Include -i (inbox), -j (journal), or -t (custom tiddler).")
		os.Exit(1)
	}
}
