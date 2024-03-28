package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/viogami/FavAni/proto"
)

func GCN_request() (map[string]float32, error) {
	conn, err := grpc.Dial("localhost:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// Failed to connect to gcnserver(py) by grpc
		return nil, err
	}
	defer conn.Close()

	client := pb.NewGCNServiceClient(conn)

	// 创建一个实例的图数据
	G_example := &pb.GraphData{
		Nodes: []*pb.Node{
			{Id: "CN-AH", Features: []float32{5000}},
			{Id: "CN-BJ", Features: []float32{6000}},
			{Id: "CN-JS", Features: []float32{7000}},
		},
		Edges: []*pb.Edge{
			{SourceId: "CN-AH", TargetId: "CN-BJ"},
			{SourceId: "CN-BJ", TargetId: "CN-AH"},
			{SourceId: "CN-AH", TargetId: "CN-JS"},
			{SourceId: "CN-JS", TargetId: "CN-AH"},
		},
	}

	// 发送请求并接收响应
	result, err := client.ProcessGraph(context.Background(), G_example)
	if err != nil {
		// Error calling ProcessGraph
		return nil, err
	}
	return result.NodeScores, err
}
