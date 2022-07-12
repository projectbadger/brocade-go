package config

import (
	"fmt"

	// "github.com/projectbadger/brocade-go/filters"

	"os"

	"gopkg.in/yaml.v2"
)

// ConfigApp holds all the application configuration data.
type ConfigApp struct {
	File         string `default:"" json:"-" yaml:"-" cli:"config Config file path\n     "`
	CreateConfig string `default:"" yaml:"-" json:"-" cli:"create-config Create a named default config file with cli parameters and environment variables.\n   "`
	DebugMode    bool   `default:"false" yaml:"-" json:"-" cli:"debug Run the app in debug mode\n   "`
	Version      bool   `default:"false" json:"-" yaml:"-" cli:"version Print version"`
	VersionShort bool   `default:"false" json:"-" yaml:"-" cli:"v Print version"`
	Log          ConfigLog
	Client       ConfigClient
}

func (c *ConfigApp) IsVersion() bool {
	return c.Version || c.VersionShort
}

func (c *ConfigApp) CheckIsVersion() {
	if c.Version || c.VersionShort {
		fmt.Printf("Package:    %s\n", PackageName)
		fmt.Printf("Version:    %s\n", Version)
		fmt.Printf("Build Time: %s\n", BuildTime)
		os.Exit(0)
	}
}

// SaveToFile saves the config to a file in YAML format
func (cfg *ConfigApp) SaveToFile(path string) error {
	if path == "" {
		return nil
	}
	configYaml, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, configYaml, 0644)
}

// SetupDefault sets up default config data.
func (c *ConfigApp) SetupDefault() {
	c.Client.SetupDefault()
	c.Log.SetupDefault()
	// c.ServiceName = "webmail"
}
