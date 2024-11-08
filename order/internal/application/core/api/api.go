package api

import (
	"log"

	"github.com/jsanchezmen/microservices/order/internal/application/core/domain"
	"github.com/jsanchezmen/microservices/order/internal/ports"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	log.Print("Calling payment service")
	paymentErr := a.payment.Charge(&order)
	log.Print("paymentErr", paymentErr)
	if paymentErr != nil {
		return domain.Order{}, err
	}

	return order, nil
}
