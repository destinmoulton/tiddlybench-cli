package cliflags

import (
	flag "github.com/spf13/pflag"
	"tiddly-cli/internal/config"
)

var (
	shouldRunConfig      = false
	shouldPaste          = false
	isBlockSelected      = false
	sendToInbox          = false
	sendToJournal        = false
	tiddlerTitle         = ""
	defaultSelectedBlock = "bullet"
	password             = ""
)

// Setup configures the cli flags
func Setup() {

	flag.BoolVarP(&shouldRunConfig, "config", "c", false, "Run the configuration prompt.")
	flag.BoolVarP(&shouldPaste, "paste", "p", false, "Paste the contents of the clipboard.")
	flag.BoolVarP(&sendToInbox, "inbox", "i", false, "Add to Inbox")
	flag.BoolVarP(&sendToJournal, "inbox", "i", false, "Add to Journal")
	flag.StringVarP(&tiddlerTitle, "tiddler", "t", "", "tiddler to create or add to")
	flag.StringVar(&password, "password", "", "password to use for authentication")

	for key := range config.Blocks {
		flag.BoolVar(&isBlockSelected, key, false, "Wrap your content with a "+key+" block")
	}

	flag.Parse()
}

// ShouldPromptForConfig returns bool as to whether the config should be rerun
func ShouldPromptForConfig() bool {
	return shouldRunConfig
}

// ShouldPaste returns bool as to whether the tiddler should be composed of a paste
func ShouldPaste() bool {
	return shouldPaste
}

// GetSelectedBlock returns the string of the selected block type
func GetSelectedBlock() string {
	selectedBlock := defaultSelectedBlock

	// Visit each flag and determine if
	// it is a block
	visitor := func(f *flag.Flag) {
		if _, ok := config.Blocks[f.Name]; ok {
			selectedBlock = f.Name
		}
	}
	flag.Visit(visitor)
	return selectedBlock
}

// GetTiddlerTitle returns the string of the tiddler title
func GetTiddlerTitle() string {
	return tiddlerTitle
}

// GetSendTo returns either "inbox" or "journal"
func GetSendTo() string {
	if sendToInbox {
		return "inbox"
	}
	if sendToJournal {
		return "journal"
	}
	return ""
}
