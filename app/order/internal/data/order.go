package data

import (
	"context"
	"errors"
	pb "mall/api/mall/service/v1"
	"mall/ent"
	"mall/ent/item"
	"mall/ent/order"
	"mall/ent/payment"
)

type OrderRepo struct {
	data *Data
}

func NewOrderRepo(data *Data) OrderRepo {
	return OrderRepo{data: data}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, userID int, address string, items []*pb.OrderItem) (*ent.Order, error) {
	o, err := r.data.db.Order.Create().SetUserID(userID).SetAddress(address).Save(ctx)
	if err != nil {
		return nil, err
	}
	oi := make([]*ent.OrderItemCreate, 0)
	for _, i := range items {
		product, err := r.data.db.Item.Query().Where(item.IDEQ(int(i.ProductId))).Only(ctx)
		if err != nil {
			return nil, err
		}
		create := r.data.db.OrderItem.Create().SetOrder(o).SetItem(product).SetQuantity(int(i.Quantity)).SetPrice(product.Price)
		oi = append(oi, create)
	}
	err = r.data.db.OrderItem.CreateBulk(oi...).Exec(ctx)
	if err != nil {
		return nil, err
	}
	o, err = r.data.db.Order.Query().Where(order.IDEQ(o.ID)).WithItems().Only(ctx)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (r *OrderRepo) GetOrderList(ctx context.Context, size, page int) ([]*ent.Order, error) {
	limit := size
	offset := (page - 1) * size
	return r.data.db.Order.Query().WithItems().Limit(limit).Offset(offset).All(ctx)
}

func (r *OrderRepo) GetOrder(ctx context.Context, id int) (*ent.Order, error) {
	return r.data.db.Order.Query().Where(order.IDEQ(id)).WithItems().Only(ctx)
}

func (r *OrderRepo) CreatePayment(ctx context.Context, id int) (*ent.Payment, error) {
	o, err := r.data.db.Order.Query().Where(order.IDEQ(id)).WithItems().Only(ctx)
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

func (r *OrderRepo) PayOrder(ctx context.Context, id int) (*ent.Payment, error) {
	p, err := r.data.db.Payment.Query().Where(payment.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	if p.Status != "pending" {
		return nil, errors.New("payment status error")
	}
	// TODO: call payment service
	p, err = r.data.db.Payment.UpdateOneID(p.ID).SetStatus("paid").Save(ctx)
	if err != nil {
		return nil, err
	}
	return p, nil
}
