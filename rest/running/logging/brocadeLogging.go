// Documentation source: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM
package logging

import "encoding/xml"

// BrocadeLogging provides a detailed view of audit and RASlog message configuration.
type BrocadeLogging struct {
	XMLName         xml.Name          `json:"-" xml:"brocade-logging"`
	Audit           Audit             `json:"audit,omitempty" xml:"audit"`
	SyslogServer    []SyslogServer    `json:"syslog-server,omitempty" xml:"syslog-server"`
	RASlog          []RASLog          `json:"raslog,omitempty" xml:"raslog"`
	RaslogModule    []RASLogModule    `json:"raslog-module,omitempty" xml:"raslog-module"`
	LogQuietControl []LogQuietControl `json:"log-quiet-control,omitempty" xml:"log-quiet-control"`
	LogSettings     LogSettings       `json:"log-settings,omitempty" xml:"log-settings"`
}

// The audit logging configuration parameters.
// You can configure certain filter classes, to set
// severity levels for audit messages, and enable or
// disable audit filters. Depending on the configuration,
// certain classes are logged to syslog for auditing. Note
// that syslog configuration is required for logging
// audit messages.
type Audit struct {
	XMLName xml.Name `json:"-" xml:"audit"`
	// Enables or disables the audit filters.
	AuditEnabled bool `json:"audit-enabled,omitempty" xml:"audit-enabled"`
	// The severity level of the log messages you want to
	// display.
	//  - info: Displays log messages of info level and higher.
	//  - warning: Displays log messages of warning level and higher.
	//  - error: Displays log messages of error level and higher.
	//  - critical: Displays log messages of critical level and higher.
	SeverityLevel string `json:"severity-level,omitempty" xml:"severity-level"`
	// A list of class types needed for audit configuration.
	// You can set more than one audit class, separated by
	// a comma.
	//
	//  - zone: Displays the zone audit class.
	//  - security: Displays the security audit class.
	//  - configuration: Displays the configuration audit class
	//  - firmware: Displays the firmware audit class
	//  - fabric: Displays the fabric audit class
	//  - ls: Displays the ls audit class
	//  - cli: Displays the cli audit class
	//  - maps: Displays the maps audit class
	FilterClassList []string `json:"filter-class-list,omitempty" xml:"filter-class-list>filter-class"`
}

// The remote syslog server.
type SyslogServer struct {
	XMLName xml.Name `json:"-" xml:"syslog-server"`
	// The IPv4 or IPv6 address or DNS name of the server.
	Server string `json:"server,omitempty" xml:"server"`
	// The target syslog server's TCP/IP port number.
	// The port can only be specified (non-default)
	// when secure mode is enabled on the server
	// (secure-mode=true).
	Port uint `json:"port,omitempty" xml:"port"`
	// Whether secure syslog mode to send the error log
	// messages securely using the TLS protocol to the
	// syslog server is enabled.
	SecureMode bool `json:"secure-mode,omitempty" xml:"secure-mode"`
}

// A list of RASlog messages’ configurable parameters, such
// as enabling message logging, the message ID, the current
// severity level, transferring to the syslog server is
// enabled. This list also includes read-only leafs, such
// as the default severity, the message text, and the
// flooded flag.
type RASLog struct {
	XMLName xml.Name `json:"-" xml:"raslog"`
	// The unique identifier for each RASlog message.
	MessageId string `json:"message-id,omitempty" xml:"message-id"`
	// Whether message logging is enabled for RASlog
	// messages identified by message ID.
	MessageEnabled bool `json:"message-enabled,omitempty" xml:"message-enabled"`
	// Whether the RASlog message is blocked due to
	// flooding.
	MessageFlooded bool `json:"message-flooded,omitempty" xml:"message-flooded"`
	// Whether internal RASlog messages to be sent to
	// syslog is enabled.
	// This parameter is only supported for internal RASlog
	// messages.
	// Internal RASlog messages have an ID number equal to
	// or greater than 5000.
	SyslogEnabled bool `json:"syslog-enabled,omitempty" xml:"syslog-enabled"`
	// The message text corresponding to the internal
	// RASlog message.
	// This parameter is only supported for internal RASlog
	// messages.
	// Internal RASlog messages
	// have an ID number equal to or greater than 5000.
	MessageText string `json:"message-text,omitempty" xml:"message-text"`
	// The current severity level of the RASlog.
	// The default severity level sets severity to a
	// pre-defined level of logging for the RASlog
	// associated with the corresponding message ID.
	// However, you can configure a different severity with
	// this parameter.
	//
	// Valid values are info, warning, error, critical, and
	// default.
	CurrentSeverity string `json:"current-severity,omitempty" xml:"current-severity"`
	// The pre-defined default severity level of the RASlog.
	//
	// Values displayed are info, warning, error, and
	// critical.
	DefaultSeverity string `json:"default-severity,omitempty" xml:"default-severity"`
}

// A list of Fabric OS modules (identified by module ID)
// that have logging enabled.
type RASLogModule struct {
	XMLName xml.Name `json:"-" xml:"raslog-module"`
	// The Fabric OS module identified by module ID.
	ModuleId string `json:"module-id,omitempty" xml:"module-id"`
	// Whether logging is enabled for the Fabric OS module.
	// This parameter is used only during a PATCH request.
	// Note that this parameter does not display when you
	// use a GET request to display RASlog modules (for
	// example: GET https://10.10.10.10/rest/running/brocade-logging/raslog-module)
	LogEnabled bool `json:"log-enabled,omitempty" xml:"log-enabled"`
}

// A list of parameters that control quiet time for logging
// based on the log type.
type LogQuietControl struct {
	XMLName xml.Name `json:"-" xml:"log-quiet-control"`
	// The log type for which quiet time is configured.
	LogType string `json:"log-type,omitempty" xml:"log-type"`
	// Whether quiet time is enabled for logs.
	QuietEnabled bool `json:"quiet-enabled,omitempty" xml:"quiet-enabled"`
	// The RAS quiet start time for the specified log type.
	// This parameter is available only when quiet time is
	// enabled (quiet-enabled = true).
	StartTime string `json:"start-time,omitempty" xml:"start-time"`
	// e RAS quiet end time for the specified log type.
	// Start time must be configured before you can set the
	// end time. This parameter is available only when
	// quiet time is enabled (quiet-enabled = true).
	EndTime string `json:"end-time,omitempty" xml:"end-time"`
	// The days of the week for which you want to set quiet
	// time for the specified log type.
	//
	// Values:
	//  - everyday: Enables quiet time everyday for the specified start time and end time
	//  - forever: Enables quiet time every day all day (you cannot enter a start or end time).
	//  - mon, tue, wed, thu, fri, sat, or sun: Enables quiet time for the specified days of the week for the specified start time and end time.
	DaysOfWeek []string `json:"days-of-week,omitempty" xml:"days-of-week>day"`
}

// The system-wide parameters to control the logs.
type LogSettings struct {
	XMLName xml.Name `json:"-" xml:"log-settings"`
	// The facility level determines the priority of the
	// syslog messages being recorded at the server.
	// A smaller facility level corresponds to higher
	// priority syslog messages.
	//
	// Valid values are log_local0, log_local1, log_local2,
	// log_local3, log_local4, log_local5, log_local6, and
	// log_local7 (default).
	SyslogFacility string `json:"syslog-facility,omitempty" xml:"syslog-facility"`
	// The RASlog keep alive timeout.
	//
	// 0 through 24: Keeps logging alive for the specified
	// time. 0: Disables the keep alive timeout.
	// Default = 1
	KeepAlivePeriod uint16 `json:"keep-alive-period,omitempty" xml:"keep-alive-period"`
}
