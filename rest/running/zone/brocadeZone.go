package zone

import "encoding/xml"

type BrocadeZone struct {
	XMLName                xml.Name               `json:"-" xml:"brocade-zone"`
	DefinedConfiguration   DefinedConfiguration   `json:"defined-configuration" xml:"defined-configuration"`
	EffectiveConfiguration EffectiveConfiguration `json:"effective-configuration" xml:"effective-configuration"`
}

type DefinedConfiguration struct {
	XMLName xml.Name `json:"-" xml:"defined-configuration"`
}

type EffectiveConfiguration struct {
	XMLName xml.Name `json:"-" xml:"effective-configuration"`
}
