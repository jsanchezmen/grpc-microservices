package db

import (
	"github.com/jsanchezmen/microservices/payment/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderId    int64
	TotalPrice float32
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})

	if openErr != nil {
		return nil, openErr
	}

	if err := db.AutoMigrate(&Payment{}); err != nil {
		return nil, err
	}
	return &Adapter{db: db}, nil
}

func (a Adapter) Get(id string) (domain.Payment, error) {
	var paymentEntity Payment
	res := a.db.First(&paymentEntity, id)
	paymentDomain := domain.Payment{
		ID:         int64(paymentEntity.ID),
		CustomerID: paymentEntity.CustomerID,
		Status:     paymentEntity.Status,
		OrderId:    paymentEntity.OrderId,
		TotalPrice: paymentEntity.TotalPrice,
		CreatedAt:  paymentEntity.CreatedAt.UnixNano(),
	}

	return paymentDomain, res.Error
}

func (a Adapter) Save(payment *domain.Payment) error {

	paymentEntity := Payment{
		CustomerID: payment.CustomerID,
		Status:     payment.Status,
		OrderId:    payment.OrderId,
		TotalPrice: payment.TotalPrice,
	}

	res := a.db.Create(&paymentEntity)

	if res.Error == nil {
		payment.ID = int64(paymentEntity.ID)
	}
	return res.Error
}
