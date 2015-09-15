package main

import (
	"log"
	"time"

	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/client"
)

func main() {
	cfg := client.Config{
                Endpoints: []string{"http://localhost:2379", "http://localhost:4001"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	resp, err := kapi.Set(context.Background(), "foo", "bar", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
