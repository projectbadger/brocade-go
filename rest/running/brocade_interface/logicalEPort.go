package brocade_interface

type LogicalEPort struct {
	PortIndex               int      `json:"port-index,omitempty" xml:"port-index"`
	FabricId                string   `json:"fabric-id,omitempty" xml:"fabric-id"`
	OperationalStatus       string   `json:"operational-status,omitempty" xml:"operational-status"`
	OfflineReason           string   `json:"offline-reason,omitempty" xml:"offline-reason"`
	NeighborNodeWWN         string   `json:"neighbor-node-wwn,omitempty" xml:"neighbor-node-wwn"`
	AssociatedPhysicalPorts []string `json:"associated-physical-ports,omitempty" xml:"associated-physical-ports>port"`
}
