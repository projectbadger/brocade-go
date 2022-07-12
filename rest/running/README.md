
# running

```go
import brocade/rest/running
```

Broadcom fibrechannel optical switch API data models

REST API documentation: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM
Endpoint list: page 32

## Index

- [type RESTRunning](#type-restrunning)
  - [NewRunning() RESTRunning](#func-newrunning-restrunning)
- [type Running](#type-running)


## type [RESTRunning](<running.go#L77>)
```go
type RESTRunning interface {
	AccessGateway() access_gateway.RESTAccessGateway
	Chassis() chassis.RESTChassis
	Fabric() fabric.RESTFabric
	FDMI() fdmi.RESTFDMI
	FibrechannelConfiguration() fibrechannel_configuration.RESTFibrechannelConfiguration
	FibrechannelInterface() brocade_interface.RESTInterface
	FRU() fru.RESTFRU
}
```

## func [NewRunning() RESTRunning](<running.go#L87>)

```go
func NewRunning(config *api_interface.RESTConfig) RESTRunning
```

## type [Running](<running.go#L49>)
```go
type Running struct {
	BrocadeAccessGateway			access_gateway.BrocadeAccessGateway
	BrocadeChassis				chassis.BrocadeChassis
	BrocadeFabric				fabric.BrocadeFabric
	BrocadeFDMI				fdmi.BrocadeFDMI
	BrocadeFibrechannelConfiguration	fibrechannel_configuration.BrocadeFibrechannelConfiguration
	BrocadeFibrechannelDiagnostics		fibrechannel_diagnostics.BrocadeFibrechannelDiagnostics
	// BrocadeFibrechannelInterface     brocade_interface.Fibrechannel
	BrocadeFibrechannelLogicalSwitch	fibrechannel_logical_switch.BrocadeFibrechannelLogicalSwitch
	BrocadeFibrechannelSwitch		fibrechannel_switch.BrocadeFibrechannelSwitch
	BrocadeFibrechannelTrunk		fibrechannel_trunk.BrocadeFibrechannelTrunk
	BrocadeFRU				fru.BrocadeFRU
	BrocadeLicense				license.BrocadeLicense
	BrocadeLogging				logging.BrocadeLogging
	BrocadeMaps				maps.BrocadeMaps
	BrocadeMedia				media.BrocadeMedia
	BrocadeNameServer			name_server.BrocadeNameServer
	// BrocadeOperations operations.Show
	BrocadeSecurity	security.BrocadeSecurity
	BrocadeSNMP	snmp.BrocadeSNMP
	BrocadeTime	time.BrocadeTime
	BrocadeZone	zone.BrocadeZone
}
```

