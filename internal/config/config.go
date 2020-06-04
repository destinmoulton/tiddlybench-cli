package config

import (
	"github.com/spf13/viper"
	"os"
	"path"
	"tiddly-cli/internal/logger"
)

var configSubdir = "tiddly-cli"
var configName = "config"
var configType = "json"
var configExt = "json"

type destination struct {
	Tags          string
	TitleTemplate string
}

var destinationDefaults = map[string]destination{
	CKInbox:   {Tags: "", TitleTemplate: "Inbox"},
	CKJournal: {Tags: "journal", TitleTemplate: "YYYY-0MM-0DD Journal"},
}

var defaults = map[string]string{
	CKURL:                "",
	CKUsername:           "",
	CKShouldSavePassword: CKNo,
	CKPassword:           "",
}

type blockParts struct {
	Begin string
	End   string
}

// Blocks define a slice of blocks for tiddlers
var Blocks = map[string]blockParts{
	"code":   {Begin: "\n\n```\n", End: "\n```\n\n"},
	"bullet": {Begin: "\n* ", End: "\n"},
	"number": {Begin: "\n# ", End: "\n"},
	"quote":  {Begin: "\n\n<<<\n", End: "\n<<<\n\n"},
	"h1":     {Begin: "! ", End: "\n"},
	"h2":     {Begin: "!! ", End: "\n"},
	"h3":     {Begin: "!!! ", End: "\n"},
	"h4":     {Begin: "!!!! ", End: "\n"},
	"h5":     {Begin: "!!!!! ", End: "\n"},
	"custom": {Begin: "\n\n", End: "\n\n"},
}

// Config is a struct for the configuration interface
type Config struct {
	log   logger.Logger
	viper *viper.Viper
}

// New creates a Config object
func New(log logger.Logger) *Config {
	c := new(Config)
	c.log = log
	c.viper = viper.New()
	c.initializeDefaults()
	c.setup()
	return c
}

func (c *Config) initializeDefaults() {
	// Basic defaults
	for key, def := range defaults {
		c.viper.SetDefault(key, def)
	}

	// Destination defaults
	c.viper.SetDefault(CKDestinations, nil)
	for key, def := range destinationDefaults {
		c.viper.SetDefault(CKDestinations+"."+key, nil)
		c.viper.SetDefault(CKDestinations+"."+key+"."+CKTags, def.Tags)
		c.viper.SetDefault(CKDestinations+"."+key+"."+CKTitleTemplate, def.TitleTemplate)
	}

	// Block defaults
	c.viper.SetDefault(CKBlocks, nil)
	for key, parts := range Blocks {
		c.viper.SetDefault(CKBlocks+"."+key, nil)
		c.viper.SetDefault(CKBlocks+"."+key+"."+CKBegin, parts.Begin)
		c.viper.SetDefault(CKBlocks+"."+key+"."+CKEnd, parts.End)
	}
}

func (c *Config) setup() {

	configPath := c.getConfigPath()
	c.setupConfigDir(configPath)

	c.viper.SetConfigName(configName)
	c.viper.SetConfigType(configType)
	c.viper.AddConfigPath(configPath)
	if err := c.viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Create a config with defaults
			c.Save()

			if serr := c.viper.ReadInConfig(); serr != nil {
				c.log.Fatal("Unable to find the config, even after trying to create it. Something weird this way goes.")
			}

		} else {
			// Config file was found but another error was produced
			c.log.Fatal(err)
		}
	}
}

// Get returns a config value by key
func (c *Config) Get(key string) string {
	return c.viper.GetString(key)
}

// GetNested returns a nested config valudflt stringe
func (c *Config) GetNested(one string, two string, three string) string {
	return c.viper.GetString(one + "." + two + "." + three)
}

// Set assigns a config value to a key
func (c *Config) Set(key string, value string) {
	c.viper.Set(key, value)
}

// Save the config to the file
func (c *Config) Save() {
	// Config file not found; ignore error if desired
	if verr := c.viper.WriteConfig(); verr != nil {
		c.log.Fatal("Unable to write the config file.", verr)
	}
}

func (c *Config) setupConfigDir(dir string) {

	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			// Folder
			// Try making the folder
			mkerr := os.Mkdir(dir, 0700)
			if mkerr != nil {
				c.log.Fatal("Unable to create the config directory [" + dir + "].")
			}

			c.createEmptyConfigFile()

			// Save the defaults
			c.Save()
		} else {
			// other error
			c.log.Fatal(err)
		}
	}
}

func (c *Config) createEmptyConfigFile() {
	configFile := c.getFullConfigPath()
	f, ferr := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY, 0700)
	if ferr != nil {
		c.log.Fatal("Unable to create config file [" + configFile + "]")
	}
	// Empty json
	f.WriteString("{}")

	if cerr := f.Close(); cerr != nil {
		c.log.Fatal(cerr)
	}
}

func (c *Config) getFullConfigPath() string {
	dir := c.getConfigPath()
	filename := configName + "." + configExt
	return path.Join(dir, filename)
}

func (c *Config) getConfigPath() string {

	configDir, err := os.UserConfigDir()

	if err != nil {
		c.log.Fatal("Unable to find a config directory for this user.")
	}

	return path.Join(configDir, configSubdir)
}

// IsConfigFileSet returns boolean if the config file is setup
func (c *Config) IsConfigFileSet() bool {
	url := c.Get(CKURL)

	// May need to alter what is checked
	if url != "" {
		return true
	}
	return false
}

// IsPasswordSaved returns boolean to determine if password is set in config
func (c *Config) IsPasswordSaved() bool {
	shouldSave := c.Get(CKShouldSavePassword) == CKYes
	isPassword := len(c.Get(CKPassword)) > 0
	return shouldSave && isPassword
}
