package config_test

import (
	"brocade/config"
	"fmt"
)

func Example() {
	// config package is initialized immediately,
	// so values can be called directly
	fmt.Println(config.Cfg.Version)
	// Output: false
}
