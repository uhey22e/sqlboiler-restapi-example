package app

import (
	"context"
	"fmt"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/uhey22e/sqlboiler-restapi-example/boiler"
	"github.com/uhey22e/sqlboiler-restapi-example/restapi"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Event struct {
	*boiler.Event
	// イベント参加者
	ParticipantUsers []*boiler.User `json:"participantUsers"`
}

type EventPopularity struct {
	boiler.Event `boil:",bind"`
	Participants int `boil:"participants" json:"participants"`
}

func ListEvents(ctx context.Context) ([]*restapi.Event, error) {
	rows, err := boiler.Events(
		qm.OrderBy(boiler.EventColumns.Date+" desc"),
		qm.Load(fmt.Sprintf("%s.%s", boiler.EventRels.EventUsers, boiler.EventUserRels.User)),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	res := make([]*restapi.Event, len(rows))
	for i, e := range rows {
		participants := make([]*restapi.User, len(e.R.EventUsers))
		for j := range e.R.EventUsers {
			u := e.R.EventUsers[j].R.User
			participants[j] = &restapi.User{
				Id:   u.ID,
				Name: u.Name,
			}
		}
		res[i] = &restapi.Event{
			Id:           e.ID,
			Name:         e.Name,
			Description:  e.Description,
			Date:         types.Date{Time: e.Date},
			Participants: participants,
		}
	}
	return res, nil
}

func ListPopularEvents(ctx context.Context) ([]*EventPopularity, error) {
	r := []*EventPopularity{}
	queries.Raw(`
		select event.*, coalesce(r1.participants, 0) as participants
		from event
		left join (
			select event_id as id, count(*) as participants
			from event_user
			group by event_id
		) r1 on event.id = r1.id
		order by participants desc
	`).Bind(ctx, db, &r)
	return r, nil
}
