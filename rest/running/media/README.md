
# media

## Index

- [type AlarmType](#type-alarmtype)
- [type BrocadeMedia](#type-brocademedia)
- [type MediaRDP](#type-mediardp)
- [type RESTMedia](#type-restmedia)
  - [NewRESTMedia() RESTMedia](#func-newrestmedia-restmedia)
- [type RemoteOpticalProductData](#type-remoteopticalproductdata)


## type [AlarmType](<brocadeMedia.go#L61>)
```go
type AlarmType struct {
	HighAlarm	bool	`json:"high-alarm,omitempty" xml:"high-alarm"`
	LowAlarm	bool	`json:"low-alarm,omitempty" xml:"low-alarm"`
	HighWarning	bool	`json:"high-warning,omitempty" xml:"high-warning"`
	LowWarning	bool	`json:"low-warning,omitempty" xml:"low-warning"`
}
```

## type [BrocadeMedia](<brocadeMedia.go#L5>)
```go
type BrocadeMedia struct {
	XMLName xml.Name `json:"-" xml:"brocade-media"`
}
```

## type [MediaRDP](<brocadeMedia.go#L9>)
```go
type MediaRDP struct {
	XMLName				xml.Name			`json:"-" xml:"media-rdp"`
	Name				string				`json:"name" xml:"name"`
	Identifier			string				`json:"identifier,omitempty" xml:"identifier"`
	Connector			string				`json:"connector,omitempty" xml:"connector"`
	MediaSpeedCapability		[]string			`json:"media-speed-capability,omitempty" xml:"media-speed-capability>speed"`
	MediaDistance			[]string			`json:"media-distance,omitempty" xml:"media-distance"`
	Encoding			string				`json:"encoding,omitempty" xml:"encoding"`
	VendorOui			string				`json:"vendor-oui,omitempty" xml:"vendor-oui"`
	PartNumber			string				`json:"part-number,omitempty" xml:"part-number"`
	SerialNumber			string				`json:"serial-number,omitempty" xml:"serial-number"`
	VendorName			string				`json:"vendor-name,omitempty" xml:"vendor-name"`
	VendorRevision			string				`json:"vendor-revision,omitempty" xml:"vendor-revision"`
	DateCode			string				`json:"date-code,omitempty" xml:"date-code"`
	Temperature			string				`json:"temperature,omitempty" xml:"temperature"`
	RxPower				string				`json:"rx-power,omitempty" xml:"rx-power"`
	TxPower				string				`json:"tx-power,omitempty" xml:"tx-power"`
	Current				string				`json:"current,omitempty" xml:"current"`
	Voltage				string				`json:"voltage,omitempty" xml:"voltage"`
	Wavelength			uint32				`json:"wavelength,omitempty" xml:"wavelength"`
	PowerOnTime			int32				`json:"power-on-time,omitempty" xml:"power-on-time"`
	PeerDataAvailable		bool				`json:"peer-data-available,omitempty" xml:"peer-data-available"`
	RemoteIdentifier		string				`json:"remote-identifier,omitempty" xml:"remote-identifier"`
	RemoteLaserType			string				`json:"remote-laser-type,omitempty" xml:"remote-laser-type"`
	RemoteMediaSpeedCapability	[]string			`json:"remote-media-speed-capability" xml:"remote-media-speed-capability"`
	RemoteOpticalProductData	RemoteOpticalProductData	`json:"remote-optical-product-data,omitempty" xml:"remote-optical-product-data"`
	RemoteMediaTemperature		string				`json:"remote-media-temperature,omitempty" xml:"remote-media-temperature"`
	RemoteMediaRxPower		string				`json:"remote-media-rx-power,omitempty" xml:"remote-media-rx-power"`
	RemoteMediaTxPower		string				`json:"remote-media-tx-power,omitempty" xml:"remote-media-tx-power"`
	RemoteMediaCurrent		string				`json:"remote-media-current,omitempty" xml:"remote-media-current"`
	RemoteMediaVoltage		string				`json:"remote-media-voltage,omitempty" xml:"remote-media-voltage"`
	RemoteMediaVoltageAlert		AlarmType			`json:"remote-media-voltage-alert,omitempty" xml:"remote-media-voltage-alert"`
	RemoteMediaTemperatureAlert	AlarmType			`json:"remote-media-temperature-alert,omitempty" xml:"remote-media-temperature-alert"`
	RemoteMediaTxBiasAlert		AlarmType			`json:"remote-media-tx-bias-alert,omitempty" xml:"remote-media-tx-bias-alert"`
	RemoteMediaTxPowerAlert		AlarmType			`json:"remote-media-tx-power-alert,omitempty" xml:"remote-media-tx-power-alert"`
	RemoteMediaRxPowerAlert		AlarmType			`json:"remote-media-rx-power-alert,omitempty" xml:"remote-media-rx-power-alert"`
	RemoteMediaVoltageAlarmType	AlarmType			`json:"remote-media-voltage-alarm-type,omitempty" xml:"remote-media-voltage-alarm-type"`
	RemoteMediaTemperatureAlarmType	AlarmType			`json:"remote-media-temperature-alarm-type,omitempty" xml:"remote-media-temperature-alarm-type"`
	RemoteMediaTxBiasAlarmType	AlarmType			`json:"remote-media-tx-bias-alarm-type,omitempty" xml:"remote-media-tx-bias-alarm-type"`
	RemoteMediaTxPowerAlarmType	AlarmType			`json:"remote-media-tx-power-alarm-type,omitempty" xml:"remote-media-tx-power-alarm-type"`
	RemoteMediaRxPowerAlarmType	AlarmType			`json:"remote-media-rx-power-alarm-type,omitempty" xml:"remote-media-rx-power-alarm-type"`
}
```

## type [RESTMedia](<methods.go#L12>)

RESTMedia describes an interface for interacting with the
*media* module.
Fetch a new instance using the NewRESTMedia function.
```go
type RESTMedia interface {
	Name() string
	URIPath() string
	GetMediaRDP(name string) ([]MediaRDP, error)
	GetMediaRDPResponse(name string) (*http.Response, error)
}
```

## func [NewRESTMedia() RESTMedia](<methods.go#L24>)

```go
func NewRESTMedia(cfg *api_interface.RESTConfig) RESTMedia
```

## type [RemoteOpticalProductData](<brocadeMedia.go#L52>)
```go
type RemoteOpticalProductData struct {
	XMLName		xml.Name	`json:"-" xml:"remote-optical-product-data"`
	PartNumber	string		`json:"part-number,omitempty" xml:"part-number"`
	SerialNumber	string		`json:"serial-number,omitempty" xml:"serial-number"`
	VendorName	string		`json:"vendor-name,omitempty" xml:"vendor-name"`
	VendorRevision	string		`json:"vendor-revision,omitempty" xml:"vendor-revision"`
	DateCode	string		`json:"date-code,omitempty" xml:"date-code"`
}
```

