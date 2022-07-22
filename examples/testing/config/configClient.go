package config

import (
	"bufio"
	"errors"
	"net/url"
	"os"
	"strings"
)

// "github.com/projectbadger/brocade-go/filters"

// ConfigApp holds all the application configuration data.
type ConfigClient struct {
	Host               string `env:"HOST" yaml:"host" json:"host" cli:"host FOS Host\nExample:\n  192.168.1.2\n  [::]:80\n  https://10.10.10.10/rest"`
	ipAddress          *url.URL
	BaseURI            string `env:"BASE_URI" yaml:"base_uri" json:"base_uri" cli:"base-uri FOS base URI for the REST API\nExample:\n  /rest\nDefault: /rest"`
	Username           string `env:"USERNAME" yaml:"username" json:"username" cli:"username API login username"`
	Password           string `env:"PASSWORD" yaml:"password" json:"password" cli:"password API login password\nDefaults to '-' as stdin"`
	Session            bool   `env:"SESSION" yaml:"session" json:"session" cli:"session Session authentication\nEstablishes a session instead of providing username:password with every request\nManual login and logout"`
	CheckPorts         bool   `env:"CHECK_PORTS" yaml:"check_ports" json:"check_ports" cli:"check-ports Check ports"`
	CheckPowerSupplies bool   `env:"CHECK_POWER_SUPPLIES" yaml:"check_power_supplies" json:"check_poer_supplies" cli:"check-power-supplies Check power supplies"`
	CheckFans          bool   `env:"CHECK_FANS" yaml:"check_fans" json:"check_fans" cli:"check-fans Check fans"`
}

func (c *ConfigClient) GetURL() string {
	return c.ipAddress.String()
}

// SetupDefault sets up default config data.
func (c *ConfigClient) SetupDefault() {
	c.BaseURI = "/rest"
	c.Password = "-"
}

func (c *ConfigClient) ParseURL() (err error) {
	if !strings.Contains(c.Host, "://") {
		c.Host = "http://" + c.Host
	} else if !strings.HasPrefix(c.Host, "http://") && !strings.HasPrefix(c.Host, "https://") {
		return errors.New("invalid schema")
	}
	c.ipAddress, err = url.Parse(c.Host)
	if err != nil {
		return
	}
	if c.ipAddress.Scheme == "" {
		c.ipAddress.Scheme = "http"
	}
	if c.ipAddress.Host == "" {
		return errors.New("invalid host")
	}
	if c.BaseURI == "" {
		c.BaseURI = c.ipAddress.Path
	} else {
		c.ipAddress.Path = c.BaseURI
	}
	return
}

func (c *ConfigClient) GetPasswordFromStdin() error {
	if c.Password == "-" {
		// nBytes, nChunks := int64(0), int64(0)
		stat, err := os.Stdin.Stat()
		if err != nil {
			return err
		}
		if stat.Size() == 0 {
			return nil
		}
		if stat.Mode()&os.ModeNamedPipe == 0 {
			return nil
		}
		r := bufio.NewReader(os.Stdin)
		buf := make([]byte, 0, 512)
		n, err := r.Read(buf[:cap(buf)])
		if err != nil {
			return err
		}
		c.Password = string(buf[:n])
	}
	return nil
}
