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

func NewItemRepo(data *Data) *ItemRepo {
	return &ItemRepo{data: data}
}

func (r *ItemRepo) FindItemByID(ctx context.Context, id int) (*ent.Item, error) {
	it, err := r.data.db.Item.Query().Where(item.ID(id)).Only(ctx)
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
