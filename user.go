package app

import (
	"context"
	"encoding/json"

	"github.com/uhey22e/sqlboiler-restapi-example/boiler"
	"github.com/uhey22e/sqlboiler-restapi-example/restapi"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CreateUser(ctx context.Context, b []byte) (*boiler.User, error) {
	u := &boiler.User{}
	if err := json.Unmarshal(b, u); err != nil {
		return nil, err
	}
	err := u.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return u, err
}

func ListUsers(ctx context.Context) ([]*boiler.User, error) {
	users, err := boiler.Users().All(ctx, db)
	if err != nil {
		return nil, err
	}
	if users == nil {
		users = []*boiler.User{}
	}
	return users, nil
}

func ListUsersAPI(ctx context.Context) ([]*restapi.User, error) {
	res := []*restapi.User{}
	err := boiler.Users().Bind(ctx, db, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
