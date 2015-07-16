package main

import (
	"flag"
	"io/ioutil"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/bigtable"
)

func getUseAppDefault() bool {
	var useAppDefault bool
	flag.BoolVar(
		&useAppDefault, "use-app-default", false,
		"Boolean to determine if app. default credentials should be used")
	flag.Parse()
	return useAppDefault
}

func getClientArgs() (*context.Context, *cloud.ClientOption, error) {
	useAppDefault := getUseAppDefault()

	var ctx context.Context
	var clientOption cloud.ClientOption
	if useAppDefault {
		tokenSrc, err := google.DefaultTokenSource(
			oauth2.NoContext, ScopeCloudPlatform)
		if err != nil {
			return nil, nil, err
		}

		ctx = oauth2.NoContext
		clientOption = cloud.WithTokenSource(tokenSrc)
	} else {
		jsonKey, err := ioutil.ReadFile(KeyFile)
		if err != nil {
			return nil, nil, err
		}

		config, err := google.JWTConfigFromJSON(jsonKey, bigtable.ClusterAdminScope)
		if err != nil {
			return nil, nil, err
		}

		ctx = cloud.NewContext(ProjectID, config.Client(oauth2.NoContext))
		clientOption = cloud.WithTokenSource(config.TokenSource(ctx))
	}
	return &ctx, &clientOption, nil
}
