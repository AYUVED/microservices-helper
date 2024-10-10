package adapters

import (
	"context"
	"log"
	"time"
	
	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices-proto/golang/logservice"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
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

	_, err := a.logservice.Add(ctx, &logservice.CreateLogRequest{
		App:  o.App,
		Name: o.Name,
		Type: o.Type,
		Status: o.Status,
		ProcessId: o.ProcessId,
		Data: o.Data.([]byte),
		User: o.User,

	})
	return err
}
