package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/viogami/FavAni/config"
	"github.com/viogami/FavAni/database"
	pb "github.com/viogami/FavAni/pb/gcn"
)

func NewGRPCClient(conf *config.GRPCConfig) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(conf.Host+":"+fmt.Sprint(conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func GCN_request(conf *config.GRPCConfig, r *pb.GCNRequest) (*pb.GCNResult, error) {
	// 创建一个gRPC连接
	conn, err := NewGRPCClient(conf)
	if err != nil {
		return nil, err
	}
	// 创建一个客户端
	client := pb.NewGCNServiceClient(conn)
	defer conn.Close()
	// 发送请求并接收响应
	result, err := client.ProcessGCN(context.Background(), r)
	if err != nil {
		// Error calling ProcessGraph
		return nil, err
	}
	return result, err
}

// 获取redis缓存数据
func CheckRequestFromRedis(rdb *database.RedisDB, r *pb.GCNRequest) (map[string]float32, error) {
	// 生成key
	key := r.String()
	// 检查redis中是否有缓存
	if rdb.Client.Exists(context.Background(), key).Val() == 0 {
		log.Println("The key does not exist in the redis database")
		// redis中缓存数据
		err := rdb.HSet(key, "Graph", r.Graph.String())
		if err != nil {
			return nil, err
		}
		err = rdb.HSet(key, "Params", r.Params.String())
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	var result string
	err := rdb.HGet(key, "NodeScores", &result)
	if err != nil {
		log.Println("Failed to get the NodeScores from the redis database")
		// 有key，但是没有NodeScores字段，删除key
		return nil, rdb.HDel(key, "graph", "params", "NodeScores")
	}

	scores := make(map[string]float32)
	err = json.Unmarshal([]byte(result), &scores)
	if err != nil {
		return nil, err
	}

	return scores, nil
}

// redis缓存请求数据和结果
func SetResultToRedis(rdb *database.RedisDB, rkey string, res *pb.GCNResult) error {
	if rdb.Client.HExists(context.Background(), rkey, "NodeScores").Val() {
		log.Println("The key of field 'NodeScores' already exists in the redis database")
		return nil
	}
	NodeScores, err := json.Marshal(res.NodeScores)
	if err != nil {
		return err
	}
	// redis中缓存数据
	return rdb.HSet(rkey, "NodeScores", NodeScores)
}
