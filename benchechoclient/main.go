package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/imroc/istio-test/chat"
	"google.golang.org/grpc"
)

var (
	serverAddr  string        = "echoserver:9000"
	reqCount    int32         = 0
	totalCount  int32         = 50000
	timeout     time.Duration = 30 * time.Millisecond
	totalTime   time.Duration = 0
	concurrency int           = 1
	slowest     time.Duration
	wg          sync.WaitGroup
)

func main() {
	fmt.Printf("%+v\n", os.Args)
	if len(os.Args) >= 2 {
		serverAddr = os.Args[1]
	}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go req()
	}

	wg.Wait()
	fmt.Printf("slowest: %v\n", slowest)
	fmt.Printf("avg: %v\n", totalTime/time.Duration(totalCount))
	fmt.Printf("qps: %v\n", reqCount/int32(totalTime/time.Second))
}

func req() {
	defer wg.Done()
	// Set up a connection to the server.
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		if shouldStop() {
			return
		}
		c := chat.NewChatServiceClient(conn)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		id := uuid.NewString()
		before := time.Now()
		_, err := c.SayHello(ctx, &chat.Message{Body: id})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		after := time.Now()
		t := after.Sub(before)
		if t > timeout {
			log.Printf("timeout: %s cost: %v\n", id, t)
		}
		if t > slowest {
			slowest = t
		}
		atomic.AddInt32(&reqCount, 1)
		timeMu.Lock()
		totalTime += t
		timeMu.Unlock()
	}
}

var mu sync.Mutex
var timeMu sync.Mutex

func shouldStop() bool {
	mu.Lock()
	defer mu.Unlock()
	if reqCount >= totalCount {
		return true
	}
	return false
}
