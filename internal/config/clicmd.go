package config

import (
	flag "github.com/spf13/pflag"
)

var (
	shouldRunConfig = false
)

// CheckCLIFlags initializes the cobra command line flags parser
func (c *Config) CheckCLIFlags() {

	flag.BoolVarP(&shouldRunConfig, "config", "c", false, "Run the configuration prompt.")
	flag.Parse()
}

// ShouldPromptForConfig returns bool as to whether the config should be rerun
func (c *Config) ShouldPromptForConfig() bool {
	return shouldRunConfig
}
