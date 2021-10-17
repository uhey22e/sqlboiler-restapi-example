//go:generate refiller -o refiller -s boiler.Event -d restapi.Event
package app

import (
	"context"
	"fmt"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/uhey22e/sqlboiler-restapi-example/boiler"
	"github.com/uhey22e/sqlboiler-restapi-example/refiller"
	"github.com/uhey22e/sqlboiler-restapi-example/restapi"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

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
		res[i] = refiller.NewRestapiEvent(e)
		res[i].Date = types.Date{Time: e.Date}
		participants := make([]*restapi.User, len(e.R.EventUsers))
		for j := range e.R.EventUsers {
			participants[j] = refiller.NewRestapiUser(e.R.EventUsers[j].R.User)
		}
	}
	return res, nil
}

func ListPopularEvents(ctx context.Context) ([]*restapi.PopularEvent, error) {
	// Re-define "Date" column because openapi_types.Date dosen't have "boil" tag.
	type popularEvent struct {
		restapi.PopularEvent `boil:",bind"`
		Date                 time.Time `boil:"date"`
	}

	rows := make([]*popularEvent, 0, 10)
	queries.Raw(`
		select
			event.*
			, coalesce(r1.participants, 0) as "participants"
			, coalesce(r1.registrations_in_the_past_week) as "registrations_in_the_past_week"
		from "event"
		left join (
			select
				event_id as "id"
				, count(*) as "participants"
				, sum( case when (now() - registered_at < '7 days') then 1 else 0 end ) as "registrations_in_the_past_week"
			from event_user
			group by event_id
		) r1 on event.id = r1.id
		order by registrations_in_the_past_week desc, participants desc
	`).Bind(ctx, db, &rows)

	res := make([]*restapi.PopularEvent, len(rows))
	for i := range rows {
		res[i] = &rows[i].PopularEvent
		res[i].Date = types.Date{Time: rows[i].Date}
	}
	return res, nil
}
