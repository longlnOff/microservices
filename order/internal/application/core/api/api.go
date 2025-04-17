package api

import (
	"context"

	"github.com/longlnOff/microservices/order/internal/application/core/domain"
	"github.com/longlnOff/microservices/order/internal/ports"
)

type Application struct {
	db ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db: db,
		payment: payment,
	}
}
func (a Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	err := a.db.Save(ctx, &order)
	if err != nil {
		return domain.Order{}, err
	}
	paymentErr := a.payment.Charge(ctx, &order)
	if paymentErr != nil {
		return domain.Order{}, paymentErr
	}

	return order, nil
}
