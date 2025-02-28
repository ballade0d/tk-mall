package data

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"mall/ent"
	"mall/ent/item"
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

func NewItemRepo(data *Data) ItemRepo {
	return ItemRepo{data: data}
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

func (r *ItemRepo) FindItemByID(ctx context.Context, id int) (*ent.Item, error) {
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

func (r *ItemRepo) AddStock(ctx context.Context, id int, stock int) (*ent.Item, error) {
	it, err := r.data.db.Item.UpdateOneID(id).AddStock(stock).Save(ctx)
	if err != nil {
		return nil, err
	}
	return it, nil
}

func (r *ItemRepo) SearchItems(ctx context.Context, query string) ([]*ESItem, error) {
	response, err := r.data.es.Search().Index(r.data.conf.ElasticSearch.Indices).Request(&search.Request{
		Query: &types.Query{
			Match: map[string]types.MatchQuery{
				"name": {Query: query},
			},
		},
	}).Do(ctx)
	if err != nil {
		return nil, err
	}
	total := response.Hits.Total.Value
	its := make([]*ESItem, 0, total)
	for _, hit := range response.Hits.Hits {
		var source ESItem
		err := json.Unmarshal(hit.Source_, &source)
		if err != nil {
			return nil, err
		}
		its = append(its, &source)
	}
	return its, nil
}
