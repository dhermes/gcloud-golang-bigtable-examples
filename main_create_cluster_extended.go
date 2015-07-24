package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	"github.com/dhermes/gcloud-golang/bigtable"
	btcspb "github.com/dhermes/gcloud-golang/bigtable/internal/cluster_service_proto"
	"github.com/golang/protobuf/proto"
)

var opNameRegexp = regexp.MustCompile(`^operations/projects/([^/]+)/zones/([^/]+)/clusters/([a-z][-a-z0-9]*)/operations/([^/]+)$`)

func main() {
	ctx, clientOption, err := getClientArgs(bigtable.ClusterAdminScope)
	if err != nil {
		log.Fatal(err)
		return
	}

	client, err := bigtable.NewClusterAdminClient(
		*ctx, ProjectID, *clientOption)
	opsClient, err := bigtable.NewClusterAdminOperationsClient(
		*ctx, ProjectID, *clientOption)

	// Insert a new cluster.
	cluster, err := client.CreateCluster(*ctx, Zone, ClusterID, DisplayName, ServeNodes)
	if err != nil {
		log.Fatal(err)
		return
	}
	clusterPretty, err := json.MarshalIndent(cluster, "", "  ")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Cluster created:\n%s\n\n", clusterPretty)

	// Get the progress of the long-running operation.
	opName := cluster.CurrentOperation.Name
	m := opNameRegexp.FindStringSubmatch(opName)
	opId := m[4]
	fmt.Printf("Op Name: %s, Op ID: %s\n\n", opName, opId)
	operation, err := opsClient.GetOperation(*ctx, Zone, ClusterID, opId)
	if err != nil {
		log.Fatal(err)
		return
	}
	opPretty, err := json.MarshalIndent(operation, "", "  ")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Operation retrieved:\n%s\n\n", opPretty)

	createClusterMetadataBytes := operation.Metadata.Value
	metadata := &btcspb.CreateClusterMetadata{}
	err = proto.Unmarshal(createClusterMetadataBytes, metadata)
	if err != nil {
		log.Fatal(err)
		return
	}
	metadataPretty, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Metadata:\n%s\n\n", metadataPretty)

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
