package data

import (
	"context"
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	pb "mall/api/mall/service/v1"
	"mall/ent"
	"mall/ent/item"
	"mall/ent/order"
)

type OrderRepo struct {
	data *Data
}

func NewOrderRepo(data *Data) *OrderRepo {
	return &OrderRepo{data: data}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, userID int, address string, items []*pb.OrderItem) (*ent.Order, error) {
	o, err := r.data.db.Order.Create().SetUserID(userID).SetAddress(address).Save(ctx)
	if err != nil {
		return nil, err
	}
	for _, i := range items {
		product, err := r.data.db.Item.Query().Where(item.IDEQ(int(i.ProductId))).Only(ctx)
		if err != nil {
			return nil, err
		}
		err = r.data.db.OrderItem.Create().SetOrder(o).SetItem(product).SetQuantity(int(i.Quantity)).SetPrice(product.Price).Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	o, err = r.data.db.Order.Query().Where(order.IDEQ(o.ID)).WithItems().WithUser().Only(ctx)
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	err = r.data.mq.Publish(
		"",
		"order_queue",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	return o, nil
}

func (r *OrderRepo) GetOrderList(ctx context.Context, size, page int) ([]*ent.Order, error) {
	limit := size
	offset := (page - 1) * size
	return r.data.db.Order.Query().WithItems().WithUser().Limit(limit).Offset(offset).All(ctx)
}

func (r *OrderRepo) GetOrder(ctx context.Context, id int) (*ent.Order, error) {
	return r.data.db.Order.Query().Where(order.IDEQ(id)).WithItems().Only(ctx)
}

func (r *OrderRepo) UpdateOrderStatus(ctx context.Context, id int, status order.Status) error {
	_, err := r.data.db.Order.UpdateOneID(id).SetStatus(status).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
