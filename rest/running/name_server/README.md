
# name_server

## Index

- [type BrocadeNameServer](#type-brocadenameserver)
- [type FibrechannelNameServer](#type-fibrechannelnameserver)


## type [BrocadeNameServer](<brocadeNameServer.go#L5>)
```go
type BrocadeNameServer struct {
	XMLName			xml.Name			`json:"-" xml:"brocade-name-server"`
	FibrechannelNameServer	[]FibrechannelNameServer	`json:"fibrechannel-name-server" xml:"fibrechannel-name-server>fibrechannel-name-server"`
}
```

## type [FibrechannelNameServer](<brocadeNameServer.go#L10>)
```go
type FibrechannelNameServer struct {
	XMLName xml.Name `json:"-" xml:"fibrechannel-name-server"`
}
```

