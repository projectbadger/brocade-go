
# fibrechannel_trunk

## Index

- [type BrocadeFibrechannelTrunk](#type-brocadefibrechanneltrunk)
- [type Performance](#type-performance)
- [type RESTFibrechannelTrunk](#type-restfibrechanneltrunk)
  - [NewRESTFibrechannelTrunk() RESTFibrechannelTrunk](#func-newrestfibrechanneltrunk-restfibrechanneltrunk)
- [type Trunk](#type-trunk)
- [type TrunkArea](#type-trunkarea)


## type [BrocadeFibrechannelTrunk](<brocadeFibrechannelTrunk.go#L15>)

This module provides a detailed view of all trunks in
the switch in native mode as well as the members of the
individual trunks. It can also provide traffic
performance and bandwidth information. For F_Port static
trunks, you can configure a trunk index that is
persistent and has the same value across all members of
the trunk. Even if the master port changes, the trunk
area is derived from the target index.

Documentation source: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM - page 158 (20220701)
```go
type BrocadeFibrechannelTrunk struct {
	XMLName		xml.Name	`json:"-" xml:"brocade-fibrechannel-trunk"`
	Trunk		Trunk		`json:"trunk,omitempty" xml:"trunk"`
	Performance	Performance	`json:"performance,omitempty" xml:"performance"`
	TrunkArea	TrunkArea	`json:"trunk-area,omitempty" xml:"trunk-area"`
}
```

## type [Performance](<brocadeFibrechannelTrunk.go#L34>)
```go
type Performance struct {
	XMLName		xml.Name	`json:"-" xml:"performance"`
	Group		string		`json:"group,omitempty" xml:"group"`
	TxBandwidth	uint32		`json:"tx-bandwidth,omitempty" xml:"tx-bandwidth"`
	TxThroughput	uint64		`json:"tx-throughput,omitempty" xml:"tx-throughput"`
	TxPercentage	string		`json:"tx-percentage,omitempty" xml:"tx-percentage"`
	RxBandwidth	uint32		`json:"rx-bandwidth,omitempty" xml:"rx-bandwidth"`
	RxThroughput	uint64		`json:"rx-throughput,omitempty" xml:"rx-throughput"`
	RxPercentage	string		`json:"rx-percentage,omitempty" xml:"rx-percentage"`
	TxRxBandwidth	uint32		`json:"txrx-bandwidth,omitempty" xml:"txrx-bandwidth"`
	TxRxThroughput	uint64		`json:"txrx-throughput,omitempty" xml:"txrx-throughput"`
	TxRxPercentage	string		`json:"txrx-percentage,omitempty" xml:"txrx-percentage"`
}
```

## type [RESTFibrechannelTrunk](<methods.go#L11>)
```go
type RESTFibrechannelTrunk interface {
	Name() string
	URIPath() string
	GetTrunk(groupSourcePort string) (*BrocadeFibrechannelTrunk, error)
	GetTrunkResponse(groupSourcePort string) (*http.Response, error)
	GetPerformance(group string) (*Performance, error)
	GetPerformanceResponse(group string) (*http.Response, error)
}
```

## func [NewRESTFibrechannelTrunk() RESTFibrechannelTrunk](<methods.go#L24>)

```go
func NewRESTFibrechannelTrunk(cfg *api_interface.RESTConfig) RESTFibrechannelTrunk
```

## type [Trunk](<brocadeFibrechannelTrunk.go#L22>)
```go
type Trunk struct {
	Group			string	`json:"group,omitempty" xml:"group"`
	SourcePort		string	`json:"source-port,omitempty" xml:"source-port"`
	Master			bool	`json:"master,omitempty" xml:"master"`
	DestinationPort		string	`json:"destination-port,omitempty" xml:"destination-port"`
	NeighborWwn		string	`json:"neighbor-wwn,omitempty" xml:"neighbor-wwn"`
	NeighborSwitchName	string	`json:"neighbor-switch-name,omitempty" xml:"neighbor-switch-name"`
	NeighborDomainId	string	`json:"neighbor-domain-id,omitempty" xml:"neighbor-domain-id"`
	Deskew			string	`json:"deskew,omitempty" xml:"deskew"`
	TrunkType		string	`json:"trunk-type,omitempty" xml:"trunk-type"`
}
```

## type [TrunkArea](<brocadeFibrechannelTrunk.go#L48>)
```go
type TrunkArea struct {
	XMLName		xml.Name	`json:"-" xml:"trunk-area"`
	TrunkIndex	string		`json:"trunk-index,omitempty" xml:"trunk-index"`
	TrunkMembers	[]string	`json:"trunk-members,omitempty" xml:"trunk-members>trunk-member"`
	MasterPort	string		`json:"master-port,omitempty" xml:"master-port"`
	TrunkActive	bool		`json:"trunk-active,omitempty" xml:"trunk-active"`
}
```

