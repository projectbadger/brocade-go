package snmp

import "encoding/xml"

type BrocadeSNMP struct {
	XMLName        xml.Name         `json:"-" xml:"brocade-snmp"`
	System         System           `json:"system" xml:"system"`
	MIBCapability  []MIBCapability  `json:"mib-capability,omitempty" xml:"mib-capability"`
	TrapCapability []TrapCapability `json:"trap-capability,omitempty" xml:"trap-capability"`
	V1Account      []V1Account      `json:"v1-account,omitempty" xml:"v1-account"`
	V1Trap         []V1Trap         `json:"v1-trap,omitempty" xml:"v1-trap"`
	V3Account      []V3Account      `json:"v3-account,omitempty" xml:"v3-account"`
	V3Trap         []V3Trap         `json:"v3-trap,omitempty" xml:"v3-trap"`
	AccessControl  []AccessControl  `json:"access-control,omitempty" xml:"access-control"`
}

type System struct {
	XMLName           xml.Name `json:"-" xml:"system"`
	Description       string   `json:"description,omitempty" xml:"description"`
	Location          string   `json:"location,omitempty" xml:"location"`
	Contact           string   `json:"contact,omitempty" xml:"contact"`
	InformsEnabled    bool     `json:"informs-enabled,omitempty" xml:"informs-enabled"`
	EncryptionEnabled bool     `json:"encryption-enabled,omitempty" xml:"encryption-enabled"`
	AuditInterval     uint16   `json:"audit-interval,omitempty" xml:"audit-interval"`
	DefaultConfig     []string `json:"default-config,omitempty" xml:"default-config>default-control"`
	SecurityGetLevel  string   `json:"security-get-level,omitempty" xml:"security-get-level"`
	SecuritySetLevel  string   `json:"security-set-level,omitempty" xml:"security-set-level"`
	SNMPv1Enabled     bool     `json:"snmpv1-enabled,omitempty" xml:"snmpv1-enabled"`
}

type MIBCapability struct {
	XMLName           xml.Name `json:"-" xml:"mib-capability"`
	MIBName           string   `json:"mib-name" xml:"mib-name"`
	IsMIBEnabledState bool     `json:"is-mib-enabled-state,omitempty" xml:"is-mib-enabled-state"`
}

type TrapCapability struct {
	XMLName            xml.Name `json:"-" xml:"trap-capability"`
	TrapName           string   `json:"trap-name" xml:"trap-name"`
	IsTrapEnabledState bool     `json:"is-trap-enabled-state,omitempty" xml:"is-trap-enabled-state"`
	Severity           string   `json:"severity,omitempty" xml:"severity"`
}

type V1Account struct {
	XMLName        xml.Name `json:"-" xml:"v1-account"`
	Index          uint16   `json:"index" xml:"index"`
	CommunityName  string   `json:"community-name,omitempty" xml:"community-name"`
	CommunityGroup string   `json:"community-group,omitempty" xml:"community-group"`
}

type V1Trap struct {
	XMLName           xml.Name `json:"-" xml:"v1-trap"`
	Index             uint16   `json:"index" xml:"index"`
	Host              string   `json:"host,omitempty" xml:"host"`
	TrapSeverityLevel string   `json:"trap-severity-level,omitempty" xml:"trap-severity-level"`
	PortNumber        int      `json:"port-number,omitempty" xml:"port-number"`
}

type V3Account struct {
	XMLName                xml.Name `json:"-" xml:"v3-account"`
	Index                  uint16   `json:"index" xml:"index"`
	UserName               string   `json:"user-name,omitempty" xml:"user-name"`
	UserGroup              string   `json:"user-group,omitempty" xml:"user-group"`
	AuthenticationProtocol string   `json:"authentication-protocol,omitempty" xml:"authentication-protocol"`
	PrivacyProtocol        string   `json:"privacy-protocol,omitempty" xml:"privacy-protocol"`
	AuthenticationPassword string   `json:"authentication-password,omitempty" xml:"authentication-password"`
	PrivacyPassword        string   `json:"privacy-password,omitempty" xml:"privacy-password"`
	ManagerEngineId        string   `json:"manager-engine-id,omitempty" xml:"manager-engine-id"`
}

type V3Trap struct {
	XMLName           xml.Name `json:"-" xml:"v3-trap"`
	TrapIndex         uint16   `json:"trap-index,omitempty" xml:"trap-index"`
	USMIndex          uint16   `json:"usm-index,omitempty" xml:"usm-index"`
	Host              string   `json:"host,omitempty" xml:"host"`
	TrapSeverityLevel string   `json:"trap-severity-level,omitempty" xml:"trap-severity-level"`
	PortNumber        int      `json:"port-number,omitempty" xml:"port-number"`
	InformsEnabled    bool     `json:"informs-enabled,omitempty" xml:"informs-enabled"`
}

type AccessControl struct {
	XMLName     xml.Name `json:"-" xml:"v3-trap"`
	Index       uint16   `json:"index,omitempty" xml:"index"`
	Host        string   `json:"host,omitempty" xml:"host"`
	AccessLevel string   `json:"access-level,omitempty" xml:"access-level"`
}
