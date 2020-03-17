package api

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"sync"
)

// AppID .
var clientAppID = "direct://default/127.0.0.1:9000"

func SetAppId(appID string) {
	once := sync.Once{}
	once.Do(func() {
		if appID != "" {
			clientAppID = appID
		}
	})
}

var tls  credentials.TransportCredentials
func AddTlsClient(pemPath, serverName string) {
	once := sync.Once{}
	once.Do(func() {
		var err error
		tls, err = credentials.NewClientTLSFromFile(pemPath, serverName)
		if err != nil {
			log.Error("credentials.NewClientTLSFromFile err: %v", err)
			panic(err)
		}
	})
}

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (DemoClient, error) {
	if tls != nil {
		opts = append(opts, grpc.WithTransportCredentials(tls))
	}
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), clientAppID)
	if err != nil {
		return nil, err
	}
	return NewDemoClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm --swagger api.proto
