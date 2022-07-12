
# license

## Index

- [type BrocadeLicense](#type-brocadelicense)
- [type License](#type-license)
- [type RESTLicense](#type-restlicense)
  - [NewRESTLicense() RESTLicense](#func-newrestlicense-restlicense)


## type [BrocadeLicense](<brocadeLicense.go#L6>)

The container for licenses installed on the switch.
```go
type BrocadeLicense struct {
	XMLName	xml.Name	`json:"-" xml:"brocade-license"`
	License	License		`json:"license" xml:"license"`
}
```

## type [License](<brocadeLicense.go#L12>)

A license installed on the switch.
```go
type License struct {
	XMLName	xml.Name	`json:"-" xml:"license"`
	// The license key for one or more features installed
	// on the switch.
	Name	string	`json:"name" xml:"name"`
	// A list of features that are integrated in a single
	// license key.
	Features	[]string	`json:"features" xml:"features"`
	// The capacity for the license installed on the
	// switch. Note that this parameter is valid only for a
	// capacity-based license.
	Capacity	uint32	`json:"capacity,omitempty" xml:"capacity"`
	// The number of slots configured to use the license
	// installed on the switch. Note that this parameter is
	// valid only for a capacity-based license.
	Consumed	uint32	`json:"consumed,omitempty" xml:"consumed"`
	// A list of slot numbers of the configured blade slots
	// for the license installed on the switch.
	ConfiguredBladeSlots	[]uint32	`json:"configured-blade-slots,omitempty" xml:"configured-blade-slots>configured-blade-slot"`
	// The expiration date for the license installed on the
	// switch.
	// (MM/DD/YYYY format)
	ExpirationDate	string	`json:"expiration-date,omitempty" xml:"expiration-date"`
}
```

## type [RESTLicense](<methods.go#L11>)
```go
type RESTLicense interface {
	Name() string
	GetLicense() (*License, errors.BrocadeErr)
}
```

## func [NewRESTLicense() RESTLicense](<methods.go#L24>)

```go
func NewRESTLicense(config *api_interface.RESTConfig) RESTLicense
```

