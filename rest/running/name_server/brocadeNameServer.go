package name_server

import "encoding/xml"

type BrocadeNameServer struct {
	XMLName                xml.Name                 `json:"-" xml:"brocade-name-server"`
	FibrechannelNameServer []FibrechannelNameServer `json:"fibrechannel-name-server" xml:"fibrechannel-name-server>fibrechannel-name-server"`
}

type FibrechannelNameServer struct {
	XMLName xml.Name `json:"-" xml:"fibrechannel-name-server"`
}
