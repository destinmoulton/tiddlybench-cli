package cliflags

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"tiddlybench-cli/internal/config"
)

var (
	shouldRunConfig      = false
	shouldPaste          = false
	isBlockSelected      = false
	sendToInbox          = false
	sendToJournal        = false
	useEditor            = false
	tiddlerTitle         = ""
	defaultSelectedBlock = "default"
	password             = ""
	addText              = ""
)

// Setup configures the cli flags
func Setup() {

	flag.StringVarP(&addText, "add", "a", "", "Text to add to tiddler")
	flag.BoolVarP(&shouldRunConfig, "config", "c", false, "Run the configuration prompt.")
	flag.BoolVarP(&useEditor, "edit", "e", false, "Edit text in editor")
	flag.BoolVarP(&sendToInbox, "inbox", "i", false, "Add to Inbox")
	flag.BoolVarP(&sendToJournal, "journal", "j", false, "Add to Journal")
	flag.BoolVarP(&shouldPaste, "paste", "p", false, "Paste the contents of the clipboard.")
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

// GetAddText returns the string to add set by the -a flag
func GetAddText() string {
	return addText
}

// IsAddTextSet returns bool for whether the -a flag contains text
func IsAddTextSet() bool {
	fmt.Println("IsAddTextSet = ", GetAddText())
	return len(GetAddText()) > 0
}

// GetPassword returns the --password flag contents
func GetPassword() string {
	return password
}

// IsPasswordSet checks whether the password flag has anything
func IsPasswordSet() bool {
	return len(GetPassword()) > 0
}

// ShouldUseEditor returns bool for whether the -e flag is set
func ShouldUseEditor() bool {
	return useEditor
}
