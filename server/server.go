package main

import (
	"net"
	"sync"

	"golang.org/x/net/context"

	"github.com/grpc/grpc-go/status"
	"github.com/juleur/grpc-around/chatpb"
	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
)

// MessengerServer struct
type MessengerServer struct {
	sync.Mutex
	Chats map[string]chan chatpb.ChatDetails
}

// NewMessenger func
func NewMessenger() *MessengerServer {
	newMessengerServer := &MessengerServer{
		Chats: make(map[string]chan chatpb.ChatDetails),
	}
	return newMessengerServer
}

// NewChat func
func (ms *MessengerServer) NewChat(ctx context.Context, req *chatpb.NewChatRequest) (*chatpb.Void, error) {
	logrus.Infof("NewChat: %v\n", req)

	ms.Chats[req.GetChatId()] = make(chan chatpb.ChatDetails, 1000)

	return &chatpb.Void{}, nil
}

// SendMessages func
func (ms *MessengerServer) SendMessages(ctx context.Context, req *chatpb.SendMessagesRequest) (*chatpb.Void, error) {
	logrus.Infof("SendMessagesRequest: %v\n", *req)

	ms.Lock()
	_, ok := ms.Chats[req.GetChatId()]
	if !ok {
		return &chatpb.Void{}, status.Errorf(codes.NotFound|codes.Canceled, "Chat n°%s cannot be found or may have been closed !\n", req.GetChatId())
	}
	ms.Unlock()

	go func() {
		ms.Chats[req.GetChatId()] <- chatpb.ChatDetails{
			Username: req.GetUsername(),
			ChatId:   req.GetChatId(),
			Message:  req.GetMessage(),
		}
	}()
	return &chatpb.Void{}, nil
}

// RcvMessages func
func (ms *MessengerServer) RcvMessages(req *chatpb.RcvMessagesRequest, stream chatpb.StreamChat_RcvMessagesServer) error {
	logrus.Infof("RcvMessagesRequest: %v\n", *req)

	ms.Lock()
	_, ok := ms.Chats[req.GetChatId()]
	if !ok {
		return status.Errorf(codes.NotFound|codes.Canceled, "Chat n°%s cannot be found or may have been closed !\n", req.GetChatId())
	}
	ms.Unlock()

	for data := range ms.Chats[req.GetChatId()] {
		stream.Send(&chatpb.RcvMessagesResponse{
			Usermane: data.GetUsername(),
			Message:  data.GetMessage(),
		})
	}
	return nil
}

// CloseChat func
func (ms *MessengerServer) CloseChat(ctx context.Context, req *chatpb.CloseChatRequest) (*chatpb.Void, error) {
	logrus.Infof("CloseChat: %v\n", req)

	close(ms.Chats[req.GetChatId()])

	return &chatpb.Void{}, nil
}

func main() {
	logrus.Println("STREAMCHAT SERVER RUNNING ...")

	ms := NewMessenger()

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		logrus.Fatalln(err)
	}

	s := grpc.NewServer()
	chatpb.RegisterStreamChatServer(s, ms)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalln(err)
	}
}
