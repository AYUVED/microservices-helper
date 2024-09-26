package adapters

import (
	"context"
	"log"
	"time"

	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices-helper/middleware"
	"github.com/ayuved/microservices-proto/golang/order"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type orderAdapter struct {
	order order.OrderClient
}

func NewOrderAdapter(orderServiceUrl string) (*orderAdapter, error) {
	cb := middleware.NewCircuitBreaker(5, 1*time.Minute)

	log.Println("NewOrderAdapter", orderServiceUrl)
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(middleware.CircuitBreakerClientInterceptor(cb)),
	)
	log.Println("NewOrderAdapter", opts)
	conn, err := grpc.NewClient(orderServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	log.Println("NewOrderAdapter", conn)
	client := order.NewOrderClient(conn)
	//defer conn.Close()
	return &orderAdapter{order: client}, nil
}

func (a *orderAdapter) CreateOrder(ctx context.Context, o *domain.Order) (*order.CreateOrderResponse, error) {
	var items []*order.OrderItem
	for _, item := range o.OrderItems {
		items = append(items, &order.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	result, err := a.order.Create(ctx, &order.CreateOrderRequest{
		UserId:     o.CustomerID,
		OrderItems: items,
	})
	return result, err
}
