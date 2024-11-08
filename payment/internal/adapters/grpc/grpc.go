package grpc

import (
	"context"
	"fmt"

	"github.com/jsanchezmen/microservices-proto/golang/payment"
	"github.com/jsanchezmen/microservices/payment/internal/application/core/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(context context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)

	result, err := a.api.Charge(newPayment)

	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("Failed to charge. %v", err)).Err()

	}

	return &payment.CreatePaymentResponse{PaymentId: result.ID}, nil

}
