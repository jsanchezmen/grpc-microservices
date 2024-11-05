package ports

import "github.com/jsanchezmen/microservices/order/internal/application/core/domain"

type ApiPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
