package data

import (
	"context"
	"mall/ent"
	"mall/ent/cart"
	"mall/ent/cartitem"
	"mall/ent/item"
	"mall/ent/user"
	"mall/pkg/util"
)

type CartRepo struct {
	data *Data
}

func NewCartRepo(data *Data) *CartRepo {
	return &CartRepo{data: data}
}

func (r *CartRepo) GetCart(ctx context.Context) (*ent.Cart, error) {
	id := ctx.Value("claims").(util.Claims).UserId
	c, err := r.data.db.Cart.Query().Where(cart.HasUserWith(user.ID(id))).Only(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CartRepo) AddToCart(ctx context.Context, itemID int, quantity int) (*ent.CartItem, error) {
	c, err := r.GetCart(ctx)
	if err != nil {
		return nil, err
	}
	old, err := r.data.db.CartItem.Query().Where(cartitem.HasCartWith(cart.ID(c.ID)), cartitem.HasItemWith(item.ID(itemID))).Only(ctx)

	if ent.IsNotFound(err) {
		it, err := r.data.db.Item.Query().Where(item.ID(itemID)).Only(ctx)
		if err != nil {
			return nil, err
		}
		ci, err := r.data.db.CartItem.Create().SetCart(c).SetItem(it).SetQuantity(quantity).Save(ctx)
		if err != nil {
			return nil, err
		}
		return ci, nil
	} else {
		old.Quantity += quantity
		_, err = old.Update().SetQuantity(old.Quantity).Save(ctx)
		if err != nil {
			return nil, err
		}
		return old, nil
	}
}

func (r *CartRepo) RemoveFromCart(ctx context.Context, itemID int) error {
	c, err := r.GetCart(ctx)
	if err != nil {
		return err
	}
	_, err = r.data.db.CartItem.Delete().Where(cartitem.HasCartWith(cart.ID(c.ID)), cartitem.HasItemWith(item.ID(itemID))).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *CartRepo) ClearCart(ctx context.Context) error {
	c, err := r.GetCart(ctx)
	if err != nil {
		return err
	}
	_, err = r.data.db.CartItem.Delete().Where(cartitem.HasCartWith(cart.ID(c.ID))).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
