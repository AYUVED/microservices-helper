package middleware
import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CircuitBreaker struct {
	mu           sync.RWMutex
	failureCount int
	lastFailure  time.Time
	maxFailures  int
	timeout      time.Duration
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures: maxFailures,
		timeout:     timeout,
	}
}

func (cb *CircuitBreaker) shouldAllow() bool {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	if cb.failureCount >= cb.maxFailures {
		if time.Since(cb.lastFailure) > cb.timeout {
			return true
		}
		return false
	}
	return true
}

func (cb *CircuitBreaker) recordSuccess() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.failureCount = 0
}

func (cb *CircuitBreaker) recordFailure() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.failureCount++
	cb.lastFailure = time.Now()
}

func CircuitBreakerClientInterceptor(cb *CircuitBreaker) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if !cb.shouldAllow() {
			return status.Errorf(codes.Unavailable, "circuit breaker is open")
		}

		err := invoker(ctx, method, req, reply, cc, opts...)

		if err != nil {
			cb.recordFailure()
			return err
		}

		cb.recordSuccess()
		return nil
	}
}
