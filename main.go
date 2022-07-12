//
/*
REST API documentation: https://docs.broadcom.com/doc/FOS-82X-REST-API-RM
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/projectbadger/brocade-go/config"
	"github.com/projectbadger/brocade-go/log"
	"github.com/projectbadger/brocade-go/rest"
)

func main() {
	err := log.SetupLogging()
	if err != nil {
		panic(err)
	}

	// rest := rest.NewRESTDefault(config.Cfg.Client.Host, config.Cfg.Client.Username, config.Cfg.Client.Password)
	rest := rest.NewRESTXMLSessionless(config.Cfg.Client.Host, config.Cfg.Client.BaseURI, config.Cfg.Client.Username, config.Cfg.Client.Password, http.DefaultClient)
	ps, errs := rest.Running().FRU().GetPowerSupply() // working
	fmt.Printf("ps: '%#v', errs: '%#v'\n", ps, errs)
	fmt.Println(rest.Running().Chassis().Name())
	fmt.Println(rest.Running().FRU().Name())
}
