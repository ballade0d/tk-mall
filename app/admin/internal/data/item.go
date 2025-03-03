package data

import (
	"context"
	"errors"
	"mall/ent"
	"mall/pkg/util"
	"time"
)

type ItemRepo struct {
	data *Data
}

type ESItem struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

func NewItemRepo(data *Data) *ItemRepo {
	return &ItemRepo{data: data}
}

func (r *ItemRepo) indexItem(ctx context.Context, it *ent.Item) error {
	_, err := r.data.es.Index(r.data.conf.ElasticSearch.Indices).Id(string(rune(it.ID))).Request(it).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *ItemRepo) CreateItem(ctx context.Context, name string, description string, price float32) (*ent.Item, error) {
	it, err := r.data.db.Item.Create().SetName(name).SetDescription(description).SetPrice(price).Save(ctx)
	if err != nil {
		return nil, err
	}
	err = r.indexItem(ctx, it)
	if err != nil {
		return nil, err
	}
	return it, nil
}

func (r *ItemRepo) deleteItem(ctx context.Context, id int) error {
	_, err := r.data.es.Delete(r.data.conf.ElasticSearch.Indices, string(rune(id))).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *ItemRepo) DeleteItem(ctx context.Context, id int) error {
	err := r.data.db.Item.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	err = r.deleteItem(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ItemRepo) EditItem(ctx context.Context, id int, name string, description string, price float32) (*ent.Item, error) {
	it, err := r.data.db.Item.UpdateOneID(id).SetName(name).SetDescription(description).SetPrice(price).Save(ctx)
	if err != nil {
		return nil, err
	}
	err = r.indexItem(ctx, it)
	if err != nil {
		return nil, err
	}
	return it, nil
}

func (r *ItemRepo) ListItems(ctx context.Context) ([]*ent.Item, error) {
	its, err := r.data.db.Item.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return its, nil
}

func (r *ItemRepo) AddStock(ctx context.Context, id int, stock int) (*ent.Item, error) {
	lock := util.NewRedisLock(r.data.rdb, id, 5*time.Second)
	if !lock.LockWithQueue(ctx, 10*time.Second) {
		return nil, errors.New("lock failed")
	}
	defer lock.Unlock(ctx)
	it, err := r.data.db.Item.UpdateOneID(id).AddStock(stock).Save(ctx)
	if err != nil {
		return nil, err
	}
	return it, nil
}
