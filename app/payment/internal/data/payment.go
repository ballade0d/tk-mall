package data

import (
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"mall/ent"
	"mall/ent/order"
	"mall/ent/payment"
)

type PaymentRepo struct {
	data *Data
}

func NewPaymentRepo(data *Data) *PaymentRepo {
	return &PaymentRepo{data: data}
}

func (r *PaymentRepo) CreatePayment(ctx context.Context, orderID int) (*ent.Payment, error) {
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

func (r *PaymentRepo) PaySuccessful(ctx context.Context, paymentID int) error {
	p, err := r.data.db.Payment.Query().Where(payment.IDEQ(paymentID)).WithOrder().Only(ctx)
	if err != nil {
		return err
	}
	if p.Status != "pending" {
		return errors.New("payment status error")
	}
	_, err = p.Update().SetStatus(payment.StatusPaid).Save(ctx)
	if err != nil {
		return err
	}

	err = r.data.mq.Publish(
		"",
		"pay_result_queue",
		false,
		false,
		amqp.Publishing{ContentType: "application/json", Body: []byte(p.String())},
	)
	return nil
}

func (r *PaymentRepo) PayFailed(ctx context.Context, paymentID int) error {
	p, err := r.data.db.Payment.Query().Where(payment.IDEQ(paymentID)).WithOrder().Only(ctx)
	if err != nil {
		return err
	}
	if p.Status != "pending" {
		return errors.New("payment status error")
	}
	_, err = p.Update().SetStatus(payment.StatusFailed).Save(ctx)
	if err != nil {
		return err
	}

	err = r.data.mq.Publish(
		"",
		"pay_result_queue",
		false,
		false,
		amqp.Publishing{ContentType: "application/json", Body: []byte(p.String())},
	)
	return nil
}
