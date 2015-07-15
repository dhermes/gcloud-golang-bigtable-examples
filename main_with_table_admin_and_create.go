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

	config, err := google.JWTConfigFromJSON(jsonKey, bigtable.AdminScope)
	if err != nil {
		log.Fatal(err)
		return
	}

	ctx := cloud.NewContext(ProjectID, config.Client(oauth2.NoContext))

	client, err := bigtable.NewAdminClient(
		ctx, ProjectID, Zone, Cluster,
		cloud.WithTokenSource(config.TokenSource(ctx)))

	// Insert a new table.
	tableName := "omg-finally"
	err = client.CreateTable(ctx, tableName)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Table created: %v\n", tableName)

	// Get the list of tables.
	tableInfo, err := client.Tables(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("List of Tables in Cluster: %v\n", tableInfo)
}
