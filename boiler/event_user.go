// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boiler

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// EventUser is an object representing the database table.
type EventUser struct {
	EventID string `boil:"event_id" json:"eventID" toml:"eventID" yaml:"eventID"`
	UserID  string `boil:"user_id" json:"userID" toml:"userID" yaml:"userID"`
	// A user registered at
	RegisteredAt time.Time `boil:"registered_at" json:"registeredAt" toml:"registeredAt" yaml:"registeredAt"`
	CreatedAt    time.Time `boil:"created_at" json:"-" toml:"-" yaml:"-"`

	R *eventUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L eventUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EventUserColumns = struct {
	EventID      string
	UserID       string
	RegisteredAt string
	CreatedAt    string
}{
	EventID:      "event_id",
	UserID:       "user_id",
	RegisteredAt: "registered_at",
	CreatedAt:    "created_at",
}

var EventUserTableColumns = struct {
	EventID      string
	UserID       string
	RegisteredAt string
	CreatedAt    string
}{
	EventID:      "event_user.event_id",
	UserID:       "event_user.user_id",
	RegisteredAt: "event_user.registered_at",
	CreatedAt:    "event_user.created_at",
}

// Generated where

var EventUserWhere = struct {
	EventID      whereHelperstring
	UserID       whereHelperstring
	RegisteredAt whereHelpertime_Time
	CreatedAt    whereHelpertime_Time
}{
	EventID:      whereHelperstring{field: "\"event_user\".\"event_id\""},
	UserID:       whereHelperstring{field: "\"event_user\".\"user_id\""},
	RegisteredAt: whereHelpertime_Time{field: "\"event_user\".\"registered_at\""},
	CreatedAt:    whereHelpertime_Time{field: "\"event_user\".\"created_at\""},
}

// EventUserRels is where relationship names are stored.
var EventUserRels = struct {
	Event string
	User  string
}{
	Event: "Event",
	User:  "User",
}

// eventUserR is where relationships are stored.
type eventUserR struct {
	Event *Event `boil:"Event" json:"Event" toml:"Event" yaml:"Event"`
	User  *User  `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*eventUserR) NewStruct() *eventUserR {
	return &eventUserR{}
}

// eventUserL is where Load methods for each relationship are stored.
type eventUserL struct{}

var (
	eventUserAllColumns            = []string{"event_id", "user_id", "registered_at", "created_at"}
	eventUserColumnsWithoutDefault = []string{"event_id", "user_id", "registered_at"}
	eventUserColumnsWithDefault    = []string{"created_at"}
	eventUserPrimaryKeyColumns     = []string{"event_id", "user_id"}
)

type (
	// EventUserSlice is an alias for a slice of pointers to EventUser.
	// This should almost always be used instead of []EventUser.
	EventUserSlice []*EventUser
	// EventUserHook is the signature for custom EventUser hook methods
	EventUserHook func(context.Context, boil.ContextExecutor, *EventUser) error

	eventUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	eventUserType                 = reflect.TypeOf(&EventUser{})
	eventUserMapping              = queries.MakeStructMapping(eventUserType)
	eventUserPrimaryKeyMapping, _ = queries.BindMapping(eventUserType, eventUserMapping, eventUserPrimaryKeyColumns)
	eventUserInsertCacheMut       sync.RWMutex
	eventUserInsertCache          = make(map[string]insertCache)
	eventUserUpdateCacheMut       sync.RWMutex
	eventUserUpdateCache          = make(map[string]updateCache)
	eventUserUpsertCacheMut       sync.RWMutex
	eventUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var eventUserBeforeInsertHooks []EventUserHook
var eventUserBeforeUpdateHooks []EventUserHook
var eventUserBeforeDeleteHooks []EventUserHook
var eventUserBeforeUpsertHooks []EventUserHook

var eventUserAfterInsertHooks []EventUserHook
var eventUserAfterSelectHooks []EventUserHook
var eventUserAfterUpdateHooks []EventUserHook
var eventUserAfterDeleteHooks []EventUserHook
var eventUserAfterUpsertHooks []EventUserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EventUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EventUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EventUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EventUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EventUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EventUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EventUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EventUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EventUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEventUserHook registers your hook function for all future operations.
func AddEventUserHook(hookPoint boil.HookPoint, eventUserHook EventUserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		eventUserBeforeInsertHooks = append(eventUserBeforeInsertHooks, eventUserHook)
	case boil.BeforeUpdateHook:
		eventUserBeforeUpdateHooks = append(eventUserBeforeUpdateHooks, eventUserHook)
	case boil.BeforeDeleteHook:
		eventUserBeforeDeleteHooks = append(eventUserBeforeDeleteHooks, eventUserHook)
	case boil.BeforeUpsertHook:
		eventUserBeforeUpsertHooks = append(eventUserBeforeUpsertHooks, eventUserHook)
	case boil.AfterInsertHook:
		eventUserAfterInsertHooks = append(eventUserAfterInsertHooks, eventUserHook)
	case boil.AfterSelectHook:
		eventUserAfterSelectHooks = append(eventUserAfterSelectHooks, eventUserHook)
	case boil.AfterUpdateHook:
		eventUserAfterUpdateHooks = append(eventUserAfterUpdateHooks, eventUserHook)
	case boil.AfterDeleteHook:
		eventUserAfterDeleteHooks = append(eventUserAfterDeleteHooks, eventUserHook)
	case boil.AfterUpsertHook:
		eventUserAfterUpsertHooks = append(eventUserAfterUpsertHooks, eventUserHook)
	}
}

// One returns a single eventUser record from the query.
func (q eventUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EventUser, error) {
	o := &EventUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for event_user")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all EventUser records from the query.
func (q eventUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (EventUserSlice, error) {
	var o []*EventUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to EventUser slice")
	}

	if len(eventUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all EventUser records in the query.
func (q eventUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count event_user rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q eventUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if event_user exists")
	}

	return count > 0, nil
}

// Event pointed to by the foreign key.
func (o *EventUser) Event(mods ...qm.QueryMod) eventQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.EventID),
	}

	queryMods = append(queryMods, mods...)

	query := Events(queryMods...)
	queries.SetFrom(query.Query, "\"event\"")

	return query
}

// User pointed to by the foreign key.
func (o *EventUser) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"user\"")

	return query
}

// LoadEvent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (eventUserL) LoadEvent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEventUser interface{}, mods queries.Applicator) error {
	var slice []*EventUser
	var object *EventUser

	if singular {
		object = maybeEventUser.(*EventUser)
	} else {
		slice = *maybeEventUser.(*[]*EventUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &eventUserR{}
		}
		args = append(args, object.EventID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &eventUserR{}
			}

			for _, a := range args {
				if a == obj.EventID {
					continue Outer
				}
			}

			args = append(args, obj.EventID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`event`),
		qm.WhereIn(`event.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Event")
	}

	var resultSlice []*Event
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Event")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for event")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for event")
	}

	if len(eventUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Event = foreign
		if foreign.R == nil {
			foreign.R = &eventR{}
		}
		foreign.R.EventUsers = append(foreign.R.EventUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.EventID == foreign.ID {
				local.R.Event = foreign
				if foreign.R == nil {
					foreign.R = &eventR{}
				}
				foreign.R.EventUsers = append(foreign.R.EventUsers, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (eventUserL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEventUser interface{}, mods queries.Applicator) error {
	var slice []*EventUser
	var object *EventUser

	if singular {
		object = maybeEventUser.(*EventUser)
	} else {
		slice = *maybeEventUser.(*[]*EventUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &eventUserR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &eventUserR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user`),
		qm.WhereIn(`user.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user")
	}

	if len(eventUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.EventUsers = append(foreign.R.EventUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.EventUsers = append(foreign.R.EventUsers, local)
				break
			}
		}
	}

	return nil
}

// SetEvent of the eventUser to the related item.
// Sets o.R.Event to related.
// Adds o to related.R.EventUsers.
func (o *EventUser) SetEvent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Event) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"event_user\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"event_id"}),
		strmangle.WhereClause("\"", "\"", 2, eventUserPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.EventID, o.UserID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.EventID = related.ID
	if o.R == nil {
		o.R = &eventUserR{
			Event: related,
		}
	} else {
		o.R.Event = related
	}

	if related.R == nil {
		related.R = &eventR{
			EventUsers: EventUserSlice{o},
		}
	} else {
		related.R.EventUsers = append(related.R.EventUsers, o)
	}

	return nil
}

// SetUser of the eventUser to the related item.
// Sets o.R.User to related.
// Adds o to related.R.EventUsers.
func (o *EventUser) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"event_user\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, eventUserPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.EventID, o.UserID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &eventUserR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			EventUsers: EventUserSlice{o},
		}
	} else {
		related.R.EventUsers = append(related.R.EventUsers, o)
	}

	return nil
}

// EventUsers retrieves all the records using an executor.
func EventUsers(mods ...qm.QueryMod) eventUserQuery {
	mods = append(mods, qm.From("\"event_user\""))
	return eventUserQuery{NewQuery(mods...)}
}

// FindEventUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEventUser(ctx context.Context, exec boil.ContextExecutor, eventID string, userID string, selectCols ...string) (*EventUser, error) {
	eventUserObj := &EventUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"event_user\" where \"event_id\"=$1 AND \"user_id\"=$2", sel,
	)

	q := queries.Raw(query, eventID, userID)

	err := q.Bind(ctx, exec, eventUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from event_user")
	}

	if err = eventUserObj.doAfterSelectHooks(ctx, exec); err != nil {
		return eventUserObj, err
	}

	return eventUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EventUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no event_user provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	eventUserInsertCacheMut.RLock()
	cache, cached := eventUserInsertCache[key]
	eventUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			eventUserAllColumns,
			eventUserColumnsWithDefault,
			eventUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(eventUserType, eventUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(eventUserType, eventUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"event_user\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"event_user\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "boiler: unable to insert into event_user")
	}

	if !cached {
		eventUserInsertCacheMut.Lock()
		eventUserInsertCache[key] = cache
		eventUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the EventUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EventUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	eventUserUpdateCacheMut.RLock()
	cache, cached := eventUserUpdateCache[key]
	eventUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			eventUserAllColumns,
			eventUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update event_user, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"event_user\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, eventUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(eventUserType, eventUserMapping, append(wl, eventUserPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update event_user row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for event_user")
	}

	if !cached {
		eventUserUpdateCacheMut.Lock()
		eventUserUpdateCache[key] = cache
		eventUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q eventUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for event_user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for event_user")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EventUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("boiler: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"event_user\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, eventUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in eventUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all eventUser")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EventUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no event_user provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventUserColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	eventUserUpsertCacheMut.RLock()
	cache, cached := eventUserUpsertCache[key]
	eventUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			eventUserAllColumns,
			eventUserColumnsWithDefault,
			eventUserColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			eventUserAllColumns,
			eventUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("boiler: unable to upsert event_user, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(eventUserPrimaryKeyColumns))
			copy(conflict, eventUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"event_user\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(eventUserType, eventUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(eventUserType, eventUserMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "boiler: unable to upsert event_user")
	}

	if !cached {
		eventUserUpsertCacheMut.Lock()
		eventUserUpsertCache[key] = cache
		eventUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single EventUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EventUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no EventUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), eventUserPrimaryKeyMapping)
	sql := "DELETE FROM \"event_user\" WHERE \"event_id\"=$1 AND \"user_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from event_user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for event_user")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q eventUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no eventUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from event_user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for event_user")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EventUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(eventUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"event_user\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, eventUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from eventUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for event_user")
	}

	if len(eventUserAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *EventUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEventUser(ctx, exec, o.EventID, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EventUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EventUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"event_user\".* FROM \"event_user\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, eventUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in EventUserSlice")
	}

	*o = slice

	return nil
}

// EventUserExists checks if the EventUser row exists.
func EventUserExists(ctx context.Context, exec boil.ContextExecutor, eventID string, userID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"event_user\" where \"event_id\"=$1 AND \"user_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, eventID, userID)
	}
	row := exec.QueryRowContext(ctx, sql, eventID, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if event_user exists")
	}

	return exists, nil
}
