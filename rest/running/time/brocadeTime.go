package time

import "encoding/xml"

type BrocadeTime struct {
	XMLName           xml.Name          `json:"-" xml:"brocade-time"`
	TimeZone          TimeZone          `json:"time-zone" xml:"time-zone"`
	ClockServer       ClockServer       `json:"clock-server" xml:"clock-server"`
	NTPClockServer    NTPClockServer    `json:"ntp-clock-server,omitempty" xml:"ntp-clock-server"`
	NTPClockServerKey NTPClockServerKey `json:"ntp-clock-server-key,omitempty" xml:"ntp-clock-server-key"`
}

type TimeZone struct {
	XMLName          xml.Name `json:"-" xml:"time-zone"`
	Name             string   `json:"name,omitempty" xml:"name"`
	GMTOffsetHours   int16    `json:"gmt-offset-hours,omitempty" xml:"gmt-offset-hours"`
	GMTOffsetMinutes int16    `json:"gmt-offset-minutes,omitempty" xml:"gmt-offset-minutes"`
}

type ClockServer struct {
	XMLName          xml.Name `json:"-" xml:"clock-server"`
	NTPServerAddress []string `json:"ntp-server-address" xml:"ntp-server-address>server-address"`
	ActiveServer     string   `json:"active-server,omitempty" xml:"active-server"`
	TSAuthType       string   `json:"ts-auth-type,omitempty" xml:"ts-auth-type"`
	TSLegacyMode     bool     `json:"ts-legacy-mode,omitempty" xml:"ts-legacy-mode"`
}

type NTPClockServer struct {
	XMLName xml.Name `json:"-" xml:"ntp-clock-server"`
	Server  string   `json:"server,omitempty" xml:"server"`
	Index   int32    `json:"index,omitempty" xml:"index"`
}

type NTPClockServerKey struct {
	XMLName xml.Name `json:"-" xml:"ntp-clock-server-key"`
	Index   int32    `json:"index" xml:"index"`
	Type    string   `json:"type,omitempty" xml:"type"`
	Key     string   `json:"key,omitempty" xml:"key"`
}
