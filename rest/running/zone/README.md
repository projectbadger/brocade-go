
# zone

## Index

- [type Alias](#type-alias)
- [type BrocadeZone](#type-brocadezone)
- [type Cfg](#type-cfg)
- [type DefinedConfiguration](#type-definedconfiguration)
- [type EffectiveConfiguration](#type-effectiveconfiguration)
- [type EnabledZone](#type-enabledzone)
- [type MemberEntry](#type-memberentry)
- [type Zone](#type-zone)


## type [Alias](<brocadeZone.go#L37>)
```go
type Alias struct {
	XMLName		xml.Name	`json:"-" xml:"alias"`
	AliasName	string		`json:"alias-name" xml:"alias-name"`
	MemberEntry	[]string	`json:"member-entry,omitempty" xml:"member-entry>alias-entry-name"`
}
```

## type [BrocadeZone](<brocadeZone.go#L5>)
```go
type BrocadeZone struct {
	XMLName			xml.Name		`json:"-" xml:"brocade-zone"`
	DefinedConfiguration	DefinedConfiguration	`json:"defined-configuration" xml:"defined-configuration"`
	EffectiveConfiguration	EffectiveConfiguration	`json:"effective-configuration" xml:"effective-configuration"`
}
```

## type [Cfg](<brocadeZone.go#L18>)
```go
type Cfg struct {
	XMLName		xml.Name	`json:"-" xml:"cfg"`
	CfgName		string		`json:"cfg-name" xml:"cfg-name"`
	MemberZone	[]string	`json:"member-zone,omitempty" xml:"member-zone>zone-name"`
}
```

## type [DefinedConfiguration](<brocadeZone.go#L11>)
```go
type DefinedConfiguration struct {
	XMLName	xml.Name	`json:"-" xml:"defined-configuration"`
	Cfg	Cfg		`json:"cfg" xml:"cfg"`
	Zone	Zone		`json:"zone" xml:"zone"`
	Alias	Alias		`json:"alias,omitempty" xml:"alias"`
}
```

## type [EffectiveConfiguration](<brocadeZone.go#L43>)
```go
type EffectiveConfiguration struct {
	XMLName			xml.Name	`json:"-" xml:"effective-configuration"`
	CfgName			string		`json:"cfg-name" xml:"cfg-name"`
	Checksum		string		`json:"checksum,omitempty" xml:"checksum"`
	CfgAction		uint8		`json:"cfg-action,omitempty" xml:"cfg-action"`
	DefaultZoneAccess	uint8		`json:"default-zone-access,omitempty" xml:"default-zone-access"`
	DBMax			uint32		`json:"db-max,omitempty" xml:"db-max"`
	DBAvail			uint32		`json:"db-avail,omitempty" xml:"db-avail"`
	DBCommited		uint32		`json:"db-commited,omitempty" xml:"db-commited"`
	DbTransaction		uint32		`json:"db-transaction,omitempty" xml:"db-transaction"`
	TransactionToken	uint32		`json:"transaction-token,omitempty" xml:"transaction-token"`
	DBChassisWideCommited	uint32		`json:"db-chassis-wide-commited,omitempty" xml:"db-chassis-wide-commited"`
	EnabledZone		EnabledZone	`json:"enabled-zone,omitempty" xml:"enabled-zone"`
}
```

## type [EnabledZone](<brocadeZone.go#L58>)
```go
type EnabledZone struct {
	XMLName		xml.Name	`json:"-" xml:"enabled-zone"`
	ZoneName	string		`json:"zone-name" xml:"zone-name"`
	ZoneType	string		`json:"zone-type,omitempty" xml:"zone-type"`
	MemberEntry	MemberEntry	`json:"member-entry,omitempty" xml:"member-entry"`
}
```

## type [MemberEntry](<brocadeZone.go#L31>)
```go
type MemberEntry struct {
	XMLName			xml.Name	`json:"-" xml:"member-entry"`
	EntryName		string		`json:"entry-name,omitempty" xml:"entry-name"`
	PrincipalEntryName	string		`json:"principal-entry-name,omitempty" xml:"principal-entry-name"`
}
```

## type [Zone](<brocadeZone.go#L24>)
```go
type Zone struct {
	XMLName		xml.Name	`json:"-" xml:"zone"`
	ZoneName	string		`json:"zone-name" xml:"zone-name"`
	ZoneType	string		`json:"zone-type,omitempty" xml:"zone-type"`
	MemberEntry	MemberEntry	`json:"member-entry,omitempty" xml:"member-entry"`
}
```

