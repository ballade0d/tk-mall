package data

import (
	"context"
	"errors"
	pb "mall/api/mall/service/v1"
	"mall/ent"
	"mall/ent/item"
	"mall/pkg/util"
	"time"
)

type ItemRepo struct {
	data *Data
}

func NewItemRepo(data *Data) *ItemRepo {
	return &ItemRepo{data: data}
}

func (r *ItemRepo) CheckAndReduceStock(ctx context.Context, orderItem []*pb.OrderItem) error {
	lock := make(map[int]*util.RedisLock)
	items := make([]*ent.Item, 0)
	for _, i := range orderItem {
		l := util.NewRedisLock(r.data.rdb, int(i.ProductId), time.Second*5)
		l.LockWithQueue(ctx, time.Second*5)
		lock[int(i.ProductId)] = l
		it, err := r.data.db.Item.Query().Where(item.IDEQ(int(i.ProductId))).Only(ctx)
		if err != nil {
			return err
		}
		if it.Stock < int(i.Quantity) {
			return errors.New("stock not enough")
		}
		items = append(items, it)
	}

	defer func() {
		for _, l := range lock {
			l.Unlock(ctx)
		}
	}()

	for index, i := range items {
		err := i.Update().SetStock(i.Stock - int(orderItem[index].Quantity)).Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
