
# Brocade FOS REST API interface in Go
A library for interfacing with Brocade FOS REST API.

The documentation is taken from the [Broadcom Brocade FOS REST API documentation](https://docs.broadcom.com/doc/FOS-82X-REST-API-RM)

# WIP!

Not all of the models have been implemented yet and methods only provide read functions.

# Using the models

The models are in individual packages in `rest/running`

# Using the API

To use the REST API, you must instantiate a session and create a REST interface implementation with it. If the session is not sessionless, it is imperative to log out after use.

```go
rest := rest.NewRESTJSONSessionless("localhost:8080", "/rest", "username", "password", http.DefaultClient)
ps, errs := rest.Running().FRU().GetPowerSupply()
```

To ensure proper login and logout, a helper session function can be used for non-sessionless login.
```go
rest := rest.NewRESTJSONSession("localhost:8080", "/rest", "username", "password", http.DefaultClient)
err := rest.Session(func(session session.Session) error {
	ps, errs := rest.Running().FRU().GetPowerSupply()
	if errs != nil {
		return errs
	}
})
```


## Subpackages

  - [main](/examples/testing)
    - [config](/examples/testing/config)
    - [log](/examples/testing/log)
- [rest](/rest)
  - [api_interface](/rest/api_interface)
  - [endpoints](/rest/endpoints)
  - [errors](/rest/errors)
  - [running](/rest/running)
    - [access_gateway](/rest/running/access_gateway)
    - [brocade_interface](/rest/running/brocade_interface)
    - [chassis](/rest/running/chassis)
    - [endpoints](/rest/running/endpoints)
    - [fabric](/rest/running/fabric)
    - [fdmi](/rest/running/fdmi)
    - [fibrechannel_configuration](/rest/running/fibrechannel_configuration)
    - [fibrechannel_diagnostics](/rest/running/fibrechannel_diagnostics)
    - [fibrechannel_logical_switch](/rest/running/fibrechannel_logical_switch)
    - [fibrechannel_switch](/rest/running/fibrechannel_switch)
    - [fibrechannel_trunk](/rest/running/fibrechannel_trunk)
    - [fru](/rest/running/fru)
    - [license](/rest/running/license)
    - [logging](/rest/running/logging)
    - [maps](/rest/running/maps)
    - [media](/rest/running/media)
    - [name_server](/rest/running/name_server)
    - [operations](/rest/running/operations)
    - [security](/rest/running/security)
    - [snmp](/rest/running/snmp)
    - [time](/rest/running/time)
    - [zone](/rest/running/zone)
  - [utils](/rest/utils)
- [session](/session)
  - [test_data](/session/test_data)
- [utils](/utils)
  - [test_data](/utils/test_data)

