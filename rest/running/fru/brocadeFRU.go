package fru

import "encoding/xml"

// Configuration and runtime information of the FRUs
// installed in the chassis.
//
// Documentation source: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM#%5B%7B%22num%22%3A370%2C%22gen%22%3A0%7D%2C%7B%22name%22%3A%22XYZ%22%7D%2C36%2C722.904%2Cnull%5D - page 160 (20220701)
type BrocadeFRU struct {
	XMLName xml.Name `json:"-" xml:"brocade-fru"`
	// A list of blade details for the specified slot
	// number.
	Blade []Blade `json:"blade,omitempty" xml:"blade"`
	// The details about the fan units.
	Fan         []Fan         `json:"fan,omitempty" xml:"fan"`
	PowerSupply []PowerSupply `json:"power-supply,omitempty" xml:"power-supply"`
}

type Blade struct {
	XMLName xml.Name `json:"-" xml:"blade"`
	// The number of the physical slot in the chassis in
	// which the blade is inserted.
	SlotNumber uint16 `json:"slot-number" xml:"slot-number"`
	// The name of the blade manufacturer.
	Manufacturer string `json:"manufacturer" xml:"manufacturer"`
	// The type of blade (sw blade, cp blade, core blade).
	BladeType string `json:"blade-type,omitempty" xml:"blade-type"`
	// The ID of the blade. This parameter is available
	// only when the blade type is not unknown
	// (blade-type != 'unknown').
	BladeId uint16 `json:"blade-id" xml:"blade-id"`
	// The current state of the blade.
	BladeState string `json:"blade-state,omitempty" xml:"blade-state"`
	// A printable ASCII string that describes the model of
	// the blade. This parameter is available only when the
	// blade type is not unknown (blade-type != 'unknown')
	// or the slot number is not zero (slot-number != 0).
	ModelName string `json:"model-name,omitempty" xml:"model-name"`
	// A human-readable string that displays the firmware
	// version running on the switch. For an AP blade, it
	// displays the AP blade firmware version. For other
	// blades, it displays the active CP firmware version.
	FirmwareVersion string `json:"firmware-version,omitempty" xml:"firmware-version"`
	// The number of FC ports supported by the blade.
	// This parameter is available only when the blade type
	// is not CP or not unknown (blade-type != 'cp blade'
	// or blade-type != 'unknown').
	FcPortCount uint16 `json:"fc-port-count,omitempty" xml:"fc-port-count"`
	// The number of GE ports supported by the blade. This
	// parameter is available only when extension is
	// enabled or the blade ID is 75
	// (extension-enabled = true or blade-id = 75).
	GePortCount uint16 `json:"ge-port-count,omitempty" xml:"ge-port-count"`
	// A list of IP addresses assigned to the CP blade.
	// This parameter is available only for CP blades
	// (blade-type = 'cp blade').
	IpAddressList []string `json:"ip-address-list,omitempty" xml:"ip-address-list>ip-address"`
	// A list of gateway IP addresses assigned to the CP
	// blade. This parameter is available only for CP blades
	// (blade-type = 'cp blade').
	IpGatewayList []string `json:"ip-gateway-list,omitempty" xml:"ip-gateway-list>ip-gateway"`
	// The subnet mask of the network. This parameter is
	// available only for CP blades
	// (blade-type = 'cp blade').
	SubnetMask string `json:"subnet-mask,omitempty" xml:"subnet-mask"`
	// A printable ASCII string that identifies the
	// firmware version running on the primary partition of
	// the CP blade. This parameter is available only for
	// CP blades (blade-type = 'cp blade').
	PrimaryFirmwareVersion string `json:"primary-firmware-version,omitempty" xml:"primary-firmware-version"`
	// The maximum power consumption allocated for the
	// blade.
	PowerConsumption string `json:"power-consumption,omitempty" xml:"power-consumption"`
	// The real-time power consumed by the FRU.
	PowerUsage string `json:"power-usage,omitempty" xml:"power-usage"`
	// A printable ASCII string that identifies the
	// firmware version running on the secondary partition
	// of the CP blade
	SecondaryFirmwareVersion string `json:"secondary-firmware-version,omitempty" xml:"secondary-firmware-version"`
	// Whether the switch or blade supports extension. For
	// blade IDs 154, 186, and 213, the value is set to
	// enabled (true).
	ExtensionEnabled string `json:"extension-enabled,omitempty" xml:"extension-enabled"`
	// The application mode configuration of the extension
	// blade or switch. This parameter is available only
	// when extension is enabled
	// (extension-enabled = true).
	//
	// Note that a change in the application mode
	// configuration is a disruptive operation and that
	// when the configuration change is successful, a
	// switch automatically reboots and a blade
	// automatically powers off and on. Blade ID 213
	// supports only hybrid mode. For blade ID 213, the
	// default is hybrid mode. For blade ID 154 and 186,
	// the default is FCIP mode.
	ExtensionAppMode string `json:"extension-app-mode,omitempty" xml:"extension-app-mode"`
	// The VE mode configuration of the extension blade or
	// switch. This parameter is available only
	// when extension is enabled
	// (extension-enabled = true).
	//
	// Note that a change in the VE mode
	// configuration is a disruptive operation and that
	// when the configuration change is successful, a
	// switch automatically reboots and a blade
	// automatically powers off and on. For blade ID 154
	// and 186, the default is 10VE mode. For blade ID 213,
	// the default is not applicable.
	ExtensionVeMode string `json:"extension-ve-mode,omitempty" xml:"extension-ve-mode"`
	// The GE mode configuration of the extension blade or
	// switch. This parameter is applicable only on the
	// Brocade 7810 Extension Switch.
	ExtensionGeMode string `json:"extension-ge-mode,omitempty" xml:"extension-ge-mode"`
	// The number of days the FRU has been powered on.
	TimeAlive string `json:"time-alive,omitempty" xml:"time-alive"`
	// The number of days since the FRU was last powered on.
	TimeAwake string `json:"time-awake,omitempty" xml:"time-awake"`
	// The part number assigned by the organization
	// responsible for producing or manufacturing of the
	// FRU.
	PartNumber string `json:"part-number,omitempty" xml:"part-number"`
	// The serial number of the FRU.
	SerialNumber string `json:"serial-number,omitempty" xml:"serial-number"`
}

// The details about a fan unit.
type Fan struct {
	XMLName xml.Name `json:"-" xml:"power-supply"`
	// The physical slot number in the chassis where the fan is located.
	UnitNumber uint `json:"unit-number,omitempty" xml:"unit-number"`
	// The maximum power consumption allocated for the fan
	// in watts.
	PowerConsumption int `json:"power-consumption,omitempty" xml:"power-consumption"`
	// The current operational state of the fan.
	OperationalState string `json:"operational-state,omitempty" xml:"operational-state"`
	// The fan speed in RPM.
	Speed uint32 `json:"speed,omitempty" xml:"speed"`
	// The air flow direction of the fan blowers.
	AirflowDirection string `json:"airflow-direction,omitempty" xml:"airflow-direction"`
	// The number of days the FRU has been powered on.
	TimeAlive int `json:"time-alive,omitempty" xml:"time-alive"`
	// The number of days since the FRU was last powered on.
	TimeAwake int `json:"time-awake,omitempty" xml:"time-awake"`
	// The part number assigned by the organization
	// responsible for producing or manufacturing of the
	// FRU.
	PartNumber string `json:"part-number,omitempty" xml:"part-number"`
	// The serial number of the FRU.
	SerialNumber string `json:"serial-number,omitempty" xml:"serial-number"`
}

// The details about a power supply unit.
type PowerSupply struct {
	XMLName xml.Name `json:"-" xml:"power-supply"`
	// The physical slot number of the chassis where the
	// power supply is located.
	UnitNumber uint16 `json:"unit-number,omitempty" xml:"unit-number"`
	// The maximum power production allocated for the power
	// supply unit in watts.
	PowerProduction uint32 `json:"power-production,omitempty" xml:"power-production"`
	// The input voltage of the power supply unit in volts.
	InputVoltage float64 `json:"input-voltage,omitempty" xml:"input-voltage"`
	// The part number assigned by the organization
	// responsible for producing or manufacturing of the
	// FRU.
	PartNumber string `json:"part-number,omitempty" xml:"part-number"`
	// The serial number of the FRU.
	SerialNumber string `json:"serial-number,omitempty" xml:"serial-number"`
	// The air flow direction of the power supply fans.
	AirflowDirection string `json:"airflow-direction,omitempty" xml:"airflow-direction"`
	// The power supply input voltage type.
	PowerSource string `json:"power-source,omitempty" xml:"power-source"`
	// The operational state of the power supply.
	OperationalState string `json:"operational-state,omitempty" xml:"operational-state"`
	// The temperature of the power supply sensor in
	// centigrade.
	Temperature float64 `json:"temperature,omitempty" xml:"temperature"`
	// Whether the temperature
	TemperatureSensorSupported bool `json:"temperature-sensor-supported,omitempty" xml:"temperature-sensor-supported"`
	// The real-time power consumed by the FRU.
	PowerUsage int `json:"power-usage,omitempty" xml:"power-usage"`
	// The number of days the FRU has been powered on.
	TimeAlive int `json:"time-alive,omitempty" xml:"time-alive"`
	// The number of days since the FRU was last powered on.
	TimeAwake int `json:"time-awake,omitempty" xml:"time-awake"`
}
