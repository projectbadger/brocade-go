
# logging

## Index

- [type Audit](#type-audit)
- [type BrocadeLogging](#type-brocadelogging)
- [type LogQuietControl](#type-logquietcontrol)
- [type LogSettings](#type-logsettings)
- [type RASLog](#type-raslog)
- [type RASLogModule](#type-raslogmodule)
- [type RESTLogging](#type-restlogging)
  - [NewRESTLogging() RESTLogging](#func-newrestlogging-restlogging)
- [type SyslogServer](#type-syslogserver)


## type [Audit](<brocadeLogging.go#L15>)
```go
type Audit struct {
	XMLName		xml.Name	`json:"-" xml:"audit"`
	AuditEnabled	bool		`json:"audit-enabled,omitempty" xml:"audit-enabled"`
	SeverityLevel	string		`json:"severity-level,omitempty" xml:"severity-level"`
	FilterClassList	[]string	`json:"filter-class-list,omitempty" xml:"filter-class-list>filter-class"`
}
```

## type [BrocadeLogging](<brocadeLogging.go#L5>)
```go
type BrocadeLogging struct {
	XMLName		xml.Name		`json:"-" xml:"brocade-logging"`
	Audit		Audit			`json:"audit,omitempty" xml:"audit"`
	SyslogServer	[]SyslogServer		`json:"syslog-server,omitempty" xml:"syslog-server"`
	RASlog		[]RASLog		`json:"raslog,omitempty" xml:"raslog"`
	RaslogModule	[]RASLogModule		`json:"raslog-module,omitempty" xml:"raslog-module"`
	LogQuietControl	[]LogQuietControl	`json:"log-quiet-control,omitempty" xml:"log-quiet-control"`
	LogSettings	LogSettings		`json:"log-settings,omitempty" xml:"log-settings"`
}
```

## type [LogQuietControl](<brocadeLogging.go#L46>)
```go
type LogQuietControl struct {
	XMLName		xml.Name	`json:"-" xml:"log-quiet-control"`
	LogType		string		`json:"log-type,omitempty" xml:"log-type"`
	QuietEnabled	bool		`json:"quiet-enabled,omitempty" xml:"quiet-enabled"`
	StartTime	string		`json:"start-time,omitempty" xml:"start-time"`
	EndTime		string		`json:"end-time,omitempty" xml:"end-time"`
	DaysOfWeek	[]string	`json:"days-of-week,omitempty" xml:"days-of-week>day"`
}
```

## type [LogSettings](<brocadeLogging.go#L55>)
```go
type LogSettings struct {
	XMLName		xml.Name	`json:"-" xml:"log-settings"`
	SyslogFacility	string		`json:"syslog-facility,omitempty" xml:"syslog-facility"`
	KeepAlivePeriod	uint16		`json:"keep-alive-period,omitempty" xml:"keep-alive-period"`
}
```

## type [RASLog](<brocadeLogging.go#L29>)
```go
type RASLog struct {
	XMLName		xml.Name	`json:"-" xml:"raslog"`
	MessageId	string		`json:"message-id,omitempty" xml:"message-id"`
	MessageEnabled	bool		`json:"message-enabled,omitempty" xml:"message-enabled"`
	MessageFlooded	bool		`json:"message-flooded,omitempty" xml:"message-flooded"`
	SyslogEnabled	bool		`json:"syslog-enabled,omitempty" xml:"syslog-enabled"`
	MessageText	string		`json:"message-text,omitempty" xml:"message-text"`
	CurrentSeverity	string		`json:"current-severity,omitempty" xml:"current-severity"`
	DefaultSeverity	string		`json:"default-severity,omitempty" xml:"default-severity"`
}
```

## type [RASLogModule](<brocadeLogging.go#L40>)
```go
type RASLogModule struct {
	XMLName		xml.Name	`json:"-" xml:"raslog-module"`
	ModuleId	string		`json:"module-id,omitempty" xml:"module-id"`
	LogEnabled	bool		`json:"log-enabled,omitempty" xml:"log-enabled"`
}
```

## type [RESTLogging](<methods.go#L10>)
```go
type RESTLogging interface {
	Name() string
	URIPath() string
	GetLogging() (*BrocadeLogging, errors.BrocadeErr)
	GetAudit() (*Audit, errors.BrocadeErr)
	GetSyslogServer() ([]SyslogServer, errors.BrocadeErr)
	GetRASLog() ([]RASLog, errors.BrocadeErr)
	GetLogQuietControl() ([]LogQuietControl, errors.BrocadeErr)
	GetRASLogModule() ([]RASLogModule, errors.BrocadeErr)
	GetLogSetting() (*LogSettings, errors.BrocadeErr)
}
```

## func [NewRESTLogging() RESTLogging](<methods.go#L30>)

```go
func NewRESTLogging(config *api_interface.RESTConfig) RESTLogging
```

## type [SyslogServer](<brocadeLogging.go#L22>)
```go
type SyslogServer struct {
	XMLName		xml.Name	`json:"-" xml:"syslog-server"`
	Server		string		`json:"server,omitempty" xml:"server"`
	Port		uint		`json:"port,omitempty" xml:"port"`
	SecureMode	bool		`json:"secure-mode,omitempty" xml:"secure-mode"`
}
```

