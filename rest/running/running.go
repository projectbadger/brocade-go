// Broadcom fibrechannel optical switch API data models
/*
REST API documentation: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM
Endpoint list: page 32
*/
package running

import (
	"brocade/rest/api_interface"
	"brocade/rest/running/access_gateway"
	"brocade/rest/running/brocade_interface"
	"brocade/rest/running/chassis"
	"brocade/rest/running/fabric"
	"brocade/rest/running/fdmi"
	"brocade/rest/running/fibrechannel_configuration"
	"brocade/rest/running/fibrechannel_diagnostics"
	"brocade/rest/running/fibrechannel_logical_switch"
	"brocade/rest/running/fibrechannel_switch"
	"brocade/rest/running/fibrechannel_trunk"
	"brocade/rest/running/fru"
	"brocade/rest/running/license"
	"brocade/rest/running/logging"
	"brocade/rest/running/maps"
	"brocade/rest/running/media"
	"brocade/rest/running/name_server"
	"brocade/rest/running/security"
	"brocade/rest/running/snmp"
	"brocade/rest/running/time"
	"brocade/rest/running/zone"
	// "brocade/rest/running/fabric"
	// "brocade/rest/running/fdmi"
	// "brocade/rest/running/fibrechannel_configuration"
	// "brocade/rest/running/fibrechannel_diagnostics"
	// "brocade/rest/running/fibrechannel_logical_switch"
	// "brocade/rest/running/fibrechannel_switch"
	// "brocade/rest/running/fibrechannel_trunk"
	// "brocade/rest/running/fru"
	// "brocade/rest/running/license"
	// "brocade/rest/running/logging"
	// "brocade/rest/running/maps"
	// "brocade/rest/running/media"
	// "brocade/rest/running/name_server"
	// "brocade/rest/running/security"
	// "brocade/rest/running/snmp"
	// "brocade/rest/running/time"
	// "brocade/rest/running/zone"
)

type Running struct {
	BrocadeAccessGateway             access_gateway.BrocadeAccessGateway
	BrocadeChassis                   chassis.BrocadeChassis
	BrocadeFabric                    fabric.BrocadeFabric
	BrocadeFDMI                      fdmi.BrocadeFDMI
	BrocadeFibrechannelConfiguration fibrechannel_configuration.BrocadeFibrechannelConfiguration
	BrocadeFibrechannelDiagnostics   fibrechannel_diagnostics.BrocadeFibrechannelDiagnostics
	// BrocadeFibrechannelInterface     brocade_interface.Fibrechannel
	BrocadeFibrechannelLogicalSwitch fibrechannel_logical_switch.BrocadeFibrechannelLogicalSwitch
	BrocadeFibrechannelSwitch        fibrechannel_switch.BrocadeFibrechannelSwitch
	BrocadeFibrechannelTrunk         fibrechannel_trunk.BrocadeFibrechannelTrunk
	BrocadeFRU                       fru.BrocadeFRU
	BrocadeLicense                   license.BrocadeLicense
	BrocadeLogging                   logging.BrocadeLogging
	BrocadeMaps                      maps.BrocadeMaps
	BrocadeMedia                     media.BrocadeMedia
	BrocadeNameServer                name_server.BrocadeNameServer
	// BrocadeOperations operations.Show
	BrocadeSecurity security.BrocadeSecurity
	BrocadeSNMP     snmp.BrocadeSNMP
	BrocadeTime     time.BrocadeTime
	BrocadeZone     zone.BrocadeZone
}

type restRunningImpl struct {
	config *api_interface.RESTConfig
}

type RESTRunning interface {
	AccessGateway() access_gateway.RESTAccessGateway
	Chassis() chassis.RESTChassis
	Fabric() fabric.RESTFabric
	FDMI() fdmi.RESTFDMI
	FibrechannelConfiguration() fibrechannel_configuration.RESTFibrechannelConfiguration
	FibrechannelInterface() brocade_interface.RESTInterface
	FRU() fru.RESTFRU
}

func NewRunning(config *api_interface.RESTConfig) RESTRunning {
	return &restRunningImpl{
		config: config,
	}
}

func (r *restRunningImpl) ModuleBrocadeAccessGateway() *access_gateway.BrocadeAccessGateway {
	return &access_gateway.BrocadeAccessGateway{}
}

func (r *restRunningImpl) ModuleBrocadeChassis() *chassis.BrocadeChassis {
	return &chassis.BrocadeChassis{}
}

func (r *restRunningImpl) AccessGateway() access_gateway.RESTAccessGateway {
	return access_gateway.NewAccessGateway(r.config)
}

func (r *restRunningImpl) Chassis() chassis.RESTChassis {
	return chassis.NewChassis(r.config)
}

func (r *restRunningImpl) Fabric() fabric.RESTFabric {
	return fabric.NewRESTFabric(r.config)
}

func (r *restRunningImpl) FDMI() fdmi.RESTFDMI {
	return fdmi.NewRESTFDMI(r.config)
}

func (r *restRunningImpl) FibrechannelConfiguration() fibrechannel_configuration.RESTFibrechannelConfiguration {
	return fibrechannel_configuration.NewRESTFibrechannelConfiguration(r.config)
}

func (r *restRunningImpl) FibrechannelInterface() brocade_interface.RESTInterface {
	return brocade_interface.NewRESTInterface(r.config)
}

func (r *restRunningImpl) FRU() fru.RESTFRU {
	return fru.NewRESTFRU(r.config)
}
