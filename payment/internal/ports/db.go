package ports

import "github.com/jsanchezmen/microservices/payment/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Payment, error)
	Save(payment *domain.Payment) error
}
