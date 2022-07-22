// // Documentation source: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM
package media

import "encoding/xml"

// A detailed summary of the Small Form-factor Pluggable
// (SFP) transceivers media data for all available ports.
type BrocadeMedia struct {
	XMLName  xml.Name   `json:"-" xml:"brocade-media"`
	MediaRdp []MediaRDP `json:"media-rdp,omitempty" xml:"media-rdp"`
}

type MediaRDP struct {
	XMLName                         xml.Name                 `json:"-" xml:"media-rdp"`
	Name                            string                   `json:"name" xml:"name"`
	Identifier                      string                   `json:"identifier,omitempty" xml:"identifier"`
	Connector                       string                   `json:"connector,omitempty" xml:"connector"`
	MediaSpeedCapability            []string                 `json:"media-speed-capability,omitempty" xml:"media-speed-capability>speed"`
	MediaDistance                   []string                 `json:"media-distance,omitempty" xml:"media-distance"`
	Encoding                        string                   `json:"encoding,omitempty" xml:"encoding"`
	VendorOui                       string                   `json:"vendor-oui,omitempty" xml:"vendor-oui"`
	PartNumber                      string                   `json:"part-number,omitempty" xml:"part-number"`
	SerialNumber                    string                   `json:"serial-number,omitempty" xml:"serial-number"`
	VendorName                      string                   `json:"vendor-name,omitempty" xml:"vendor-name"`
	VendorRevision                  string                   `json:"vendor-revision,omitempty" xml:"vendor-revision"`
	DateCode                        string                   `json:"date-code,omitempty" xml:"date-code"`
	Temperature                     string                   `json:"temperature,omitempty" xml:"temperature"`
	RxPower                         string                   `json:"rx-power,omitempty" xml:"rx-power"`
	TxPower                         string                   `json:"tx-power,omitempty" xml:"tx-power"`
	Current                         string                   `json:"current,omitempty" xml:"current"`
	Voltage                         string                   `json:"voltage,omitempty" xml:"voltage"`
	Wavelength                      uint32                   `json:"wavelength,omitempty" xml:"wavelength"`
	PowerOnTime                     int32                    `json:"power-on-time,omitempty" xml:"power-on-time"`
	PeerDataAvailable               bool                     `json:"peer-data-available,omitempty" xml:"peer-data-available"`
	RemoteIdentifier                string                   `json:"remote-identifier,omitempty" xml:"remote-identifier"`
	RemoteLaserType                 string                   `json:"remote-laser-type,omitempty" xml:"remote-laser-type"`
	RemoteMediaSpeedCapability      []string                 `json:"remote-media-speed-capability" xml:"remote-media-speed-capability"`
	RemoteOpticalProductData        RemoteOpticalProductData `json:"remote-optical-product-data,omitempty" xml:"remote-optical-product-data"`
	RemoteMediaTemperature          string                   `json:"remote-media-temperature,omitempty" xml:"remote-media-temperature"`
	RemoteMediaRxPower              string                   `json:"remote-media-rx-power,omitempty" xml:"remote-media-rx-power"`
	RemoteMediaTxPower              string                   `json:"remote-media-tx-power,omitempty" xml:"remote-media-tx-power"`
	RemoteMediaCurrent              string                   `json:"remote-media-current,omitempty" xml:"remote-media-current"`
	RemoteMediaVoltage              string                   `json:"remote-media-voltage,omitempty" xml:"remote-media-voltage"`
	RemoteMediaVoltageAlert         AlarmType                `json:"remote-media-voltage-alert,omitempty" xml:"remote-media-voltage-alert"`
	RemoteMediaTemperatureAlert     AlarmType                `json:"remote-media-temperature-alert,omitempty" xml:"remote-media-temperature-alert"`
	RemoteMediaTxBiasAlert          AlarmType                `json:"remote-media-tx-bias-alert,omitempty" xml:"remote-media-tx-bias-alert"`
	RemoteMediaTxPowerAlert         AlarmType                `json:"remote-media-tx-power-alert,omitempty" xml:"remote-media-tx-power-alert"`
	RemoteMediaRxPowerAlert         AlarmType                `json:"remote-media-rx-power-alert,omitempty" xml:"remote-media-rx-power-alert"`
	RemoteMediaVoltageAlarmType     AlarmType                `json:"remote-media-voltage-alarm-type,omitempty" xml:"remote-media-voltage-alarm-type"`
	RemoteMediaTemperatureAlarmType AlarmType                `json:"remote-media-temperature-alarm-type,omitempty" xml:"remote-media-temperature-alarm-type"`
	RemoteMediaTxBiasAlarmType      AlarmType                `json:"remote-media-tx-bias-alarm-type,omitempty" xml:"remote-media-tx-bias-alarm-type"`
	RemoteMediaTxPowerAlarmType     AlarmType                `json:"remote-media-tx-power-alarm-type,omitempty" xml:"remote-media-tx-power-alarm-type"`
	RemoteMediaRxPowerAlarmType     AlarmType                `json:"remote-media-rx-power-alarm-type,omitempty" xml:"remote-media-rx-power-alarm-type"`
}

type RemoteOpticalProductData struct {
	XMLName        xml.Name `json:"-" xml:"remote-optical-product-data"`
	PartNumber     string   `json:"part-number,omitempty" xml:"part-number"`
	SerialNumber   string   `json:"serial-number,omitempty" xml:"serial-number"`
	VendorName     string   `json:"vendor-name,omitempty" xml:"vendor-name"`
	VendorRevision string   `json:"vendor-revision,omitempty" xml:"vendor-revision"`
	DateCode       string   `json:"date-code,omitempty" xml:"date-code"`
}

type AlarmType struct {
	HighAlarm   bool `json:"high-alarm,omitempty" xml:"high-alarm"`
	LowAlarm    bool `json:"low-alarm,omitempty" xml:"low-alarm"`
	HighWarning bool `json:"high-warning,omitempty" xml:"high-warning"`
	LowWarning  bool `json:"low-warning,omitempty" xml:"low-warning"`
}
