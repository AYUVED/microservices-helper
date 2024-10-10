module github.com/ayuved/microservices-helper

go 1.22.3

require (
	github.com/ayuved/microservices-proto/golang/logservice v1.0.9
	github.com/ayuved/microservices-proto/golang/order v1.0.9
	github.com/ayuved/microservices-proto/golang/shipping v1.0.9
	google.golang.org/grpc v1.66.2
)

require github.com/stretchr/testify v1.9.0 // indirect

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
