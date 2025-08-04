package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client   *mongo.Client
	database *mongo.Database
)

// Init 初始化MongoDB连接
func Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// 连接MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	
	// 测试连接
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	
	// 设置数据库
	database = client.Database("random_today")
	
	log.Println("MongoDB connected successfully")
	return nil
}

// GetClient 获取MongoDB客户端
func GetClient() *mongo.Client {
	return client
}

// GetDatabase 获取数据库实例
func GetDatabase() *mongo.Database {
	return database
}

// Close 关闭数据库连接
func Close() error {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return client.Disconnect(ctx)
	}
	return nil
} 