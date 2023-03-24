package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

const MAX_WAIT_SEC = 20

const END_KEY = "END"

type Env struct {
	Brokers []string
	Topic   string
	Id      string
	OutDir  string
	LogFile string
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
	outDir := MustGetEnv("OUT_DIR")
	logFile := MustGetEnv("LOG_FILE")
	if len(brokers) == 0 {
		panic("no broker address")
	}

	return Env{Brokers: brokers, Topic: topic, Id: id, OutDir: outDir, LogFile: logFile}
}

func main() {
	env := getEnv()
	consume(context.Background(), env)
}

func consume(ctx context.Context, env Env) {

	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	consumerName := "consumer %" + env.Id
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     env.Brokers,
		Topic:       env.Topic,
		GroupID:     "my-group",
		StartOffset: kafka.FirstOffset,
		// FirstOffset: will start consuming messages from the earliest available
		// LastOffset: will only consume new messages (this only applies for new consumer groups.)
		Logger: log.New(os.Stdout, fmt.Sprintf("kafka %s: ", consumerName), 0),
	})
	f, err := os.Create(env.OutDir + env.LogFile + "-" + env.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for i := 0; ; i++ {
		ctx_timeout, _ := context.WithTimeout(ctx, time.Second*MAX_WAIT_SEC) // timeout
		msg, err := r.ReadMessage(ctx_timeout)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
		if string(msg.Value) == END_KEY {
			break
		}
		log := fmt.Sprintf("%d  %s\n", time.Now().Unix(), msg.Value)
		f.WriteString(log)
		// after receiving the message, log its value
	}
}
