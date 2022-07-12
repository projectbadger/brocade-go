
# time

## Index

- [type BrocadeTime](#type-brocadetime)
- [type ClockServer](#type-clockserver)
- [type NTPClockServer](#type-ntpclockserver)
- [type NTPClockServerKey](#type-ntpclockserverkey)
- [type TimeZone](#type-timezone)


## type [BrocadeTime](<brocadeTime.go#L5>)
```go
type BrocadeTime struct {
	XMLName xml.Name `json:"-" xml:"brocade-time"`
}
```

## type [ClockServer](<brocadeTime.go#L13>)
```go
type ClockServer struct {
	XMLName xml.Name `json:"-" xml:"clock-server"`
}
```

## type [NTPClockServer](<brocadeTime.go#L17>)
```go
type NTPClockServer struct {
	XMLName xml.Name `json:"-" xml:"ntp-clock-server"`
}
```

## type [NTPClockServerKey](<brocadeTime.go#L21>)
```go
type NTPClockServerKey struct {
	XMLName xml.Name `json:"-" xml:"ntp-clock-server-key"`
}
```

## type [TimeZone](<brocadeTime.go#L9>)
```go
type TimeZone struct {
	XMLName xml.Name `json:"-" xml:"time-zone"`
}
```

