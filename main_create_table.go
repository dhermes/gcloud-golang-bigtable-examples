package main

import (
	"fmt"
	"log"

	"github.com/dhermes/gcloud-golang/bigtable"
)

func main() {
	ctx, clientOption, err := getClientArgs(bigtable.AdminScope)
	if err != nil {
		log.Fatal(err)
		return
	}

	client, err := bigtable.NewAdminClient(
		*ctx, ProjectID, Zone, Cluster, *clientOption)

	// Insert a new table.
	tableName := "omg-finally"
	err = client.CreateTable(*ctx, tableName)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Table created: %v\n", tableName)

	// Get the list of tables.
	tableInfo, err := client.Tables(*ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("List of Tables in Cluster: %v\n", tableInfo)
}
