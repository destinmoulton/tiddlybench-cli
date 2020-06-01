package config

import (
	"github.com/spf13/cobra"
)

var (
	shouldRunConfig = false
	rootCmd         = &cobra.Command{
		Use:   "tikli",
		Short: "A Tiddly Wiki cli application.",
		Long: `tikli is a CLI application for Tiddly Wiki.
Add journal entries to tiddlywiki directly from the command line.`,
	}
)

// SetupCLICommands initializes the cobra command line flags
func (c *Config) setupCLICommands() {

	cobra.OnInitialize(c.initConfig)

	rootCmd.PersistentFlags().BoolVarP(&shouldRunConfig, "config", "c", true, "Run the configuration prompt.")
}

func (c *Config) initConfig() {
	if shouldRunConfig {
		c.PromptForConfig()
	}
}
