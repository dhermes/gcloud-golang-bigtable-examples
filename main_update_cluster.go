package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dhermes/gcloud-golang/bigtable"
)

func main() {
	ctx, clientOption, err := getClientArgs(bigtable.ClusterAdminScope)
	if err != nil {
		log.Fatal(err)
		return
	}

	client, err := bigtable.NewClusterAdminClient(
		*ctx, ProjectID, *clientOption)

	// Update the cluster.
	cluster, err := client.UpdateCluster(
		*ctx, Zone, ClusterID, DisplayName+"Foo", ServeNodes)
	if err != nil {
		log.Fatal(err)
		return
	}
	clusterPretty, err := json.MarshalIndent(cluster, "", "  ")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Cluster updated:\n%s\n\n", clusterPretty)

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
