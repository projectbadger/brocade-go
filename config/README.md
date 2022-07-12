
# config

```go
import github.com/projectbadger/brocade-go/config
```

Package config handles all the configuration data.

Before use, the package must be initialized by the function `SetupConfig() error`. It is normally already initialized by the `init()` function.
Example:

```go
err := config.SetupConfig()

```
The configuration data is accessible from the `Cfg` variable.

Example:

```go
configFilePath := config.Cfg.File

```


<details>
<summary>Example</summary>
<p>

```go
{

	fmt.Println(config.Cfg.Version)

}
```
</p></details>

## Index

- [SetupConfig() error](#func-setupconfig-error)

- [type Config](#type-config)
- [type ConfigApp](#type-configapp)
  - [GetParsedConfig() (ConfigApp, error)](#func-getparsedconfig-configapp-error)
  - [NewDefaultConfigApp() ConfigApp](#func-newdefaultconfigapp-configapp)
  - [CheckIsVersion()](#func-configapp-checkisversion)
  - [IsVersion() bool](#func-configapp-isversion-bool)
  - [SaveToFile(string) error](#func-configapp-savetofilestring-error)
  - [SetupDefault()](#func-configapp-setupdefault)
- [type ConfigClient](#type-configclient)
  - [GetPasswordFromStdin() error](#func-configclient-getpasswordfromstdin-error)
  - [GetURL() string](#func-configclient-geturl-string)
  - [ParseURL() error](#func-configclient-parseurl-error)
  - [SetupDefault()](#func-configclient-setupdefault)
- [type ConfigLog](#type-configlog)
  - [SetupDefault()](#func-configlog-setupdefault)
- [type LogLevel](#type-loglevel)
- [type LogType](#type-logtype)
  - [String() string](#func-logtype-string-string)
- [Variables](#variables)

## Variables
```go
var (
	// Main application configuration data.
	Cfg	*ConfigApp
	// supplied through ldflags
	PackageName	= "brocade"
	Version		= "development"
	BuildTime	= ""
)

```

## func [SetupConfig() error](<config.go#L78>)

SetupConfig sets Cfg variable to the parsed *ConfigApp


```go
func SetupConfig() error
```


## type [Config](<config.go#L19>)
```go
type Config interface {
	SetupDefault()
}
```

## type [ConfigApp](<configApp.go#L14>)

ConfigApp holds all the application configuration data.
```go
type ConfigApp struct {
	File		string	`default:"" json:"-" yaml:"-" cli:"config Config file path\n     "`
	CreateConfig	string	`default:"" yaml:"-" json:"-" cli:"create-config Create a named default config file with cli parameters and environment variables.\n   "`
	DebugMode	bool	`default:"false" yaml:"-" json:"-" cli:"debug Run the app in debug mode\n   "`
	Version		bool	`default:"false" json:"-" yaml:"-" cli:"version Print version"`
	VersionShort	bool	`default:"false" json:"-" yaml:"-" cli:"v Print version"`
	Log		ConfigLog
	Client		ConfigClient
}
```

## func [GetParsedConfig() (ConfigApp, error)](<config.go#L46>)

GetParsedConfig returns a config, filled with
environment variables, config file and CLI arguments
values.

Variable source parsing order:
```go
1. config file (if defined)
2. environment variables
3. CLI arguments

```
CLI arguments always take precedence.

If the CLI flag "-create-config" was provided along with
a filepath, a config file will be created with default
values and any parsed environment variables and CLI
arguments.


```go
func GetParsedConfig() (cfg *ConfigApp, err error)
```
## func [NewDefaultConfigApp() ConfigApp](<config.go#L25>)

NewDefaultConfigApp returns a *ConfigApp with all the
default values filled.


```go
func NewDefaultConfigApp() *ConfigApp
```

## func (*ConfigApp) [CheckIsVersion()](<configApp.go#L28>)

```go
func (c *ConfigApp) CheckIsVersion()
```
## func (*ConfigApp) [IsVersion() bool](<configApp.go#L24>)

```go
func (c *ConfigApp) IsVersion() bool
```
## func (*ConfigApp) [SaveToFile(string) error](<configApp.go#L38>)

SaveToFile saves the config to a file in YAML format


```go
func (cfg *ConfigApp) SaveToFile(path string) error
```
## func (*ConfigApp) [SetupDefault()](<configApp.go#L50>)

SetupDefault sets up default config data.


```go
func (c *ConfigApp) SetupDefault()
```

## type [ConfigClient](<configClient.go#L14>)

ConfigApp holds all the application configuration data.
```go
type ConfigClient struct {
	Host	string	`env:"HOST" yaml:"host" json:"host" cli:"host FOS Host\nExample:\n  192.168.1.2\n  [::]:80\n  https://10.10.10.10/rest"`

	BaseURI			string	`env:"BASE_URI" yaml:"base_uri" json:"base_uri" cli:"base-uri FOS base URI for the REST API\nExample:\n  /rest\nDefault: /rest"`
	Username		string	`env:"USERNAME" yaml:"username" json:"username" cli:"username API login username"`
	Password		string	`env:"PASSWORD" yaml:"password" json:"password" cli:"password API login password\nDefaults to '-' as stdin"`
	Session			bool	`env:"SESSION" yaml:"session" json:"session" cli:"session Session authentication\nEstablishes a session instead of providing username:password with every request\nManual login and logout"`
	CheckPorts		bool	`env:"CHECK_PORTS" yaml:"check_ports" json:"check_ports" cli:"check-ports Check ports"`
	CheckPowerSupplies	bool	`env:"CHECK_POWER_SUPPLIES" yaml:"check_power_supplies" json:"check_poer_supplies" cli:"check-power-supplies Check power supplies"`
	CheckFans		bool	`env:"CHECK_FANS" yaml:"check_fans" json:"check_fans" cli:"check-fans Check fans"`
	// contains filtered or unexported fields
}
```

## func (*ConfigClient) [GetPasswordFromStdin() error](<configClient.go#L60>)

```go
func (c *ConfigClient) GetPasswordFromStdin() error
```
## func (*ConfigClient) [GetURL() string](<configClient.go#L26>)

```go
func (c *ConfigClient) GetURL() string
```
## func (*ConfigClient) [ParseURL() error](<configClient.go#L36>)

```go
func (c *ConfigClient) ParseURL() (err error)
```
## func (*ConfigClient) [SetupDefault()](<configClient.go#L31>)

SetupDefault sets up default config data.


```go
func (c *ConfigClient) SetupDefault()
```

## type [ConfigLog](<configLog.go#L5>)
```go
type ConfigLog struct {
	Type		LogType		`env:"GOSERVERHEADERS_LOG_TYPE" yaml:"type" cli:"log-type Logging output type\n      1: stdout\n      2: file (must have -log-file defined)\n      3: both\n     " default:"1"`
	Level		logrus.Level	`env:"GOSERVERHEADERS_LOG_LEVEL" yaml:"level" cli:"log-level Logging level\n  Logging level in the range from 0 to 6.\n      0: Panic only logs panic-level messages\n      1: Fatal logs fatal-level messages\n      2: Error logs error-level messages\n      3: Warn logs warning-level messages\n      4: Info logs information-level messages\n      5: Debug logs debugging-level messages\n      6: Trace logs trace-level messages\n     " default:"6"`
	FilePath	string		`env:"GOSERVERHEADERS_LOG_FILEPATH"  yaml:"filepath" cli:"log-file Path to the log file if using logging type 2 or 3\n   " default:"./server.log"`
	Format		string		`json:"format" yaml:"format" cli:"log-format Log entries formatting\n[ text | json ]\n " default:"text"`
	ErrFilePath	string
}
```

## func (*ConfigLog) [SetupDefault()](<configLog.go#L14>)

SetupDefault sets up default config data.


```go
func (c *ConfigLog) SetupDefault()
```

## type [LogLevel](<configLog.go#L31>)
```go
type LogLevel int
```

## type [LogType](<configLog.go#L22>)
```go
type LogType int
```

## func (LogType) [String() string](<configLog.go#L33>)

```go
func (l LogType) String() string
```

