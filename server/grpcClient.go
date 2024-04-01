package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/viogami/FavAni/proto"
)

func GCN_request(G_example *pb.GraphData) (map[string]float32, error) {
	conn, err := grpc.Dial("localhost:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// Failed to connect to gcnserver(py) by grpc
		return nil, err
	}
	defer conn.Close()

	client := pb.NewGCNServiceClient(conn)

	// 发送请求并接收响应
	result, err := client.ProcessGraph(context.Background(), G_example)
	if err != nil {
		// Error calling ProcessGraph
		return nil, err
	}
	return result.NodeScores, err
}
