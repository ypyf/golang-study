package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"gopkg.in/redis.v3"
)

func openRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "10.10.149.62:6379",
		Password: "",
		DB:       0, // use default DB
	})
}

func eat() {
	consumer, err := sarama.NewConsumer([]string{"10.10.40.151:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	redis := openRedis()
	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			redis.SAdd("atime_ap", string(msg.Value))
			//log.Printf("Consumed message offset %d\n", msg.Offset)
			//log.Printf("Consumed message value %s\n", string(msg.Value))
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}

func main() {
	eat()
}
