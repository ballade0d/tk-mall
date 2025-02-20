package data

import (
	"context"
	"mall/ent"
	"mall/ent/item"
)

type ItemRepo struct {
	data *Data
}

func NewItemRepo(data *Data) ItemRepo {
	return ItemRepo{data: data}
}

func (r *ItemRepo) indexItem(ctx context.Context, it *ent.Item) error {
	_, err := r.data.es.Index(r.data.conf.ElasticSearch.Indices).Id(string(it.ID)).Request(it).Do(ctx)
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

func (r *ItemRepo) deleteItem(ctx context.Context, id int32) error {
	_, err := r.data.es.Delete(r.data.conf.ElasticSearch.Indices, string(id)).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *ItemRepo) DeleteItem(ctx context.Context, id int32) error {
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

func (r *ItemRepo) FindItemByID(ctx context.Context, id int32) (*ent.Item, error) {
	it, err := r.data.db.Item.Query().Where(item.ID(id)).Only(ctx)
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
