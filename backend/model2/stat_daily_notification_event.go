// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model2

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

// StatDailyNotificationEvent is an object representing the database table.
type StatDailyNotificationEvent struct {
	Date    time.Time `boil:"date" json:"date" toml:"date" yaml:"date"`
	Event   string    `boil:"event" json:"event" toml:"event" yaml:"event"`
	Count   uint      `boil:"count" json:"count" toml:"count" yaml:"count"`
	UuCount uint      `boil:"uu_count" json:"uu_count" toml:"uu_count" yaml:"uu_count"`

	R *statDailyNotificationEventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L statDailyNotificationEventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StatDailyNotificationEventColumns = struct {
	Date    string
	Event   string
	Count   string
	UuCount string
}{
	Date:    "date",
	Event:   "event",
	Count:   "count",
	UuCount: "uu_count",
}

var StatDailyNotificationEventTableColumns = struct {
	Date    string
	Event   string
	Count   string
	UuCount string
}{
	Date:    "stat_daily_notification_event.date",
	Event:   "stat_daily_notification_event.event",
	Count:   "stat_daily_notification_event.count",
	UuCount: "stat_daily_notification_event.uu_count",
}

// Generated where

var StatDailyNotificationEventWhere = struct {
	Date    whereHelpertime_Time
	Event   whereHelperstring
	Count   whereHelperuint
	UuCount whereHelperuint
}{
	Date:    whereHelpertime_Time{field: "`stat_daily_notification_event`.`date`"},
	Event:   whereHelperstring{field: "`stat_daily_notification_event`.`event`"},
	Count:   whereHelperuint{field: "`stat_daily_notification_event`.`count`"},
	UuCount: whereHelperuint{field: "`stat_daily_notification_event`.`uu_count`"},
}

// StatDailyNotificationEventRels is where relationship names are stored.
var StatDailyNotificationEventRels = struct {
}{}

// statDailyNotificationEventR is where relationships are stored.
type statDailyNotificationEventR struct {
}

// NewStruct creates a new relationship struct
func (*statDailyNotificationEventR) NewStruct() *statDailyNotificationEventR {
	return &statDailyNotificationEventR{}
}

// statDailyNotificationEventL is where Load methods for each relationship are stored.
type statDailyNotificationEventL struct{}

var (
	statDailyNotificationEventAllColumns            = []string{"date", "event", "count", "uu_count"}
	statDailyNotificationEventColumnsWithoutDefault = []string{"date", "event", "count", "uu_count"}
	statDailyNotificationEventColumnsWithDefault    = []string{}
	statDailyNotificationEventPrimaryKeyColumns     = []string{"date", "event"}
)

type (
	// StatDailyNotificationEventSlice is an alias for a slice of pointers to StatDailyNotificationEvent.
	// This should almost always be used instead of []StatDailyNotificationEvent.
	StatDailyNotificationEventSlice []*StatDailyNotificationEvent
	// StatDailyNotificationEventHook is the signature for custom StatDailyNotificationEvent hook methods
	StatDailyNotificationEventHook func(context.Context, boil.ContextExecutor, *StatDailyNotificationEvent) error

	statDailyNotificationEventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	statDailyNotificationEventType                 = reflect.TypeOf(&StatDailyNotificationEvent{})
	statDailyNotificationEventMapping              = queries.MakeStructMapping(statDailyNotificationEventType)
	statDailyNotificationEventPrimaryKeyMapping, _ = queries.BindMapping(statDailyNotificationEventType, statDailyNotificationEventMapping, statDailyNotificationEventPrimaryKeyColumns)
	statDailyNotificationEventInsertCacheMut       sync.RWMutex
	statDailyNotificationEventInsertCache          = make(map[string]insertCache)
	statDailyNotificationEventUpdateCacheMut       sync.RWMutex
	statDailyNotificationEventUpdateCache          = make(map[string]updateCache)
	statDailyNotificationEventUpsertCacheMut       sync.RWMutex
	statDailyNotificationEventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var statDailyNotificationEventBeforeInsertHooks []StatDailyNotificationEventHook
var statDailyNotificationEventBeforeUpdateHooks []StatDailyNotificationEventHook
var statDailyNotificationEventBeforeDeleteHooks []StatDailyNotificationEventHook
var statDailyNotificationEventBeforeUpsertHooks []StatDailyNotificationEventHook

var statDailyNotificationEventAfterInsertHooks []StatDailyNotificationEventHook
var statDailyNotificationEventAfterSelectHooks []StatDailyNotificationEventHook
var statDailyNotificationEventAfterUpdateHooks []StatDailyNotificationEventHook
var statDailyNotificationEventAfterDeleteHooks []StatDailyNotificationEventHook
var statDailyNotificationEventAfterUpsertHooks []StatDailyNotificationEventHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StatDailyNotificationEvent) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range statDailyNotificationEventBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StatDailyNotificationEvent) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range statDailyNotificationEventBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StatDailyNotificationEvent) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range statDailyNotificationEventBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StatDailyNotificationEvent) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range statDailyNotificationEventBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StatDailyNotificationEvent) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range statDailyNotificationEventAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StatDailyNotificationEvent) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range statDailyNotificationEventAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StatDailyNotificationEvent) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range statDailyNotificationEventAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StatDailyNotificationEvent) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range statDailyNotificationEventAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StatDailyNotificationEvent) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range statDailyNotificationEventAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStatDailyNotificationEventHook registers your hook function for all future operations.
func AddStatDailyNotificationEventHook(hookPoint boil.HookPoint, statDailyNotificationEventHook StatDailyNotificationEventHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		statDailyNotificationEventBeforeInsertHooks = append(statDailyNotificationEventBeforeInsertHooks, statDailyNotificationEventHook)
	case boil.BeforeUpdateHook:
		statDailyNotificationEventBeforeUpdateHooks = append(statDailyNotificationEventBeforeUpdateHooks, statDailyNotificationEventHook)
	case boil.BeforeDeleteHook:
		statDailyNotificationEventBeforeDeleteHooks = append(statDailyNotificationEventBeforeDeleteHooks, statDailyNotificationEventHook)
	case boil.BeforeUpsertHook:
		statDailyNotificationEventBeforeUpsertHooks = append(statDailyNotificationEventBeforeUpsertHooks, statDailyNotificationEventHook)
	case boil.AfterInsertHook:
		statDailyNotificationEventAfterInsertHooks = append(statDailyNotificationEventAfterInsertHooks, statDailyNotificationEventHook)
	case boil.AfterSelectHook:
		statDailyNotificationEventAfterSelectHooks = append(statDailyNotificationEventAfterSelectHooks, statDailyNotificationEventHook)
	case boil.AfterUpdateHook:
		statDailyNotificationEventAfterUpdateHooks = append(statDailyNotificationEventAfterUpdateHooks, statDailyNotificationEventHook)
	case boil.AfterDeleteHook:
		statDailyNotificationEventAfterDeleteHooks = append(statDailyNotificationEventAfterDeleteHooks, statDailyNotificationEventHook)
	case boil.AfterUpsertHook:
		statDailyNotificationEventAfterUpsertHooks = append(statDailyNotificationEventAfterUpsertHooks, statDailyNotificationEventHook)
	}
}

// One returns a single statDailyNotificationEvent record from the query.
func (q statDailyNotificationEventQuery) One(ctx context.Context, exec boil.ContextExecutor) (*StatDailyNotificationEvent, error) {
	o := &StatDailyNotificationEvent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model2: failed to execute a one query for stat_daily_notification_event")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all StatDailyNotificationEvent records from the query.
func (q statDailyNotificationEventQuery) All(ctx context.Context, exec boil.ContextExecutor) (StatDailyNotificationEventSlice, error) {
	var o []*StatDailyNotificationEvent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model2: failed to assign all query results to StatDailyNotificationEvent slice")
	}

	if len(statDailyNotificationEventAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all StatDailyNotificationEvent records in the query.
func (q statDailyNotificationEventQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model2: failed to count stat_daily_notification_event rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q statDailyNotificationEventQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model2: failed to check if stat_daily_notification_event exists")
	}

	return count > 0, nil
}

// StatDailyNotificationEvents retrieves all the records using an executor.
func StatDailyNotificationEvents(mods ...qm.QueryMod) statDailyNotificationEventQuery {
	mods = append(mods, qm.From("`stat_daily_notification_event`"))
	return statDailyNotificationEventQuery{NewQuery(mods...)}
}

// FindStatDailyNotificationEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStatDailyNotificationEvent(ctx context.Context, exec boil.ContextExecutor, date time.Time, event string, selectCols ...string) (*StatDailyNotificationEvent, error) {
	statDailyNotificationEventObj := &StatDailyNotificationEvent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `stat_daily_notification_event` where `date`=? AND `event`=?", sel,
	)

	q := queries.Raw(query, date, event)

	err := q.Bind(ctx, exec, statDailyNotificationEventObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model2: unable to select from stat_daily_notification_event")
	}

	if err = statDailyNotificationEventObj.doAfterSelectHooks(ctx, exec); err != nil {
		return statDailyNotificationEventObj, err
	}

	return statDailyNotificationEventObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *StatDailyNotificationEvent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model2: no stat_daily_notification_event provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(statDailyNotificationEventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	statDailyNotificationEventInsertCacheMut.RLock()
	cache, cached := statDailyNotificationEventInsertCache[key]
	statDailyNotificationEventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			statDailyNotificationEventAllColumns,
			statDailyNotificationEventColumnsWithDefault,
			statDailyNotificationEventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(statDailyNotificationEventType, statDailyNotificationEventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(statDailyNotificationEventType, statDailyNotificationEventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `stat_daily_notification_event` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `stat_daily_notification_event` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `stat_daily_notification_event` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, statDailyNotificationEventPrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model2: unable to insert into stat_daily_notification_event")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Date,
		o.Event,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model2: unable to populate default values for stat_daily_notification_event")
	}

CacheNoHooks:
	if !cached {
		statDailyNotificationEventInsertCacheMut.Lock()
		statDailyNotificationEventInsertCache[key] = cache
		statDailyNotificationEventInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the StatDailyNotificationEvent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *StatDailyNotificationEvent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	statDailyNotificationEventUpdateCacheMut.RLock()
	cache, cached := statDailyNotificationEventUpdateCache[key]
	statDailyNotificationEventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			statDailyNotificationEventAllColumns,
			statDailyNotificationEventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model2: unable to update stat_daily_notification_event, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `stat_daily_notification_event` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, statDailyNotificationEventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(statDailyNotificationEventType, statDailyNotificationEventMapping, append(wl, statDailyNotificationEventPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model2: unable to update stat_daily_notification_event row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model2: failed to get rows affected by update for stat_daily_notification_event")
	}

	if !cached {
		statDailyNotificationEventUpdateCacheMut.Lock()
		statDailyNotificationEventUpdateCache[key] = cache
		statDailyNotificationEventUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q statDailyNotificationEventQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to update all for stat_daily_notification_event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to retrieve rows affected for stat_daily_notification_event")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StatDailyNotificationEventSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("model2: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), statDailyNotificationEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `stat_daily_notification_event` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, statDailyNotificationEventPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to update all in statDailyNotificationEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to retrieve rows affected all in update all statDailyNotificationEvent")
	}
	return rowsAff, nil
}

var mySQLStatDailyNotificationEventUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *StatDailyNotificationEvent) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model2: no stat_daily_notification_event provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(statDailyNotificationEventColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLStatDailyNotificationEventUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	statDailyNotificationEventUpsertCacheMut.RLock()
	cache, cached := statDailyNotificationEventUpsertCache[key]
	statDailyNotificationEventUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			statDailyNotificationEventAllColumns,
			statDailyNotificationEventColumnsWithDefault,
			statDailyNotificationEventColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			statDailyNotificationEventAllColumns,
			statDailyNotificationEventPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("model2: unable to upsert stat_daily_notification_event, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`stat_daily_notification_event`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `stat_daily_notification_event` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(statDailyNotificationEventType, statDailyNotificationEventMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(statDailyNotificationEventType, statDailyNotificationEventMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model2: unable to upsert for stat_daily_notification_event")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(statDailyNotificationEventType, statDailyNotificationEventMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model2: unable to retrieve unique values for stat_daily_notification_event")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model2: unable to populate default values for stat_daily_notification_event")
	}

CacheNoHooks:
	if !cached {
		statDailyNotificationEventUpsertCacheMut.Lock()
		statDailyNotificationEventUpsertCache[key] = cache
		statDailyNotificationEventUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single StatDailyNotificationEvent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StatDailyNotificationEvent) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model2: no StatDailyNotificationEvent provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), statDailyNotificationEventPrimaryKeyMapping)
	sql := "DELETE FROM `stat_daily_notification_event` WHERE `date`=? AND `event`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to delete from stat_daily_notification_event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model2: failed to get rows affected by delete for stat_daily_notification_event")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q statDailyNotificationEventQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model2: no statDailyNotificationEventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to delete all from stat_daily_notification_event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model2: failed to get rows affected by deleteall for stat_daily_notification_event")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StatDailyNotificationEventSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(statDailyNotificationEventBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), statDailyNotificationEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `stat_daily_notification_event` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, statDailyNotificationEventPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to delete all from statDailyNotificationEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model2: failed to get rows affected by deleteall for stat_daily_notification_event")
	}

	if len(statDailyNotificationEventAfterDeleteHooks) != 0 {
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
func (o *StatDailyNotificationEvent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStatDailyNotificationEvent(ctx, exec, o.Date, o.Event)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StatDailyNotificationEventSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StatDailyNotificationEventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), statDailyNotificationEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `stat_daily_notification_event`.* FROM `stat_daily_notification_event` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, statDailyNotificationEventPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model2: unable to reload all in StatDailyNotificationEventSlice")
	}

	*o = slice

	return nil
}

// StatDailyNotificationEventExists checks if the StatDailyNotificationEvent row exists.
func StatDailyNotificationEventExists(ctx context.Context, exec boil.ContextExecutor, date time.Time, event string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `stat_daily_notification_event` where `date`=? AND `event`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, date, event)
	}
	row := exec.QueryRowContext(ctx, sql, date, event)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model2: unable to check if stat_daily_notification_event exists")
	}

	return exists, nil
}