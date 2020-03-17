package api

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/naming/discovery"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/bilibili/kratos/pkg/net/rpc/warden/resolver"
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

var tls credentials.TransportCredentials

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

func init() {
	// NOTE: 注意这段代码，表示要使用discovery进行服务发现
	// NOTE: 还需注意的是，resolver.Register是全局生效的，所以建议该代码放在进程初始化的时候执行
	// NOTE: ！！！切记不要在一个进程内进行多个不同中间件的Register！！！
	// NOTE: 在启动应用时，可以通过flag(-discovery.nodes) 或者 环境配置(DISCOVERY_NODES)指定discovery节点
	resolver.Register(discovery.Builder())
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
