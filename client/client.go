package main

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	"golang.org/x/net/context"

	"github.com/icrowley/fake"
	"github.com/juleur/grpc-around/chatpb"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

var (
	streamers []string
)

func init() {
	// fake 10 streamers
	for i := 0; i < 10; i++ {
		streamers = append(streamers, fake.FirstName())
	}
}

func main() {
	logrus.Println("STREAMCHAT CLIENT RUNNING ...")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		logrus.Fatalln(err)
	}
	defer conn.Close()

	client := chatpb.NewStreamChatClient(conn)

	endShowMsg := make(chan struct{})
	noMoreMsg := make(chan struct{})

	// NewChat
	makeChats(client)
	// SendMessages
	go fakeMessages(client, noMoreMsg)
	// RcvMessages
	go receiverMsg(client, endShowMsg)
	// CloseChat
	go closeChats(client, noMoreMsg)

	// when all messages were displayed, end
	<-endShowMsg
}

func makeChats(c chatpb.StreamChatClient) {
	for _, streamer := range streamers {
		_, err := c.NewChat(context.Background(), &chatpb.NewChatRequest{ChatId: streamer})
		if err != nil {
			logrus.Fatalln(err)
		}
	}
}

func fakeMessages(c chatpb.StreamChatClient, noMoreMsg chan<- struct{}) {
	for i := 0; i < 20; i++ {
		streamer := streamers[rand.Intn(len(streamers)-1)]

		fakeReq := &chatpb.SendMessagesRequest{Username: fake.FirstName(), ChatId: streamer, Message: fake.WordsN(6)}
		// fake users delay
		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		_, err := c.SendMessages(context.Background(), fakeReq)
		if err != nil {
			logrus.Fatalln(err)
		}
	}

	logrus.Warningln("WAITING 10 SECONDS")
	time.Sleep(10 * time.Second)

	for i := 0; i < 15; i++ {
		fakeReq := &chatpb.SendMessagesRequest{Username: fake.FirstName(), ChatId: streamers[rand.Intn(len(streamers)-1)], Message: fake.WordsN(6)}
		// fake users delay
		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		_, err := c.SendMessages(context.Background(), fakeReq)
		if err != nil {
			logrus.Fatalln(err)
		}
	}
	noMoreMsg <- struct{}{}
}

func receiverMsg(c chatpb.StreamChatClient, endShowMsg chan<- struct{}) {
	// fetching stream msgs from one streamer
	fakeReq := &chatpb.RcvMessagesRequest{ChatId: streamers[rand.Intn(len(streamers)-1)]}

	stream, err := c.RcvMessages(context.Background(), fakeReq)
	if err != nil {
		logrus.Fatalln(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			logrus.Infoln("No more Messages")
		}
		if err != nil {
			endShowMsg <- struct{}{}
		}
		fmt.Printf("<%s>: %s\n", res.GetUsermane(), res.GetMessage())
	}
}

func closeChats(c chatpb.StreamChatClient, noMoreMsg <-chan struct{}) {
	// Waiting that all fake messages are sent
	<-noMoreMsg
	for i := 0; i < len(streamers); i++ {
		_, err := c.CloseChat(context.Background(), &chatpb.CloseChatRequest{ChatId: streamers[i]})
		if err != nil {
			logrus.Fatalln(err)
		}
	}
}
