package security

import "encoding/xml"

type BrocadeSecurity struct {
	XMLName xml.Name `json:"-" xml:"brocade-security"`
}
