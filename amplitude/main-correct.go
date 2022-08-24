package main

import (
	"bitbucket.org/ntuclink/devex-sdk-go/logging/v2"
	"context"
	"fmt"
)

const (
	chunkSize int = 1000
)

func main() {
	fmt.Println("Hello world")
	ctx := context.Background()
	pushDateToAmplitude(ctx)
	fmt.Println("DONE!")
}

func pushDateToAmplitude(ctx context.Context) {
	orderCheckouts := GetOrderCheckouts(ctx)
	orderCheckoutsChunks := ChuckArray(orderCheckouts, chunkSize)
	for i := 0; i < len(orderCheckoutsChunks); i++ {
		request := MapEventsToAmplitudeRequests(orderCheckoutsChunks[i])
		SendBatchEvent(ctx, request)
		logResults(ctx, orderCheckoutsChunks[i])
	}
}

func logResults(ctx context.Context, orderCheckouts []OrderCheckout) {
	for i := 0; i < len(orderCheckouts); i++ {
		//todo: remove print
		fmt.Printf("Entry: %d \t customerId: %d \t loyalty card no: %s\n", i+1, orderCheckouts[i].CustomerId, orderCheckouts[i].LoyaltyCardNo)
		logging.WithCtx(ctx).Printf("Successfully pushed Amplitude request for userId: %v", orderCheckouts[i].CustomerId)
	}
}
