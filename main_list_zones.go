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

	altClient, err := NewAltClusterAdminClient(
		*ctx, ProjectID, *clientOption)
	if err != nil {
		log.Fatal(err)
		return
	}

	zones, err := altClient.ListZones(*ctx)
	for i := 0; i < len(zones); i++ {
		fmt.Printf("Zone %d\n", i)
		fmt.Println("===========================================================")
		fmt.Printf("       Name: %v\n", zones[i].Name)
		fmt.Printf("DisplayName: %v\n", zones[i].DisplayName)
		fmt.Printf("     Status: %v\n", zones[i].Status)
	}
}
