package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/chen2eric/tag-service/proto"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/naming"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx, "tag-service", nil)
	defer clientConn.Close()
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{})
	log.Printf("resp : %v\n", resp)
}

func GetClientConn(ctx context.Context, serviceName string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	config := clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: time.Second * 60,
	}
	cli, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}
	r := &naming.GRPCResolver{Client: cli}
	target := fmt.Sprintf("/etcdv3://gp-programming-tour/grpc/%s", serviceName)
	opts = append(opts, grpc.WithInsecure(), grpc.WithBalancer(grpc.RoundRobin(r)), grpc.WithBlock())

	return grpc.DialContext(ctx, target, opts...)
}
