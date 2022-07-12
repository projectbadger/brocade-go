package license

import "encoding/xml"

// The container for licenses installed on the switch.
type BrocadeLicense struct {
	XMLName xml.Name `json:"-" xml:"brocade-license"`
	License License  `json:"license" xml:"license"`
}

// A license installed on the switch.
type License struct {
	XMLName xml.Name `json:"-" xml:"license"`
	// The license key for one or more features installed
	// on the switch.
	Name string `json:"name" xml:"name"`
	// A list of features that are integrated in a single
	// license key.
	Features []string `json:"features" xml:"features"`
	// The capacity for the license installed on the
	// switch. Note that this parameter is valid only for a
	// capacity-based license.
	Capacity uint32 `json:"capacity,omitempty" xml:"capacity"`
	// The number of slots configured to use the license
	// installed on the switch. Note that this parameter is
	// valid only for a capacity-based license.
	Consumed uint32 `json:"consumed,omitempty" xml:"consumed"`
	// A list of slot numbers of the configured blade slots
	// for the license installed on the switch.
	ConfiguredBladeSlots []uint32 `json:"configured-blade-slots,omitempty" xml:"configured-blade-slots>configured-blade-slot"`
	// The expiration date for the license installed on the
	// switch.
	// (MM/DD/YYYY format)
	ExpirationDate string `json:"expiration-date,omitempty" xml:"expiration-date"`
}
