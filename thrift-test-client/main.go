package main

import (
	"os"
	"time"
	"wifi/session/rpc"

	"github.com/Sirupsen/logrus"

	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	NETWORK_ADDR = "127.0.0.1:9090"
)

func main() {
	startTime := currentTimeMillis()

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(NETWORK_ADDR)
	if err != nil {
		logrus.Fatal(os.Stderr, "error resolving address:", err)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewSessionManagerClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		logrus.Fatal(os.Stderr, "Error opening socket to "+NETWORK_ADDR, err)
	}
	defer transport.Close()

	// 开始调用服务的接口
	ctx := rpc.NewSessionContext()

	sid, _ := client.CreateSession(ctx)
	logrus.Infof("创新新的会话id => %s\n", sid)

	ctx, _ = client.GetSession(sid)
	logrus.Infof("获取会话上下文 => %+v\n", ctx)

	endTime := currentTimeMillis()
	logrus.Infof("本次调用用时: %d 毫秒\n", endTime-startTime)

}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
