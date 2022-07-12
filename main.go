//
/*
# Nagios check for Brocade fibrechannel switches

A basic check for Brocade switches that checks if psus and fans are working
and if all ports with non-default description are online and working.

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
	// fmt.Println(rest.Running().AccessGateway().GetNPortMap())
	// fmt.Println(rest.Running().AccessGateway().GetFPortList())
	// fc, errs := rest.Running().FibrechannelInterface().GetFibrechannelData() // working
	// blade, errs := rest.Running().FRU().GetBlade()
	// fmt.Printf("blade: '%#v', errs: '%#v'\n", blade, errs)
	// fan, errs := rest.Running().FRU().GetFan()
	// fmt.Printf("fan: '%#v', errs: '%#v'\n", fan, errs)
	ps, errs := rest.Running().FRU().GetPowerSupply() // working
	fmt.Printf("ps: '%#v', errs: '%#v'\n", ps, errs)
	fmt.Println(rest.Running().Chassis().Name())
	fmt.Println(rest.Running().FRU().Name())
}
