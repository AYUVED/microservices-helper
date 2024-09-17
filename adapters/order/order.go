package Order

import (
	"context"

	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices-proto/golang/order"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	order order.OrderClient
}

func NewOrderAdapter(orderServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	conn, err := grpc.Dial(orderServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := order.NewOrderClient(conn)
	//defer conn.Close()
	return &Adapter{order: client}, nil
}

func (a *Adapter) Order(ctx context.Context, o *domain.Order) error {
	var items []*order.OrderItem
	for _, item := range o.OrderItems {
		items = append(items, &order.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	_, err := a.order.Create(ctx, &order.CreateOrderRequest{
		UserId:     o.CustomerID,
		OrderItems: items,
	})
	return err
}
