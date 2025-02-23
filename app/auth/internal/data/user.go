package data

import (
	"context"
	"mall/ent"
	"mall/ent/user"
)

type UserRepo struct {
	data *Data
}

func NewUserRepo(data *Data) UserRepo {
	return UserRepo{data: data}
}

func (r *UserRepo) FindUserByID(ctx context.Context, id int) (*ent.User, error) {
	u, err := r.data.db.User.Query().Where(user.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepo) FindUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	usr, err := r.data.db.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (r *UserRepo) CreateUser(ctx context.Context, u *ent.User, p *ent.Password) (*ent.User, error) {
	u, err := r.data.db.User.Create().SetName(u.Name).SetEmail(u.Email).AddPassword(p).Save(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}
