package api
//
//import (
//	"context"
//	"github.com/bilibili/kratos/pkg/conf/env"
//	"github.com/bilibili/kratos/pkg/log"
//	"github.com/bilibili/kratos/pkg/naming"
//	"github.com/bilibili/kratos/pkg/naming/discovery"
//	"github.com/bilibili/kratos/pkg/naming/etcd"
//	"github.com/bilibili/kratos/pkg/net/rpc/warden"
//	"github.com/bilibili/kratos/pkg/net/rpc/warden/resolver"
//	"go.etcd.io/etcd/clientv3"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials"
//	"sync"
//	"time"
//)
//
//// AppID .
//var clientAppID = "direct://default/127.0.0.1:9000"
//
//func SetAppId(appID string) {
//	once := sync.Once{}
//	once.Do(func() {
//		if appID != "" {
//			clientAppID = appID
//		}
//	})
//}
//
//var tls credentials.TransportCredentials
//
//func AddTlsClient(pemPath, serverName string) {
//	once := sync.Once{}
//	once.Do(func() {
//		var err error
//		tls, err = credentials.NewClientTLSFromFile(pemPath, serverName)
//		if err != nil {
//			log.Error("credentials.NewClientTLSFromFile err: %v", err)
//			panic(err)
//		}
//	})
//}
//
//// NewClient new grpc client
//func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (DemoClient, error) {
//	if tls != nil {
//		opts = append(opts, grpc.WithTransportCredentials(tls))
//	}
//	config := &clientv3.Config{
//		Endpoints:   []string{"192.168.3.241.1:2379"},
//		DialTimeout: time.Second * 3,
//		DialOptions: []grpc.DialOption{grpc.WithBlock()},
//	}
//
//	client := warden.NewClient(cfg, opts...)
//	cc, err := client.Dial(context.Background(), clientAppID)
//	if err != nil {
//		return nil, err
//	}
//	return NewDemoClient(cc), nil
//}
//func init()  {
//	config := &clientv3.Config{
//		Endpoints:   []string{"192.168.3.241.1:2379"},
//		DialTimeout: time.Second * 3,
//		DialOptions: []grpc.DialOption{grpc.WithBlock()},
//	}
//	builder, err := etcd.New(config)
//	if err != nil {
//		panic(err)
//	}
//	demoService := builder.Build("demo.service")
//	demoService.Watch()
//}

//type consumer struct {
//	conf *discovery.Config
//	appID string
//	dis naming.Resolver
//	ins []*naming.Instance
//}
//func ResolverWatch()  {
//	conf := &discovery.Config{
//		Nodes: []string{"192.168.3.241:7171"},
//		Zone:  "sh1",
//		Env:   env.DeployEnv,
//	}
//	dis := discovery.New(conf)
//	c := &consumer{
//		conf:  conf,
//		appID: "demo.server",
//		dis:   dis.Build("demo.server"),
//		ins:   nil,
//	}
//	rsl := dis.Build(c.appID)
//	ch := rsl.Watch()
//
//}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm --swagger api.proto
