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
var serverAddr string
var verbose bool

func init() {
	flag.UintVar(&retryCount, "retry", defaultRetryCount, "Retry count of request")
	flag.DurationVar(&timeout, "timeout", defaultTimeout*time.Second, "Request timeout")
	flag.StringVar(&serverAddr, "addr", "", "The server address in the format of host:port")
	flag.BoolVar(&verbose, "verbose", false, "verbose mode")
}

func sendManualStartRequest(addr string, timeout time.Duration, retryCount uint) {
	if verbose {
		log.Printf("Connecting %s", addr)
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()

	c := pb.NewAppClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if verbose {
		log.Printf("Requesting ManualStartRequest to %s", addr)
	}
	_, err = c.ManualStart(ctx, &pb.ManualStartRequest{}, grpc_retry.WithMax(retryCount))
	if err != nil {
		log.Fatalf("ManualStartRequest failed: %v", err)
		os.Exit(1)
	}

	log.Println("ManualStartRequest succeeded")
}

func main() {
	flag.Parse()

	if serverAddr == "" {
		fmt.Fprintln(os.Stderr, "-addr must be specified")
		flag.Usage()
	}

	sendManualStartRequest(serverAddr, timeout, retryCount)
}
