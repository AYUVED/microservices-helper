package adapters

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices-proto/golang/eventEmitter"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"

)

type eventEmitterAdapter struct {
	eventEmitter eventEmitter.EventEmitterClient
}

func NewEventEmitterAdapter(orderServiceUrl string) (*eventEmitterAdapter, error) {
	log.Println("NewEventEmitterAdapter", orderServiceUrl)
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
			grpc_retry.WithMax(5),
			grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second))),
		))
	log.Println("NewEventEmitterAdapter", opts)

	conn, err := grpc.NewClient(orderServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	log.Println("NewEventEmitterAdapter", conn)
	client := eventEmitter.NewEventEmitterClient(conn)
	//defer conn.Close()
	log.Println("NewEventEmitterAdapter", client)
	return &eventEmitterAdapter{eventEmitter: client}, nil
}

func (a *eventEmitterAdapter) AddLog(ctx context.Context, o *domain.Logservice) error {
	jsonBytes, err := json.Marshal(o.Data)
	if err != nil {
		log.Fatalf("Failed to marshal data to JSON: %v", err)
	  }

	_, err = a.eventEmitter.AddLogEvent(ctx, &eventEmitter.CreateLogEventRequest{
		App:  o.App,
		Name: o.Name,
		Type: o.Type,
		Status: o.Status,
		ProcessId: o.ProcessId,
		Data: string(jsonBytes),
		User: o.User,

	})
	return err
}
