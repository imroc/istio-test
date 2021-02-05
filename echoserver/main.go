package main

import (
	"context"
	"fmt"
	"github.com/imroc/istio-test/chat"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

var (
	GRPCRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",
			Help: "Number of the grpc requests received since the server started",
		},
		[]string{"body"},
	)
)

func init() {
	prometheus.MustRegister(GRPCRequests)
}

type server struct {
	chat.UnimplementedChatServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *chat.Message) (*chat.Message, error) {
	log.Printf("Received: %v", in.GetBody())
	GRPCRequests.WithLabelValues(in.GetBody()).Inc()
	return &chat.Message{Body: "hi"}, nil
}

func main() {

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println(http.ListenAndServe(":9001", nil))
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("echoserver!")

	s := server{}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
