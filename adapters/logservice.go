package adapters

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices-proto/golang/logservice"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	// "google.golang.org/protobuf/proto"
	// "google.golang.org/protobuf/types/known/anypb"
)

type logserviceAdapter struct {
	logservice logservice.LogClient
}

func NewLogServiceAdapter(orderServiceUrl string) (*logserviceAdapter, error) {
	log.Println("NewLogServiceAdapter", orderServiceUrl)
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
			grpc_retry.WithMax(5),
			grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second))),
		))
	log.Println("NewLogServiceAdapter", opts)

	conn, err := grpc.NewClient(orderServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	log.Println("NewLogServiceAdapter", conn)
	client := logservice.NewLogClient(conn)
	//defer conn.Close()
	log.Println("NewLogServiceAdapter", client)
	return &logserviceAdapter{logservice: client}, nil
}

func (a *logserviceAdapter) AddLog(ctx context.Context, o *domain.Logservice) error {
	//data := map[string] o.Data
	jsonBytes, err := json.Marshal(o.Data)
	if err != nil {
		log.Fatalf("Failed to marshal data to JSON: %v", err)
	  }


	_, err = a.logservice.Add(ctx, &logservice.CreateLogRequest{
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
// func interfaceToProtoBytes(data interface{}) ([]byte, error) {
// 	// Convert the data to a protobuf Any message
// 	// anyMsg, err := anypb.New(&any.Any{Value: []byte(data.(string))})
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// Marshal the Any message to bytes
// 	return proto.Marshal(data)
// }