
# zone

## Index

- [type BrocadeZone](#type-brocadezone)
- [type DefinedConfiguration](#type-definedconfiguration)
- [type EffectiveConfiguration](#type-effectiveconfiguration)


## type [BrocadeZone](<brocadeZone.go#L5>)
```go
type BrocadeZone struct {
	XMLName			xml.Name		`json:"-" xml:"brocade-zone"`
	DefinedConfiguration	DefinedConfiguration	`json:"defined-configuration" xml:"defined-configuration"`
	EffectiveConfiguration	EffectiveConfiguration	`json:"effective-configuration" xml:"effective-configuration"`
}
```

## type [DefinedConfiguration](<brocadeZone.go#L11>)
```go
type DefinedConfiguration struct {
	XMLName xml.Name `json:"-" xml:"defined-configuration"`
}
```

## type [EffectiveConfiguration](<brocadeZone.go#L15>)
```go
type EffectiveConfiguration struct {
	XMLName xml.Name `json:"-" xml:"effective-configuration"`
}
```

