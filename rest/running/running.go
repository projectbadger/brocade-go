// Broadcom fibrechannel optical switch API data models
/*
REST API documentation: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM
Endpoint list: page 32
*/
package running

import (
	"github.com/projectbadger/brocade-go/rest/api_interface"
	"github.com/projectbadger/brocade-go/rest/running/access_gateway"
	"github.com/projectbadger/brocade-go/rest/running/brocade_interface"
	"github.com/projectbadger/brocade-go/rest/running/chassis"
	"github.com/projectbadger/brocade-go/rest/running/fabric"
	"github.com/projectbadger/brocade-go/rest/running/fdmi"
	"github.com/projectbadger/brocade-go/rest/running/fibrechannel_configuration"
	"github.com/projectbadger/brocade-go/rest/running/fibrechannel_diagnostics"
	"github.com/projectbadger/brocade-go/rest/running/fibrechannel_logical_switch"
	"github.com/projectbadger/brocade-go/rest/running/fibrechannel_switch"
	"github.com/projectbadger/brocade-go/rest/running/fibrechannel_trunk"
	"github.com/projectbadger/brocade-go/rest/running/fru"
	"github.com/projectbadger/brocade-go/rest/running/license"
	"github.com/projectbadger/brocade-go/rest/running/logging"
	"github.com/projectbadger/brocade-go/rest/running/maps"
	"github.com/projectbadger/brocade-go/rest/running/media"
	"github.com/projectbadger/brocade-go/rest/running/name_server"
	"github.com/projectbadger/brocade-go/rest/running/security"
	"github.com/projectbadger/brocade-go/rest/running/snmp"
	"github.com/projectbadger/brocade-go/rest/running/time"
	"github.com/projectbadger/brocade-go/rest/running/zone"
)

type Running struct {
	BrocadeAccessGateway             access_gateway.BrocadeAccessGateway
	BrocadeChassis                   chassis.BrocadeChassis
	BrocadeFabric                    fabric.BrocadeFabric
	BrocadeFDMI                      fdmi.BrocadeFDMI
	BrocadeFibrechannelConfiguration fibrechannel_configuration.BrocadeFibrechannelConfiguration
	BrocadeFibrechannelDiagnostics   fibrechannel_diagnostics.BrocadeFibrechannelDiagnostics
	BrocadeFibrechannelInterface     brocade_interface.Fibrechannel
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

// Running interface for interacting with modules
type RESTRunning interface {
	// Returns Brocade Access Gateway interface
	AccessGateway() access_gateway.RESTAccessGateway
	// Returns Brocade Chassis interface
	Chassis() chassis.RESTChassis
	// Returns Brocade Fabric interface
	Fabric() fabric.RESTFabric
	// Returns FDMI interface
	FDMI() fdmi.RESTFDMI
	// Returns Brocade Fibrechannel Configuration interface
	FibrechannelConfiguration() fibrechannel_configuration.RESTFibrechannelConfiguration
	// Returns Brocade Fibrechannel Interface interface
	FibrechannelInterface() brocade_interface.RESTInterface
	// Returns Brocade FRU interface
	FRU() fru.RESTFRU
}

// Returns a new Running interface from config
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
