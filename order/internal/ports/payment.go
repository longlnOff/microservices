package ports

import (
	"context"
	"github.com/longlnOff/microservices/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
