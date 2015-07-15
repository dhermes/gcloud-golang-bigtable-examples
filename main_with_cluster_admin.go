package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/bigtable"
)

func main() {
	jsonKey, err := ioutil.ReadFile(KeyFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	config, err := google.JWTConfigFromJSON(jsonKey, bigtable.ClusterAdminScope)
	if err != nil {
		log.Fatal(err)
		return
	}

	ctx := cloud.NewContext(ProjectID, config.Client(oauth2.NoContext))

	client, err := bigtable.NewClusterAdminClient(
		ctx, ProjectID, cloud.WithTokenSource(config.TokenSource(ctx)))

	// Get the list of clusters.
	clusterInfo, err := client.Clusters(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Result of List Clusters: %v\n", clusterInfo)
}
