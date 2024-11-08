package ports

import "github.com/jsanchezmen/microservices/payment/internal/application/core/domain"

type APIPort interface {
	Charge(payment domain.Payment) (domain.Payment, error)
}
