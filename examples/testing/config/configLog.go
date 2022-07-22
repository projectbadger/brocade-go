package config

import "github.com/sirupsen/logrus"

type ConfigLog struct {
	Type        LogType      `env:"GOSERVERHEADERS_LOG_TYPE" yaml:"type" cli:"log-type Logging output type\n      1: stdout\n      2: file (must have -log-file defined)\n      3: both\n     " default:"1"`
	Level       logrus.Level `env:"GOSERVERHEADERS_LOG_LEVEL" yaml:"level" cli:"log-level Logging level\n  Logging level in the range from 0 to 6.\n      0: Panic only logs panic-level messages\n      1: Fatal logs fatal-level messages\n      2: Error logs error-level messages\n      3: Warn logs warning-level messages\n      4: Info logs information-level messages\n      5: Debug logs debugging-level messages\n      6: Trace logs trace-level messages\n     " default:"6"`
	FilePath    string       `env:"GOSERVERHEADERS_LOG_FILEPATH"  yaml:"filepath" cli:"log-file Path to the log file if using logging type 2 or 3\n   " default:"./server.log"`
	Format      string       `json:"format" yaml:"format" cli:"log-format Log entries formatting\n[ text | json ]\n " default:"text"`
	ErrFilePath string
}

// SetupDefault sets up default config data.
func (c *ConfigLog) SetupDefault() {
	// c.Level = logrus.ErrorLevel
	c.Level = logrus.DebugLevel
	c.Type = LogTypeStdOut
	c.Format = "json"
	c.FilePath = "./brocade.log"
}

type LogType int

var logTypeStrings = []string{
	"",
	"stdout",
	"file",
	"stdout AND file",
}

type LogLevel int

func (l LogType) String() string {
	if l > 0 && l <= 3 {
		return logTypeStrings[l]
	}
	return logTypeStrings[1]
}

const (
	LogTypeStdOut LogType = 1 << iota
	LogTypeFile
	LogTypeMask = (1 << iota) - 1
)
