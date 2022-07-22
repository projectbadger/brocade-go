{{ define "overview" }}# {{ index .CustomVars "name" }}
A library for interfacing with Brocade FOS REST API.

The documentation is taken from the [Broadcom Brocade FOS REST API documentation](https://docs.broadcom.com/doc/FOS-82X-REST-API-RM)

# WIP!

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
{{ end }}