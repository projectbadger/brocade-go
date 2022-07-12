
# fibrechannel_diagnostics

## Index

- [type BrocadeFibrechannelDiagnostics](#type-brocadefibrechanneldiagnostics)
- [type CR](#type-cr)
- [type EgressPowerLoss](#type-egresspowerloss)
- [type ElectricalLoopbackTest](#type-electricalloopbacktest)
- [type FEC](#type-fec)
- [type FailureReport](#type-failurereport)
- [type FailureReportLocalErrors](#type-failurereportlocalerrors)
- [type FailureReportRemoteErrors](#type-failurereportremoteerrors)
- [type FibrechannelDiagnostics](#type-fibrechanneldiagnostics)
- [type IngressPowerLoss](#type-ingresspowerloss)
- [type LinkTrafficTest](#type-linktraffictest)
- [type OpticalLoopbackTest](#type-opticalloopbacktest)
- [type PayloadPattern](#type-payloadpattern)


## type [BrocadeFibrechannelDiagnostics](<brocadeFibrechannelDiagnostics.go#L6>)

Documentation source: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM - page 138 (20220701)
```go
type BrocadeFibrechannelDiagnostics struct {
	XMLName xml.Name `json:"-" xml:"brocade-fibrechannel-diagnostics"`
}
```

## type [CR](<brocadeFibrechannelDiagnostics.go#L77>)
```go
type CR struct {
	XMLName	xml.Name	`json:"-" xml:"cr"`
	Enable	string		`json:"enable,omitempty" xml:"enable"`
	Active	string		`json:"active,omitempty" xml:"active"`
	Option	string		`json:"option,omitempty" xml:"option"`
}
```

## type [EgressPowerLoss](<brocadeFibrechannelDiagnostics.go#L100>)
```go
type EgressPowerLoss struct {
	XMLName		xml.Name	`json:"-" xml:"egress-power-loss"`
	Tx		float64		`json:"tx,omitempty" xml:"tx"`
	Rx		float64		`json:"rx,omitempty" xml:"rx"`
	Loss		float64		`json:"loss,omitempty" xml:"loss"`
	Comments	string		`json:"comments,omitempty" xml:"comments"`
}
```

## type [ElectricalLoopbackTest](<brocadeFibrechannelDiagnostics.go#L40>)
```go
type ElectricalLoopbackTest struct {
	XMLName		xml.Name	`json:"-" xml:"electrical-loopback-test"`
	Result		string		`json:"result,omitempty" xml:"result"`
	Comments	string		`json:"comments,omitempty" xml:"comments"`
	StartTime	string		`json:"start-time,omitempty" xml:"start-time"`
	EstimateTime	string		`json:"estimate-time,omitempty" xml:"estimate-time"`
}
```

## type [FEC](<brocadeFibrechannelDiagnostics.go#L70>)
```go
type FEC struct {
	XMLName	xml.Name	`json:"-" xml:"fec"`
	Enable	string		`json:"enable,omitempty" xml:"enable"`
	Active	string		`json:"active,omitempty" xml:"active"`
	Option	string		`json:"option,omitempty" xml:"option"`
}
```

## type [FailureReport](<brocadeFibrechannelDiagnostics.go#L84>)
```go
type FailureReport struct {
	XMLName			xml.Name	`json:"-" xml:"failure-report"`
	ErrorsDetectedLocal	string		`json:"errors-detected-local,omitempty" xml:"errors-detected-local"`
	ErrorsDetectedRemote	string		`json:"errors-detected-remote,omitempty" xml:"errors-detected-remote"`
}
```

## type [FailureReportLocalErrors](<brocadeFibrechannelDiagnostics.go#L90>)
```go
type FailureReportLocalErrors struct {
	XMLName	xml.Name	`json:"-" xml:"failure-report-local-errors"`
	Error	string		`json:"error,omitempty" xml:"error"`
}
```

## type [FailureReportRemoteErrors](<brocadeFibrechannelDiagnostics.go#L95>)
```go
type FailureReportRemoteErrors struct {
	XMLName	xml.Name	`json:"-" xml:"failure-report-remote-errors"`
	Error	string		`json:"error,omitempty" xml:"error"`
}
```

## type [FibrechannelDiagnostics](<brocadeFibrechannelDiagnostics.go#L10>)
```go
type FibrechannelDiagnostics struct {
	XMLName				xml.Name			`json:"-" xml:"fibrechannel-diagnostics"`
	Name				string				`json:"name" xml:"name"`
	DiagnosticControl		uint32				`json:"diagnostic-control,omitempty" xml:"diagnostic-control"`
	Mode				string				`json:"mode,omitempty" xml:"mode"`	// enum
	State				string				`json:"state" xml:"state"`
	ErrorMessage			string				`json:"error-message" xml:"error-message"`
	Distance			int32				`json:"distance" xml:"distance"`
	RemoteSwitchWWN			string				`json:"remote-switch-wwn" xml:"remote-switch-wwn"`
	RemotePortIndex			string				`json:"remote-port-index" xml:"remote-port-index"`
	ElectricalLoopbackTest		ElectricalLoopbackTest		`json:"electrical-loopback-test" xml:"electrical-loopback-test"`
	OpticalLoopbackTest		OpticalLoopbackTest		`json:"optical-loopback-test" xml:"optical-loopback-test"`
	LinkTrafficTest			LinkTrafficTest			`json:"link-traffic-test" xml:"link-traffic-test"`
	StartTime			string				`json:"start-time" xml:"start-time"`
	FrameCount			uint32				`json:"frame-count" xml:"frame-count"`
	FrameSize			uint32				`json:"frame-size" xml:"frame-size"`
	Time				string				`json:"time" xml:"time"`
	PayloadPattern			PayloadPattern			`json:"payload-pattern" xml:"payload-pattern"`
	FEC				FEC				`json:"fec" xml:"fec"`
	RTLatency			uint32				`json:"rt-latency,omitempty" xml:"rt-latency"`
	BuffersRequired			string				`json:"buffers-required" xml:"buffers-required"`
	EndTime				string				`json:"end-time" xml:"end-time"`
	CR				CR				`json:"cr" xml:"cr"`
	FailureReport			FailureReport			`json:"failure-report" xml:"failure-report"`
	FailureReportLocalErrors	FailureReportLocalErrors	`json:"failure-report-local-errors" xml:"failure-report-local-errors"`
	FailureReportRemoteErrors	FailureReportRemoteErrors	`json:"failure-report-remote-errors" xml:"failure-report-remote-errors"`
	EgressPowerLoss			EgressPowerLoss			`json:"egress-power-loss" xml:"egress-power-loss"`
	IngressPowerLoss		IngressPowerLoss		`json:"ingress-power-loss" xml:"ingress-power-loss"`
}
```

## type [IngressPowerLoss](<brocadeFibrechannelDiagnostics.go#L108>)
```go
type IngressPowerLoss struct {
	XMLName		xml.Name	`json:"-" xml:"ingress-power-loss"`
	Tx		float64		`json:"tx,omitempty" xml:"tx"`
	Rx		float64		`json:"rx,omitempty" xml:"rx"`
	Loss		float64		`json:"loss,omitempty" xml:"loss"`
	Comments	string		`json:"comments,omitempty" xml:"comments"`
}
```

## type [LinkTrafficTest](<brocadeFibrechannelDiagnostics.go#L56>)
```go
type LinkTrafficTest struct {
	XMLName		xml.Name	`json:"-" xml:"link-traffic-test"`
	Result		string		`json:"result,omitempty" xml:"result"`
	Comments	string		`json:"comments,omitempty" xml:"comments"`
	StartTime	string		`json:"start-time,omitempty" xml:"start-time"`
	EstimateTime	string		`json:"estimate-time,omitempty" xml:"estimate-time"`
}
```

## type [OpticalLoopbackTest](<brocadeFibrechannelDiagnostics.go#L48>)
```go
type OpticalLoopbackTest struct {
	XMLName		xml.Name	`json:"-" xml:"optical-loopback-test"`
	Result		string		`json:"result,omitempty" xml:"result"`
	Comments	string		`json:"comments,omitempty" xml:"comments"`
	StartTime	string		`json:"start-time,omitempty" xml:"start-time"`
	EstimateTime	string		`json:"estimate-time,omitempty" xml:"estimate-time"`
}
```

## type [PayloadPattern](<brocadeFibrechannelDiagnostics.go#L64>)
```go
type PayloadPattern struct {
	XMLName	xml.Name	`json:"-" xml:"payload-pattern"`
	Pattern	string		`json:"pattern,omitempty" xml:"pattern"`
	Payload	string		`json:"payload,omitempty" xml:"payload"`
}
```

