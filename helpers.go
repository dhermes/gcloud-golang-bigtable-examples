package main

import (
	"flag"
)

func getUseAppDefault() bool {
	var useAppDefault bool
	flag.BoolVar(
		&useAppDefault, "use-app-default", false,
		"Boolean to determine if app. default credentials should be used")
	flag.Parse()
	return useAppDefault
}
