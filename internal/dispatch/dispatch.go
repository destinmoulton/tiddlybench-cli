package dispatch

import (
	//"fmt"
	"tiddly-cli/internal/apicall"
	"tiddly-cli/internal/clipboard"
	"tiddly-cli/internal/config"
	"tiddly-cli/internal/logger"
)

// Dispatch the right activity for the app
func Dispatch(log logger.Logger) {

	cfg := config.New(log)
	cfg.CheckCLIFlags()
	if cfg.ShouldPromptForConfig() || !cfg.IsConfigFileSet() {
		// Prompt to configure the username/password
		cfg.PromptForConfig()
	}

	if cfg.IsConfigFileSet() {

		savepassword := cfg.Get("SavePassword")

		if savepassword == config.No {
			// Password is not saved
			password := config.PromptForPassword()
			cfg.Set(config.Password, password)

			// DO NOT c.Save after this point as we don't want to
			// write the password to the file per user request
		}

		tiddlerTitle := cfg.PromptTiddlerTitle()
		api := apicall.New(log, cfg)

		tiddler := api.GetTiddlerByName(tiddlerTitle)

		tidtext := ""
		if cfg.FlagShouldPaste() {
			// Use the clipboard contents for the tiddler
			tidtext = clipboard.Paste(log)
		} else {
			// Prompt the user for the tiddler
			tidtext = cfg.PromptTiddlerText()
		}
		if tiddler.Title != "" {
			fulltext := tiddler.Text + "\n" + tidtext
			api.UpdateTiddler(tiddler.Title, fulltext)
		} else {
			creator := cfg.Get(config.Username)
			api.AddNewTiddler(tiddlerTitle, creator, tidtext)
		}

		// The first argument is the journal entry
		// fmt.Println("Journal entry")
		// fmt.Println(os.Args[1])
	}
}
