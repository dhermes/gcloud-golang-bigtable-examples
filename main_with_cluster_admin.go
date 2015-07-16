package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/bigtable"
)

func main() {
	var useAppDefault bool
	flag.BoolVar(
		&useAppDefault, "use-app-default", false,
		"Boolean to determine if app. default credentials should be used")
	flag.Parse()

	jsonKey, err := ioutil.ReadFile(KeyFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	var ctx context.Context
	var clientOption cloud.ClientOption
	if useAppDefault {
		tokenSrc, err := google.DefaultTokenSource(
			oauth2.NoContext, ScopeCloudPlatform)
		if err != nil {
			log.Fatal(err)
			return
		}

		ctx = oauth2.NoContext
		clientOption = cloud.WithTokenSource(tokenSrc)
	} else {
		config, err := google.JWTConfigFromJSON(jsonKey, bigtable.ClusterAdminScope)
		if err != nil {
			log.Fatal(err)
			return
		}

		ctx = cloud.NewContext(ProjectID, config.Client(oauth2.NoContext))
		clientOption = cloud.WithTokenSource(config.TokenSource(ctx))
	}

	client, err := bigtable.NewClusterAdminClient(
		ctx, ProjectID, clientOption)

	// Get the list of clusters.
	clusterInfo, err := client.Clusters(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Result of List Clusters: %v\n", clusterInfo)
}
