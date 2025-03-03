package data

import (
	"context"
	"mall/ent/payment"
)

type PaymentRepo struct {
	data *Data
}

func NewPaymentRepo(data *Data) *PaymentRepo {
	return &PaymentRepo{data: data}
}

func (r *PaymentRepo) PaymentCallback(ctx context.Context, id int) error {
	p, err := r.data.db.Payment.Query().Where(payment.ID(id)).Only(ctx)
	if err != nil {
		return nil
	}
	_, err = p.Update().SetStatus(payment.StatusPaid).Save(ctx)
	return err
}
