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

// UserGoogle is an object representing the database table.
type UserGoogle struct {
	GoogleID  string    `boil:"google_id" json:"google_id" toml:"google_id" yaml:"google_id"`
	UserID    uint      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *userGoogleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userGoogleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserGoogleColumns = struct {
	GoogleID  string
	UserID    string
	CreatedAt string
	UpdatedAt string
}{
	GoogleID:  "google_id",
	UserID:    "user_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var UserGoogleTableColumns = struct {
	GoogleID  string
	UserID    string
	CreatedAt string
	UpdatedAt string
}{
	GoogleID:  "user_google.google_id",
	UserID:    "user_google.user_id",
	CreatedAt: "user_google.created_at",
	UpdatedAt: "user_google.updated_at",
}

// Generated where

var UserGoogleWhere = struct {
	GoogleID  whereHelperstring
	UserID    whereHelperuint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	GoogleID:  whereHelperstring{field: "`user_google`.`google_id`"},
	UserID:    whereHelperuint{field: "`user_google`.`user_id`"},
	CreatedAt: whereHelpertime_Time{field: "`user_google`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`user_google`.`updated_at`"},
}

// UserGoogleRels is where relationship names are stored.
var UserGoogleRels = struct {
}{}

// userGoogleR is where relationships are stored.
type userGoogleR struct {
}

// NewStruct creates a new relationship struct
func (*userGoogleR) NewStruct() *userGoogleR {
	return &userGoogleR{}
}

// userGoogleL is where Load methods for each relationship are stored.
type userGoogleL struct{}

var (
	userGoogleAllColumns            = []string{"google_id", "user_id", "created_at", "updated_at"}
	userGoogleColumnsWithoutDefault = []string{"google_id", "user_id", "created_at", "updated_at"}
	userGoogleColumnsWithDefault    = []string{}
	userGooglePrimaryKeyColumns     = []string{"google_id"}
)

type (
	// UserGoogleSlice is an alias for a slice of pointers to UserGoogle.
	// This should almost always be used instead of []UserGoogle.
	UserGoogleSlice []*UserGoogle
	// UserGoogleHook is the signature for custom UserGoogle hook methods
	UserGoogleHook func(context.Context, boil.ContextExecutor, *UserGoogle) error

	userGoogleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userGoogleType                 = reflect.TypeOf(&UserGoogle{})
	userGoogleMapping              = queries.MakeStructMapping(userGoogleType)
	userGooglePrimaryKeyMapping, _ = queries.BindMapping(userGoogleType, userGoogleMapping, userGooglePrimaryKeyColumns)
	userGoogleInsertCacheMut       sync.RWMutex
	userGoogleInsertCache          = make(map[string]insertCache)
	userGoogleUpdateCacheMut       sync.RWMutex
	userGoogleUpdateCache          = make(map[string]updateCache)
	userGoogleUpsertCacheMut       sync.RWMutex
	userGoogleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userGoogleBeforeInsertHooks []UserGoogleHook
var userGoogleBeforeUpdateHooks []UserGoogleHook
var userGoogleBeforeDeleteHooks []UserGoogleHook
var userGoogleBeforeUpsertHooks []UserGoogleHook

var userGoogleAfterInsertHooks []UserGoogleHook
var userGoogleAfterSelectHooks []UserGoogleHook
var userGoogleAfterUpdateHooks []UserGoogleHook
var userGoogleAfterDeleteHooks []UserGoogleHook
var userGoogleAfterUpsertHooks []UserGoogleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserGoogle) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGoogleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserGoogle) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGoogleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserGoogle) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGoogleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserGoogle) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGoogleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserGoogle) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGoogleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserGoogle) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGoogleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserGoogle) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGoogleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserGoogle) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGoogleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserGoogle) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userGoogleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserGoogleHook registers your hook function for all future operations.
func AddUserGoogleHook(hookPoint boil.HookPoint, userGoogleHook UserGoogleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		userGoogleBeforeInsertHooks = append(userGoogleBeforeInsertHooks, userGoogleHook)
	case boil.BeforeUpdateHook:
		userGoogleBeforeUpdateHooks = append(userGoogleBeforeUpdateHooks, userGoogleHook)
	case boil.BeforeDeleteHook:
		userGoogleBeforeDeleteHooks = append(userGoogleBeforeDeleteHooks, userGoogleHook)
	case boil.BeforeUpsertHook:
		userGoogleBeforeUpsertHooks = append(userGoogleBeforeUpsertHooks, userGoogleHook)
	case boil.AfterInsertHook:
		userGoogleAfterInsertHooks = append(userGoogleAfterInsertHooks, userGoogleHook)
	case boil.AfterSelectHook:
		userGoogleAfterSelectHooks = append(userGoogleAfterSelectHooks, userGoogleHook)
	case boil.AfterUpdateHook:
		userGoogleAfterUpdateHooks = append(userGoogleAfterUpdateHooks, userGoogleHook)
	case boil.AfterDeleteHook:
		userGoogleAfterDeleteHooks = append(userGoogleAfterDeleteHooks, userGoogleHook)
	case boil.AfterUpsertHook:
		userGoogleAfterUpsertHooks = append(userGoogleAfterUpsertHooks, userGoogleHook)
	}
}

// One returns a single userGoogle record from the query.
func (q userGoogleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserGoogle, error) {
	o := &UserGoogle{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model2: failed to execute a one query for user_google")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserGoogle records from the query.
func (q userGoogleQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserGoogleSlice, error) {
	var o []*UserGoogle

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model2: failed to assign all query results to UserGoogle slice")
	}

	if len(userGoogleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserGoogle records in the query.
func (q userGoogleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model2: failed to count user_google rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userGoogleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model2: failed to check if user_google exists")
	}

	return count > 0, nil
}

// UserGoogles retrieves all the records using an executor.
func UserGoogles(mods ...qm.QueryMod) userGoogleQuery {
	mods = append(mods, qm.From("`user_google`"))
	return userGoogleQuery{NewQuery(mods...)}
}

// FindUserGoogle retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserGoogle(ctx context.Context, exec boil.ContextExecutor, googleID string, selectCols ...string) (*UserGoogle, error) {
	userGoogleObj := &UserGoogle{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `user_google` where `google_id`=?", sel,
	)

	q := queries.Raw(query, googleID)

	err := q.Bind(ctx, exec, userGoogleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model2: unable to select from user_google")
	}

	if err = userGoogleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userGoogleObj, err
	}

	return userGoogleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserGoogle) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model2: no user_google provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userGoogleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userGoogleInsertCacheMut.RLock()
	cache, cached := userGoogleInsertCache[key]
	userGoogleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userGoogleAllColumns,
			userGoogleColumnsWithDefault,
			userGoogleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userGoogleType, userGoogleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userGoogleType, userGoogleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `user_google` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `user_google` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `user_google` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, userGooglePrimaryKeyColumns))
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
		return errors.Wrap(err, "model2: unable to insert into user_google")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.GoogleID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model2: unable to populate default values for user_google")
	}

CacheNoHooks:
	if !cached {
		userGoogleInsertCacheMut.Lock()
		userGoogleInsertCache[key] = cache
		userGoogleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserGoogle.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserGoogle) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userGoogleUpdateCacheMut.RLock()
	cache, cached := userGoogleUpdateCache[key]
	userGoogleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userGoogleAllColumns,
			userGooglePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model2: unable to update user_google, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `user_google` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, userGooglePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userGoogleType, userGoogleMapping, append(wl, userGooglePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model2: unable to update user_google row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model2: failed to get rows affected by update for user_google")
	}

	if !cached {
		userGoogleUpdateCacheMut.Lock()
		userGoogleUpdateCache[key] = cache
		userGoogleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userGoogleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to update all for user_google")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to retrieve rows affected for user_google")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserGoogleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userGooglePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `user_google` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userGooglePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to update all in userGoogle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to retrieve rows affected all in update all userGoogle")
	}
	return rowsAff, nil
}

var mySQLUserGoogleUniqueColumns = []string{
	"google_id",
	"user_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserGoogle) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model2: no user_google provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userGoogleColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLUserGoogleUniqueColumns, o)

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

	userGoogleUpsertCacheMut.RLock()
	cache, cached := userGoogleUpsertCache[key]
	userGoogleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userGoogleAllColumns,
			userGoogleColumnsWithDefault,
			userGoogleColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userGoogleAllColumns,
			userGooglePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("model2: unable to upsert user_google, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`user_google`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `user_google` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(userGoogleType, userGoogleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userGoogleType, userGoogleMapping, ret)
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
		return errors.Wrap(err, "model2: unable to upsert for user_google")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(userGoogleType, userGoogleMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model2: unable to retrieve unique values for user_google")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model2: unable to populate default values for user_google")
	}

CacheNoHooks:
	if !cached {
		userGoogleUpsertCacheMut.Lock()
		userGoogleUpsertCache[key] = cache
		userGoogleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserGoogle record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserGoogle) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model2: no UserGoogle provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userGooglePrimaryKeyMapping)
	sql := "DELETE FROM `user_google` WHERE `google_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to delete from user_google")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model2: failed to get rows affected by delete for user_google")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userGoogleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model2: no userGoogleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to delete all from user_google")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model2: failed to get rows affected by deleteall for user_google")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserGoogleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userGoogleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userGooglePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `user_google` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userGooglePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model2: unable to delete all from userGoogle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model2: failed to get rows affected by deleteall for user_google")
	}

	if len(userGoogleAfterDeleteHooks) != 0 {
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
func (o *UserGoogle) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserGoogle(ctx, exec, o.GoogleID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserGoogleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserGoogleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userGooglePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `user_google`.* FROM `user_google` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userGooglePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model2: unable to reload all in UserGoogleSlice")
	}

	*o = slice

	return nil
}

// UserGoogleExists checks if the UserGoogle row exists.
func UserGoogleExists(ctx context.Context, exec boil.ContextExecutor, googleID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `user_google` where `google_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, googleID)
	}
	row := exec.QueryRowContext(ctx, sql, googleID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model2: unable to check if user_google exists")
	}

	return exists, nil
}