//go:generate refiller -o refiller -s boiler.User -d restapi.User
package app

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/uhey22e/sqlboiler-restapi-example/boiler"
	"github.com/uhey22e/sqlboiler-restapi-example/restapi"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func ListUsers(ctx context.Context) ([]*restapi.User, error) {
	res := make([]*restapi.User, 0, 10)
	err := boiler.Users(qm.OrderBy(fmt.Sprintf("%s desc", boiler.UserColumns.CreatedAt))).Bind(ctx, db, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
