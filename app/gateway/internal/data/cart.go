package data

import (
	"context"
	"mall/ent"
)

type CartRepo struct {
	data *Data
}

func NewCartRepo(data *Data) *CartRepo {
	return &CartRepo{data: data}
}

func (r *CartRepo) CreateCart(ctx context.Context, userID int) (*ent.Cart, error) {
	c, err := r.data.db.Cart.Create().
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}
