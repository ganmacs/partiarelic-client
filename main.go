package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/ganmacs/partiarelic-client/partiarelic"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	defaultTimeout    = 1
	defaultRetryCount = 3
)

var retryCount uint
var timeout time.Duration
var requestUrl string

func init() {
	flag.UintVar(&retryCount, "retry", defaultRetryCount, "Retry count of request")
	flag.DurationVar(&timeout, "timeout", defaultTimeout*time.Second, "Request timeout")
	flag.StringVar(&requestUrl, "url", "", "URL to request")
}

func sendManualStartRequest(url string, timeout time.Duration, retryCount uint) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()

	c := pb.NewAppClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err = c.ManualStart(ctx, &pb.ManualStartRequest{}, grpc_retry.WithMax(retryCount))
	if err != nil {
		log.Fatalf("ManualStartRequest failed: %v", err)
		os.Exit(1)
	}

	log.Println("ManualStartRequest succeeded")
}

func main() {
	flag.Parse()

	if requestUrl == "" {
		fmt.Fprintln(os.Stderr, "-url must be specified")
		flag.Usage()
	}

	sendManualStartRequest(requestUrl, timeout, retryCount)
}
