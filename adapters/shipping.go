package adapters

import (
	"context"
	"log"
	"time"

	"github.com/ayuved/microservices-helper/domain"
	"github.com/ayuved/microservices-helper/middleware"
	"github.com/ayuved/microservices-proto/golang/shipping"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type shippingAdapter struct {
	shipping shipping.ShippingClient
}

func NewShippingAdapter(shippingServiceUrl string) (*shippingAdapter, error) {
	cb := middleware.NewCircuitBreaker(5, 1*time.Minute)

	log.Println("NewShippingAdapter", shippingServiceUrl)
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(middleware.CircuitBreakerClientInterceptor(cb)),
	)
	log.Println("NewShippingAdapter", opts)
	conn, err := grpc.NewClient(shippingServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	log.Println("NewShippingAdapter", conn)
	client := shipping.NewShippingClient(conn)
	//defer conn.Close()
	return &shippingAdapter{shipping: client}, nil
}

func (a *shippingAdapter) CreateShipping(ctx context.Context, o *domain.Shipping) (*shipping.CreateShippingResponse, error) {
	var items []*shipping.ShippingItem
	for _, item := range o.ShippingItems {
		items = append(items, &shipping.ShippingItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	result, err := a.shipping.Create(ctx, &shipping.CreateShippingRequest{
		OrderId:      o.OrderID,
		AddressId:    o.AddressID,
		ShippingItems: items,

	})
	return result, err
}
