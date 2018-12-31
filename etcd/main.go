package main

import (
	"context"
	"fmt"
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
	resp, err := kapi.Set(context.Background(), "foo", "123", nil)
	if err != nil {
		log.Fatalf("Couldn't set node: %v\n", err)
	}
	resp, err = kapi.Get(ctx, "foo", nil)
	if err != nil {
		log.Fatalf("Couldn't get node: %v\n", err)
	}
	fmt.Println(resp.Node.Key, resp.Node.Value)
}
