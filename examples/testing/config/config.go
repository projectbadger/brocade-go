package config

import (
	"fmt"
	"os"

	"github.com/eschao/config"
)

var (
	// Main application configuration data.
	Cfg *ConfigApp
	// supplied through ldflags
	PackageName = "brocade"
	Version     = "development"
	BuildTime   = ""
)

type Config interface {
	SetupDefault()
}

// NewDefaultConfigApp returns a *ConfigApp with all the
// default values filled.
func NewDefaultConfigApp() *ConfigApp {
	cfg := &ConfigApp{}
	cfg.SetupDefault()
	return cfg
}

// GetParsedConfig returns a config, filled with
// environment variables, config file and CLI arguments
// values.
//
// Variable source parsing order:
// 	1. config file (if defined)
// 	2. environment variables
// 	3. CLI arguments
//
// CLI arguments always take precedence.
//
// If the CLI flag "-create-config" was provided along with
// a filepath, a config file will be created with default
// values and any parsed environment variables and CLI
// arguments.
func GetParsedConfig() (cfg *ConfigApp, err error) {
	cfg = NewDefaultConfigApp()
	config.ParseEnv(cfg)
	config.ParseCli(cfg)
	if cfg.File != "" {
		config.ParseConfigFile(cfg, cfg.File)
		// CLI and ENV arguments override config values.
		config.ParseEnv(cfg)
		config.ParseCli(cfg)
	}
	cfg.CheckIsVersion()
	err = cfg.Client.ParseURL()
	if err != nil {
		return nil, err
	}
	err = cfg.Client.GetPasswordFromStdin()
	if err != nil {
		return nil, err
	}

	if cfg.CreateConfig != "" {
		err = cfg.SaveToFile(cfg.CreateConfig)
		if err != nil {
			fmt.Printf("Error saving to file: '%q'", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	return cfg, nil
}

// SetupConfig sets Cfg variable to the parsed *ConfigApp
func SetupConfig() error {
	var err error
	Cfg, err = GetParsedConfig()
	return err
}

// init function. Runs when the module is included.
func init() {
	err := SetupConfig()
	if err != nil {
		panic(err)
	}
}
