package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/client"
)

func main() {
	cfg := client.Config{
		Endpoints: []string{"http://localhost:2379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	ctx := context.Background()
	resp, err := kapi.Get(ctx, "/foo", nil)
	//resp, err := kapi.Set(context.Background(), "foo", "bar", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Node.Key, resp.Node.Value)
}
