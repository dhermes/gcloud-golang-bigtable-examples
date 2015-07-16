package main

import (
	"flag"
	"io/ioutil"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
)

func getUseAppDefault() bool {
	var useAppDefault bool
	flag.BoolVar(
		&useAppDefault, "use-app-default", false,
		"Boolean to determine if app. default credentials should be used")
	flag.Parse()
	return useAppDefault
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
	useAppDefault := getUseAppDefault()
	if useAppDefault {
		return getAppDefaultClientArgs()
	} else {
		return getServiceAccountClientArgs(jwtScope)
	}
}
