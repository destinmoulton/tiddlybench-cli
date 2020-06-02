package dispatch

import (
	"fmt"
	"tiddly-cli/internal/apicall"
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

		if savepassword == "N" {
			// Password is not saved
			password := config.PromptForPassword()
			cfg.Set("Password", password)

			// DO NOT c.Save after this point as we don't want to
			// write the password to the file per user request
		}

		tiddlerTitle := cfg.PromptTiddlerTitle()
		api := apicall.New(log, cfg)

		tiddler := api.GetTiddlerByName(tiddlerTitle)

		fmt.Print(tiddler)
		// The first argument is the journal entry
		// fmt.Println("Journal entry")
		// fmt.Println(os.Args[1])
	}
}
