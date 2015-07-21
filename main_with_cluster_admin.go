package main

import (
	"fmt"
	"log"

	"google.golang.org/cloud/bigtable"
)

func main() {
	ctx, clientOption, err := getClientArgs(bigtable.ClusterAdminScope)
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
	fmt.Println("Result of List Clusters:")
	for i := 0; i < len(clusterInfo); i++ {
		fmt.Printf("    %v\n", clusterInfo[i].Name)
	}
}
