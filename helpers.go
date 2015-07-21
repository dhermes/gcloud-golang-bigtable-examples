package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
)

func getFlags() (bool, bool) {
	var useAppDefault bool
	var verbose bool
	flag.BoolVar(
		&useAppDefault, "use-app-default", false,
		"Boolean to determine if app. default credentials should be used")
	flag.BoolVar(
		&verbose, "verbose", false,
		"Boolean to determine if verbosity level should be increased")
	flag.Parse()
	return useAppDefault, verbose
}

func getAppDefaultClientArgs() (*context.Context, *cloud.ClientOption, error) {
	ctx := oauth2.NoContext
	tokenSrc, err := google.DefaultTokenSource(ctx, ScopeCloudPlatform)
	if err != nil {
		return nil, nil, err
	}

	clientOption := cloud.WithTokenSource(tokenSrc)
	return &ctx, &clientOption, nil
}

func getServiceAccountClientArgs(jwtScope string) (*context.Context, *cloud.ClientOption, error) {
	jsonKey, err := ioutil.ReadFile(KeyFile)
	if err != nil {
		return nil, nil, err
	}

	config, err := google.JWTConfigFromJSON(jsonKey, jwtScope)
	if err != nil {
		return nil, nil, err
	}

	ctx := cloud.NewContext(ProjectID, config.Client(oauth2.NoContext))
	clientOption := cloud.WithTokenSource(config.TokenSource(ctx))
	return &ctx, &clientOption, nil
}

func getClientArgs(jwtScope string) (*context.Context, *cloud.ClientOption, error) {
	useAppDefault, verbose := getFlags()
	if useAppDefault {
		if verbose {
			fmt.Println("Using App. Default Credentials.")
		}
		return getAppDefaultClientArgs()
	} else {
		if verbose {
			fmt.Println("Using Service Account.")
		}
		return getServiceAccountClientArgs(jwtScope)
	}
}
