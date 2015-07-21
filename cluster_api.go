package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/cloud"
	"google.golang.org/cloud/bigtable"
	btcdpb "google.golang.org/cloud/bigtable/internal/cluster_data_proto"
	btcspb "google.golang.org/cloud/bigtable/internal/cluster_service_proto"
	"google.golang.org/grpc"
)

const clusterAdminAddr = "bigtableclusteradmin.googleapis.com:443"

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
