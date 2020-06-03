package config

import (
	flag "github.com/spf13/pflag"
)

var (
	shouldRunConfig = false
	shouldPaste     = false
	selectedBlock   = "bullet"
)

// CheckCLIFlags initializes the cobra command line flags parser
func (c *Config) CheckCLIFlags() {

	flag.BoolVarP(&shouldRunConfig, "config", "c", false, "Run the configuration prompt.")
	flag.BoolVarP(&shouldPaste, "paste", "p", false, "Paste the contents of the clipboard.")

	for key := range Blocks {
		flag.StringVar(&selectedBlock, key, key, "Wrap your content with a "+key+" block")
	}
	flag.Parse()
}

// ShouldPromptForConfig returns bool as to whether the config should be rerun
func (c *Config) ShouldPromptForConfig() bool {
	return shouldRunConfig
}

// FlagShouldPaste returns bool as to whether the tiddler should be composed of a paste
func (c *Config) FlagShouldPaste() bool {
	return shouldPaste
}

func (c *Config) GetSelectedBlock() string {
	return selectedBlock
}
