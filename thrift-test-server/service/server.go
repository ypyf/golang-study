package service

import (
	"crypto/tls"
	"git.apache.org/thrift.git/lib/go/thrift"
	"wifi/session/rpc"
	"github.com/Sirupsen/logrus"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TServerTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		if cert, err := tls.LoadX509KeyPair("server.crt", "server.key"); err == nil {
			cfg.Certificates = append(cfg.Certificates, cert)
		} else {
			return err
		}
		transport, err = thrift.NewTSSLServerSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTServerSocket(addr)
	}

	if err != nil {
		return err
	}
	logrus.Debugf("%T\n", transport)
	handler := NewSessionHandler()
	processor := rpc.NewSessionManagerProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	logrus.Info("Starting the simple server... on ", addr)
	return server.Serve()
}

func RunSSLServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	return runServer(transportFactory, protocolFactory, addr, true)
}

func RunNonSecureServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	return runServer(transportFactory, protocolFactory, addr, false)
}
