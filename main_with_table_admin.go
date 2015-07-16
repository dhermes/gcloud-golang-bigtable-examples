package main

import (
	"fmt"
	"log"

	"google.golang.org/cloud/bigtable"
)

func main() {
	ctx, clientOption, err := getClientArgs(bigtable.AdminScope)
	if err != nil {
		log.Fatal(err)
		return
	}

	client, err := bigtable.NewAdminClient(
		*ctx, ProjectID, Zone, Cluster, *clientOption)

	// Get the list of tables.
	tableInfo, err := client.Tables(*ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("List of Tables in Cluster: %v\n", tableInfo)
}
