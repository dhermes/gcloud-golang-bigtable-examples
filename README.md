# Example Calls with `gcloud-golang

Make sure you have [downloaded][2] `go` (GOLANG) and
that it is on your `${PATH}`. After doing this, the
`Makefile` will handle installing isolated dependencies
to run this sample.

To list clusters run

```
make run_cluster
```

This consistently fails (at least with a service account),

To list tables in a cluster

```
make run_table
```

This will need to be a cluster you have created (see
"Creating a Cluster in the UI" below).

Finally, to create a table in an existing cluster, run

```
make run_table_with_create
```

This will create a table named `omg-finally`. If you'd like
to use a different table name, you can edit

```go
tableName := "omg-finally"
```

in `main_with_table_admin_and_create.go`. However, the script
is just for demonstration, so it shouldn't matter.

## Enabling the BigTable API

1.  Visit [Google Cloud Console][1]
1.  Either create a new project or visit an existing one
1.  In the project, click **"APIs & auth > APIs"**. The URI
    should resemble

    ```
    https://console.developers.google.com/project/{project-id}/apiui/apis/library
    ```

1.  On this page, search for **bigtable**, and click both `Cloud Bigtable API`
    and `Cloud Bigtable Table Admin API`.
1.  For each API, click "Enable API" (if not already enabled)

## Getting a Service Account Keyfile

1.  Visit [Google Cloud Console][1]
1.  Either create a new project or visit an existing one
1.  In the project, click **"APIs & auth > Credentials"**. The URI
    should resemble

    ```
    https://console.developers.google.com/project/{project-id}/apiui/credential
    ```

1.  On this page, click "Create new Client ID", select "Service account" as
    your "Application type" and then download the JSON key provided. The
    downloaded file should resemble `keyfile.json.sample`.

After downloading, move this key to the local directory holding this code.

## Creating a Cluster in the UI

1.  Visit [Google Cloud Console][1]
1.  Either create a new project or visit an existing one
1.  In the project, click **"Storage > Cloud Bigtable"**. The URI
    should resemble

    ```
    https://console.developers.google.com/project/{project-id}/bigtable/clusters
    ```

1.  On this page, click **Create a cluster** and take note of the "Cluster ID"
    and "Zone" you use when creating it.

## Setting Up Local Files

You will need configuration for your own account and the code
pulls this from `consts.go`.

1.  Execute

    ```bash
    cp consts.go.sample consts.go
    ```

1.  Edit `consts.go` to match your own project

    1.  The `ProjectID` in `consts.go` to match the project ID
        in the project you used above. (Make sure you use the
        Project ID, not the Project Number)
    1.  You may name `Cluster` and `Zone` anything you like, but these
        should come from a cluster that already exists (see above for
        how to create a cluster).
    1.  Change `KeyFile` to the path of the service account key
        file that you downloaded above.

[1]: https://console.developers.google.com/
[2]: http://golang.org/doc/install
