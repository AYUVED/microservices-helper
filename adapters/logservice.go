package adapters

import (
	"context"
	"log"

	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices-proto/golang/logservice"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
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
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	log.Println("NewLogServiceAdapter", opts)
	conn, err := grpc.Dial(orderServiceUrl, opts...)
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
		Data: o.Data,
	})
	return err
}
