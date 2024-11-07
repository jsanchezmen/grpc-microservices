package payment

import (
	"context"
	"log"

	"github.com/jsanchezmen/microservices-proto/golang/payment"
	"github.com/jsanchezmen/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	paymentClient payment.PaymentClient
	Connection    *grpc.ClientConn
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	client := payment.NewPaymentClient(conn)
	return &Adapter{paymentClient: client, Connection: conn}, nil

}

func (a *Adapter) Charge(order *domain.Order) error {
	response, err := a.paymentClient.Create(context.Background(), &payment.CreatePaymentRequest{
		UserId:     order.CustomerId,
		OrderId:    order.Id,
		TotalPrice: order.TotalPrice(),
	})
	log.Printf("Response %v", response)
	return err
}
