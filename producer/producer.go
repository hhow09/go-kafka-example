package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

const END_KEY = "END"

type Env struct {
	Brokers        []string
	Topic          string
	Id             string
	MaxCount       int
	Async          bool
	BatchTimeoutMs int
}

func GetEnvInt(s string, _default int) int {
	v := os.Getenv(s)
	iv, err := strconv.Atoi(v)
	if err != nil {
		fmt.Printf("%s not provided, use default %d\n", s, _default)
		iv = _default
	}
	return iv
}

func MustGetEnv(s string) string {
	v := os.Getenv(s)
	if v == "" {
		log.Fatalf("no %s", s)
	}
	return v
}

func getEnv() Env {
	brokers := strings.Split(os.Getenv("BROKER_ADDRESS"), ",")
	topic := MustGetEnv("TOPIC")
	id := MustGetEnv("ID")
	maxCount := GetEnvInt("MAX_COUNT", 100)
	async := os.Getenv("ASYNC")
	batchTimeoutMs := GetEnvInt("BATCH_TIMEOUT_MS", 1000)
	if len(brokers) == 0 {
		panic("no broker address")
	}
	return Env{Brokers: brokers, Topic: topic, Id: id, MaxCount: maxCount, Async: async == "true", BatchTimeoutMs: batchTimeoutMs}
}

func produce(ctx context.Context, env Env, msgChan chan string) {
	producerName := fmt.Sprintf("kafka producer %s, ", env.Id)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      env.Brokers,
		Topic:        env.Topic,
		Logger:       log.New(os.Stdout, producerName, 0),
		Async:        env.Async,
		BatchTimeout: time.Duration(env.BatchTimeoutMs) * time.Millisecond,
	})

	// send completed callback
	w.Completion = func(messages []kafka.Message, err error) {
		if err != nil {
			fmt.Println("callback error:", err)
		}
	}
	defer w.Close()
	for msgKey := range msgChan {
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(msgKey),
			// create an arbitrary message payload for the value
			Value: []byte("message" + msgKey),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("write to writer:", msgKey)
	}
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(env.Id + "-" + strconv.Itoa(env.MaxCount+1)),
		// create an arbitrary message payload for the value
		Value: []byte(END_KEY),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func getMessageChan(env Env) chan string {
	ch := make(chan string, env.MaxCount)
	for i := 0; i < env.MaxCount; i++ {
		msgKey := env.Id + "-" + strconv.Itoa(i)
		ch <- msgKey
	}
	return ch
}

func main() {
	env := getEnv()
	fmt.Printf("producer env %+v \n", env)
	ch := getMessageChan(env)
	close(ch)
	produce(context.Background(), env, ch)
}
