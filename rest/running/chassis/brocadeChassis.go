package chassis

import (
	"encoding/xml"
)

// Configuration and runtime information of the chassis.
//
// Documentation source: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM - page 99 (20220701)
type BrocadeChassis struct {
	XMLName xml.Name `json:"-" xml:"brocade-chassis"`
	// The complete details of the chassis.
	Chassis Chassis `json:"chassis" xml:"chassis"`
	// The High Availability control processor (CP) status
	HAStatus HAStatus `json:"ha-status" xml:"ha-status"`
}

// The complete details of the chassis.
type Chassis struct {
	XMLName xml.Name `json:"-" xml:"chassis"`
	// An ASCII name assigned to the switch chassis by the
	// administrator.
	ChassisUserFriendlyName string `json:"chassis-user-friendly-name,omitempty" xml:"chassis-user-friendly-name"`
	// The WWN of the chassis, which is used for the
	// license.
	ChassisWWN string `json:"chassis-wwn,omitempty" xml:"chassis-wwn"`
	// A printable ASCII string that specifies the serial
	// number of the chassis.
	SerialNumber string `json:"serial-number,omitempty" xml:"serial-number"`
	// The manufacturer of the chassis.
	Manufacturer string `json:"manufacturer,omitempty" xml:"manufacturer"`
	// The part number for the physical element assigned by
	// the manufacturer.
	PartNumber string `json:"part-number,omitempty" xml:"part-number"`
	// A serial number that is used for entitlement support.
	EntitlementSerialNumber string `json:"entitlement-serial-number,omitempty" xml:"entitlement-serial-number"`
	// The maximum number of blades that can fit in the
	// physical chassis. This includes switch, control
	// processor, application, and core routing blades.
	MaxBladesSupported int `json:"max-blades-supported,omitempty" xml:"max-blades-supported"`
	// The serial number of the chassis assigned by the
	// vendor.
	VendorSerialNumber string `json:"vendor-serial-number,omitempty" xml:"vendor-serial-number"`
	// The part number of the chassis assigned by the
	// vendor.
	VendorPartNumber string `json:"vendor-part-number,omitempty" xml:"vendor-part-number"`
	// The revision number of the chassis assigned by the
	// vendor.
	VendorRevisionNumber string `json:"vendor-revision-number,omitempty" xml:"vendor-revision-number"`
	// The product name of the chassis.
	ProductName string `json:"product-name,omitempty" xml:"product-name"`
	// Whether Virtual Fabrics is enabled on the chassis.
	VFEnabled bool `json:"vf-enabled,omitempty" xml:"vf-enabled"`
	// Whether Virtual Fabrics is supported on the chassis.
	VFSupported bool `json:"vf-supported,omitempty" xml:"vf-supported"`
	// The current date of the switch.
	Date string `json:"date,omitempty" xml:"date"`
}

func (c *BrocadeChassis) Name() string {
	return "brocade-chassis"
}

func (b *Chassis) URI() string {
	return "/rest/brocade-chassis/chassis"
}

// The High Availability control processor (CP) status,
// which includes the following information:
//   * Local CP state (slot number and CP ID) and warm or cold
//   * Remote CP state (slot number and CP ID)
//   * High Availability (enabled or disabled)
//   * Heartbeat (up or down)
//   * Health of standby CP
//   * HA synchronization status
type HAStatus struct {
	XMLName xml.Name `json:"-" xml:"ha-status"`
	// The ID of the active CP.
	//
	// cp0 or cp1
	ActiveCP string `json:"active-cp,omitempty" xml:"active-cp"`
	// The ID of the standby CP.
	//
	// cp0 or cp1
	StandbyCP string `json:"standby-cp,omitempty" xml:"standby-cp"`
	// The slot number of the active CP.
	//
	// 0 through 12
	ActiveSlot uint16 `json:"active-slot,omitempty" xml:"active-slot"`
	// The slot number of the standby CP.
	//
	// 0 through 12
	StandbySlot uint16 `json:"standby-slot,omitempty" xml:"standby-slot"`
	// The recovery status of the switch
	//
	// cold or warm
	RecoveryType string `json:"recovery-type,omitempty" xml:"recovery-type"`
	// The health status of the standby CP.
	//
	// healthy | faulted | unknown | non-redundant | not-available
	StandbyHealth string `json:"standby-health,omitempty" xml:"standby-health"`
	// Whether High Availability (HA) is enabled or
	// disabled.
	HAEnabled bool `json:"ha-enabled,omitempty" xml:"ha-enabled"`
	// Whether the heartbeat to the standby CP is up or
	// down.
	HeartbeatUp bool `json:"heartbeat-up,omitempty" xml:"heartbeat-up"`
	// Whether HA is in a synchronized state.
	HASynchronized bool `json:"ha-synchronized,omitempty" xml:"ha-synchronized"`
}

func (b *HAStatus) URI() string {
	return "ha-status"
}
