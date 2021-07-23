package app

import (
	"context"
	"fmt"

	"github.com/uhey22e/sqlboiler-restapi-example/boiler"
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

func ListEvents(ctx context.Context) ([]*Event, error) {
	es, err := boiler.Events(
		qm.OrderBy(boiler.EventColumns.Date+" desc"),
		qm.Load(fmt.Sprintf("%s.%s", boiler.EventRels.EventUsers, boiler.EventUserRels.User)),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}
	res := make([]*Event, len(es))
	for i, e := range es {
		ps := make([]*boiler.User, len(e.R.EventUsers))
		for j, u := range e.R.EventUsers {
			ps[j] = u.R.User
		}
		res[i] = &Event{
			Event:            e,
			ParticipantUsers: ps,
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
			order by participants desc
			limit 10
		) r1 on event.id = r1.id
		order by participants desc
	`).Bind(ctx, db, &r)
	return r, nil
}
