package main

import (
	"flag"
	"git.apache.org/thrift.git/lib/go/thrift"
	"thrift-test-server/service"
	"github.com/Sirupsen/logrus"
)

func main() {
	var networkAddr = flag.String("addr", "localhost:9090", "会话服务器地址")
	flag.Parse()

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	if err := service.RunNonSecureServer(transportFactory, protocolFactory, *networkAddr); err != nil {
		logrus.Fatal(err)
	}
}
