package logservice

import (
	"context"

	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices-proto/golang/logservice"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	logservice logservice.LogClient
}

func NewLogServiceAdapter(orderServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	conn, err := grpc.Dial(orderServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := logservice.NewLogClient(conn)
	//defer conn.Close()
	return &Adapter{logservice: client}, nil
}

func (a *Adapter) LogService(ctx context.Context, o *domain.Logservice) error {

	_, err := a.logservice.Add(ctx, &logservice.CreateLogRequest{
		App:  o.App,
		Name: o.Name,
		Data: o.Data,
	})
	return err
}
