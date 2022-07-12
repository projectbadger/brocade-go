package time

import "encoding/xml"

type BrocadeTime struct {
	XMLName xml.Name `json:"-" xml:"brocade-time"`
}

type TimeZone struct {
	XMLName xml.Name `json:"-" xml:"time-zone"`
}

type ClockServer struct {
	XMLName xml.Name `json:"-" xml:"clock-server"`
}

type NTPClockServer struct {
	XMLName xml.Name `json:"-" xml:"ntp-clock-server"`
}

type NTPClockServerKey struct {
	XMLName xml.Name `json:"-" xml:"ntp-clock-server-key"`
}
