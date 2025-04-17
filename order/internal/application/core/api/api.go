package api

import (
	"context"
	"fmt"

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
	fmt.Println("LINE 27 ====================")
	fmt.Printf("order: %v\n", order)
	paymentErr := a.payment.Charge(ctx, &order)
	fmt.Printf("123order: %v\n", order)
	if paymentErr != nil {
		println(paymentErr.Error())
		return domain.Order{}, paymentErr
	}

	return order, nil
}
