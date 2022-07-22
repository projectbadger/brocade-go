
# time

## Index

- [type BrocadeTime](#type-brocadetime)
- [type ClockServer](#type-clockserver)
- [type NTPClockServer](#type-ntpclockserver)
- [type NTPClockServerKey](#type-ntpclockserverkey)
- [type RESTTime](#type-resttime)
  - [NewRESTTime() RESTTime](#func-newresttime-resttime)
- [type TimeZone](#type-timezone)


## type [BrocadeTime](<brocadeTime.go#L5>)
```go
type BrocadeTime struct {
	XMLName			xml.Name		`json:"-" xml:"brocade-time"`
	TimeZone		TimeZone		`json:"time-zone" xml:"time-zone"`
	ClockServer		ClockServer		`json:"clock-server" xml:"clock-server"`
	NTPClockServer		NTPClockServer		`json:"ntp-clock-server,omitempty" xml:"ntp-clock-server"`
	NTPClockServerKey	NTPClockServerKey	`json:"ntp-clock-server-key,omitempty" xml:"ntp-clock-server-key"`
}
```

## type [ClockServer](<brocadeTime.go#L20>)
```go
type ClockServer struct {
	XMLName			xml.Name	`json:"-" xml:"clock-server"`
	NTPServerAddress	[]string	`json:"ntp-server-address" xml:"ntp-server-address>server-address"`
	ActiveServer		string		`json:"active-server,omitempty" xml:"active-server"`
	TSAuthType		string		`json:"ts-auth-type,omitempty" xml:"ts-auth-type"`
	TSLegacyMode		bool		`json:"ts-legacy-mode,omitempty" xml:"ts-legacy-mode"`
}
```

## type [NTPClockServer](<brocadeTime.go#L28>)
```go
type NTPClockServer struct {
	XMLName	xml.Name	`json:"-" xml:"ntp-clock-server"`
	Server	string		`json:"server,omitempty" xml:"server"`
	Index	int32		`json:"index,omitempty" xml:"index"`
}
```

## type [NTPClockServerKey](<brocadeTime.go#L34>)
```go
type NTPClockServerKey struct {
	XMLName	xml.Name	`json:"-" xml:"ntp-clock-server-key"`
	Index	int32		`json:"index" xml:"index"`
	Type	string		`json:"type,omitempty" xml:"type"`
	Key	string		`json:"key,omitempty" xml:"key"`
}
```

## type [RESTTime](<methods.go#L13>)

RESTTime describes an interface for interacting with the
*time* module.
Fetch a new instance using the NewRESTTime function.
```go
type RESTTime interface {
	Name() string
	URIPath() string
	GetTimeZone() (*TimeZone, error)
	GetTimeZoneResponse() (*http.Response, error)
	GetClockServer() (*ClockServer, error)
	GetClockServerResponse() (*http.Response, error)
	GetNTPClockServer(server string) (*NTPClockServer, error)
	GetNTPClockServerResponse(server string) (*http.Response, error)
	GetNTPClockServerKey(index int32) (*NTPClockServerKey, error)
	GetNTPClockServerKeyResponse(index int32) (*http.Response, error)
}
```

## func [NewRESTTime() RESTTime](<methods.go#L32>)

```go
func NewRESTTime(cfg *api_interface.RESTConfig) RESTTime
```

## type [TimeZone](<brocadeTime.go#L13>)
```go
type TimeZone struct {
	XMLName			xml.Name	`json:"-" xml:"time-zone"`
	Name			string		`json:"name,omitempty" xml:"name"`
	GMTOffsetHours		int16		`json:"gmt-offset-hours,omitempty" xml:"gmt-offset-hours"`
	GMTOffsetMinutes	int16		`json:"gmt-offset-minutes,omitempty" xml:"gmt-offset-minutes"`
}
```

