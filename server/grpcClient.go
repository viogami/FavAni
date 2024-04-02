package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/viogami/FavAni/config"
	pb "github.com/viogami/FavAni/pb/gcn"
)

func NewGRPCClient(conf *config.GRPCConfig) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(conf.Host+":"+fmt.Sprint(conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func GCN_request(conn *grpc.ClientConn, r *pb.GCNRequest) (map[string]float32, error) {
	// 创建一个客户端
	client := pb.NewGCNServiceClient(conn)
	defer conn.Close()
	// 发送请求并接收响应
	result, err := client.ProcessGCN(context.Background(), r)
	if err != nil {
		// Error calling ProcessGraph
		return nil, err
	}
	return result.NodeScores, err
}
