// Package config handles all the configuration data.
/*
Before use, the package must be initialized by the function `SetupConfig() error`. It is normally already initialized by the `init()` function.
Example:

	err := config.SetupConfig()

The configuration data is accessible from the `Cfg` variable.

Example:

	configFilePath := config.Cfg.File
*/
package config
