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

type setting struct {
	datatype     string
	defaultvalue string
}

var defaults = map[string]setting{
	"Username":     {"string", ""},
	"Password":     {"string", ""},
	"URL":          {"string", ""},
	"SavePassword": {"string", ""},
}

type Config struct {
	log   logger.Logger
	viper *viper.Viper
}

func New(log logger.Logger) *Config {
	c := new(Config)
	c.log = log
	c.viper = viper.New()
	c.initializeDefaults()
	c.setup()
	return c
}

func (c *Config) initializeDefaults() {
	for key, setting := range defaults {
		c.viper.SetDefault(key, setting.defaultvalue)
	}
}

func (c *Config) Get(key string) string {
	_, ok := defaults[key]
	if !ok {
		c.log.Fatal("The config key " + key + " does not exist.")
	}
	return c.viper.GetString(key)
}

func (c *Config) Set(key string, value string) {
	c.viper.Set(key, value)
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
