package main

import (
	"fmt"
	"regexp"

	"golang.org/x/net/context"
	"google.golang.org/cloud"
	"google.golang.org/cloud/bigtable"
	btcdpb "google.golang.org/cloud/bigtable/internal/cluster_data_proto"
	btcspb "google.golang.org/cloud/bigtable/internal/cluster_service_proto"
	"google.golang.org/grpc"
)

const clusterAdminAddr = "bigtableclusteradmin.googleapis.com:443"

var clusterNameRegexp = regexp.MustCompile(`^projects/([^/]+)/zones/([^/]+)/clusters/([a-z][-a-z0-9]*)$`)

type AltClusterAdminClient struct {
	conn    *grpc.ClientConn
	cClient btcspb.BigtableClusterServiceClient
	project string
}

func (cac *AltClusterAdminClient) ListZones(ctx context.Context) ([]*btcdpb.Zone, error) {
	req := &btcspb.ListZonesRequest{
		Name: "projects/" + cac.project,
	}
	res, err := cac.cClient.ListZones(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return res.Zones, nil
	}
}

// BEGIN: Copied and pasted from gcloud-golang/bigtable/admin.go
func NewAltClusterAdminClient(ctx context.Context, project string, opts ...cloud.ClientOption) (*AltClusterAdminClient, error) {
	o := []cloud.ClientOption{
		cloud.WithEndpoint(clusterAdminAddr),
		cloud.WithScopes(bigtable.ClusterAdminScope),
	}
	o = append(o, opts...)
	conn, err := cloud.DialGRPC(ctx, o...)
	if err != nil {
		return nil, fmt.Errorf("dialing: %v", err)
	}
	return &AltClusterAdminClient{
		conn:    conn,
		cClient: btcspb.NewBigtableClusterServiceClient(conn),

		project: project,
	}, nil
}

func (cac *AltClusterAdminClient) Close() {
	cac.conn.Close()
}

// Clusters returns a list of clusters in the project.
func (cac *AltClusterAdminClient) Clusters(ctx context.Context) ([]*bigtable.ClusterInfo, error) {
	req := &btcspb.ListClustersRequest{
		Name: "projects/" + cac.project,
	}
	res, err := cac.cClient.ListClusters(ctx, req)
	if err != nil {
		return nil, err
	}
	// TODO(dsymonds): Deal with failed_zones.
	var cis []*bigtable.ClusterInfo
	for _, c := range res.Clusters {
		m := clusterNameRegexp.FindStringSubmatch(c.Name)
		if m == nil {
			return nil, fmt.Errorf("malformed cluster name %q", c.Name)
		}
		cis = append(cis, &bigtable.ClusterInfo{
			Name:        m[3],
			Zone:        m[2],
			DisplayName: c.DisplayName,
			ServeNodes:  int(c.ServeNodes),
		})
	}
	return cis, nil
}

//   END: Copied and pasted from gcloud-golang/bigtable/admin.go
