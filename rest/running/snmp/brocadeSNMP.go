package snmp

import "encoding/xml"

type BrocadeSNMP struct {
	XMLName xml.Name `json:"-" xml:"brocade-snmp"`
}

type System struct {
	XMLName xml.Name `json:"-" xml:"system"`
}
