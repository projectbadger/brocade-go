package config_test

import (
	"fmt"

	"github.com/projectbadger/brocade-go/examples/testing/config"
)

func Example() {
	// config package is initialized immediately,
	// so values can be called directly
	fmt.Println(config.Cfg.Version)
	// Output: false
}
