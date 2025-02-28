package data

import (
	"context"
	"mall/ent"
)

type PasswordRepo struct {
	data *Data
}

func NewPasswordRepo(data *Data) PasswordRepo {
	return PasswordRepo{data: data}
}

func (r *PasswordRepo) CreatePassword(ctx context.Context, password string) (*ent.Password, error) {
	pwd, err := r.data.db.Password.Create().SetPassword(password).Save(ctx)
	if err != nil {
		return nil, err
	}
	return pwd, nil
}
