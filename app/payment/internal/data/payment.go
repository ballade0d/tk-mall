package data

import (
	"context"
	"errors"
	"mall/ent"
	"mall/ent/order"
)

type PaymentRepo struct {
	data *Data
}

func NewPaymentRepo(data *Data) PaymentRepo {
	return PaymentRepo{data: data}
}

func (r *PaymentRepo) PayOrder(ctx context.Context, orderID int) (*ent.Payment, error) {
	o, err := r.data.db.Order.Query().Where(order.IDEQ(orderID)).WithItems().Only(ctx)
	if err != nil {
		return nil, err
	}
	if o.Status != "pending" {
		return nil, errors.New("order status error")
	}
	var amount float32
	for _, i := range o.Edges.Items {
		amount += i.Price * float32(i.Quantity)
	}
	p, err := r.data.db.Payment.Create().SetOrder(o).SetAmount(amount).Save(ctx)
	if err != nil {
		return nil, err
	}
	return p, nil
}
