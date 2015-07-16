package main

import (
	"fmt"
	"log"

	"google.golang.org/cloud/bigtable"
)

func main() {
	ctx, clientOption, err := getClientArgs()
	if err != nil {
		log.Fatal(err)
		return
	}

	client, err := bigtable.NewClusterAdminClient(
		*ctx, ProjectID, *clientOption)

	// Get the list of clusters.
	clusterInfo, err := client.Clusters(*ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Result of List Clusters: %v\n", clusterInfo)
}
