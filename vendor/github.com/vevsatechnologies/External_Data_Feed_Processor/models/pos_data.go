// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v6"
)

// PosDatum is an object representing the database table.
type PosDatum struct {
	ID                   int          `boil:"id" json:"id" toml:"id" yaml:"id"`
	Posid                null.String  `boil:"posid" json:"posid,omitempty" toml:"posid" yaml:"posid,omitempty"`
	Apienabled           null.String  `boil:"apienabled" json:"apienabled,omitempty" toml:"apienabled" yaml:"apienabled,omitempty"`
	Apiversionssupported null.Float64 `boil:"apiversionssupported" json:"apiversionssupported,omitempty" toml:"apiversionssupported" yaml:"apiversionssupported,omitempty"`
	Network              null.String  `boil:"network" json:"network,omitempty" toml:"network" yaml:"network,omitempty"`
	URL                  null.String  `boil:"url" json:"url,omitempty" toml:"url" yaml:"url,omitempty"`
	Launched             null.String  `boil:"launched" json:"launched,omitempty" toml:"launched" yaml:"launched,omitempty"`
	Lastupdated          null.String  `boil:"lastupdated" json:"lastupdated,omitempty" toml:"lastupdated" yaml:"lastupdated,omitempty"`
	Immature             null.String  `boil:"immature" json:"immature,omitempty" toml:"immature" yaml:"immature,omitempty"`
	Live                 null.String  `boil:"live" json:"live,omitempty" toml:"live" yaml:"live,omitempty"`
	Voted                null.Float64 `boil:"voted" json:"voted,omitempty" toml:"voted" yaml:"voted,omitempty"`
	Missed               null.Float64 `boil:"missed" json:"missed,omitempty" toml:"missed" yaml:"missed,omitempty"`
	Poolfees             null.Float64 `boil:"poolfees" json:"poolfees,omitempty" toml:"poolfees" yaml:"poolfees,omitempty"`
	Proportionlive       null.Float64 `boil:"proportionlive" json:"proportionlive,omitempty" toml:"proportionlive" yaml:"proportionlive,omitempty"`
	Proportionmissed     null.Float64 `boil:"proportionmissed" json:"proportionmissed,omitempty" toml:"proportionmissed" yaml:"proportionmissed,omitempty"`
	Usercount            null.Float64 `boil:"usercount" json:"usercount,omitempty" toml:"usercount" yaml:"usercount,omitempty"`
	Usercountactive      null.Float64 `boil:"usercountactive" json:"usercountactive,omitempty" toml:"usercountactive" yaml:"usercountactive,omitempty"`
	Timestamp            null.String  `boil:"timestamp" json:"timestamp,omitempty" toml:"timestamp" yaml:"timestamp,omitempty"`

	R *posDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L posDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PosDatumColumns = struct {
	ID                   string
	Posid                string
	Apienabled           string
	Apiversionssupported string
	Network              string
	URL                  string
	Launched             string
	Lastupdated          string
	Immature             string
	Live                 string
	Voted                string
	Missed               string
	Poolfees             string
	Proportionlive       string
	Proportionmissed     string
	Usercount            string
	Usercountactive      string
	Timestamp            string
}{
	ID:                   "id",
	Posid:                "posid",
	Apienabled:           "apienabled",
	Apiversionssupported: "apiversionssupported",
	Network:              "network",
	URL:                  "url",
	Launched:             "launched",
	Lastupdated:          "lastupdated",
	Immature:             "immature",
	Live:                 "live",
	Voted:                "voted",
	Missed:               "missed",
	Poolfees:             "poolfees",
	Proportionlive:       "proportionlive",
	Proportionmissed:     "proportionmissed",
	Usercount:            "usercount",
	Usercountactive:      "usercountactive",
	Timestamp:            "timestamp",
}

// posDatumR is where relationships are stored.
type posDatumR struct {
}

// posDatumL is where Load methods for each relationship are stored.
type posDatumL struct{}

var (
	posDatumColumns               = []string{"id", "posid", "apienabled", "apiversionssupported", "network", "url", "launched", "lastupdated", "immature", "live", "voted", "missed", "poolfees", "proportionlive", "proportionmissed", "usercount", "usercountactive", "timestamp"}
	posDatumColumnsWithoutDefault = []string{"posid", "apienabled", "apiversionssupported", "network", "url", "launched", "lastupdated", "immature", "live", "voted", "missed", "poolfees", "proportionlive", "proportionmissed", "usercount", "usercountactive", "timestamp"}
	posDatumColumnsWithDefault    = []string{"id"}
	posDatumPrimaryKeyColumns     = []string{"id"}
)

type (
	// PosDatumSlice is an alias for a slice of pointers to PosDatum.
	// This should generally be used opposed to []PosDatum.
	PosDatumSlice []*PosDatum
	// PosDatumHook is the signature for custom PosDatum hook methods
	PosDatumHook func(boil.Executor, *PosDatum) error

	posDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	posDatumType                 = reflect.TypeOf(&PosDatum{})
	posDatumMapping              = queries.MakeStructMapping(posDatumType)
	posDatumPrimaryKeyMapping, _ = queries.BindMapping(posDatumType, posDatumMapping, posDatumPrimaryKeyColumns)
	posDatumInsertCacheMut       sync.RWMutex
	posDatumInsertCache          = make(map[string]insertCache)
	posDatumUpdateCacheMut       sync.RWMutex
	posDatumUpdateCache          = make(map[string]updateCache)
	posDatumUpsertCacheMut       sync.RWMutex
	posDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var posDatumBeforeInsertHooks []PosDatumHook
var posDatumBeforeUpdateHooks []PosDatumHook
var posDatumBeforeDeleteHooks []PosDatumHook
var posDatumBeforeUpsertHooks []PosDatumHook

var posDatumAfterInsertHooks []PosDatumHook
var posDatumAfterSelectHooks []PosDatumHook
var posDatumAfterUpdateHooks []PosDatumHook
var posDatumAfterDeleteHooks []PosDatumHook
var posDatumAfterUpsertHooks []PosDatumHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PosDatum) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range posDatumBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PosDatum) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range posDatumBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PosDatum) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range posDatumBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PosDatum) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range posDatumBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PosDatum) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range posDatumAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PosDatum) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range posDatumAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PosDatum) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range posDatumAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PosDatum) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range posDatumAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PosDatum) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range posDatumAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPosDatumHook registers your hook function for all future operations.
func AddPosDatumHook(hookPoint boil.HookPoint, posDatumHook PosDatumHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		posDatumBeforeInsertHooks = append(posDatumBeforeInsertHooks, posDatumHook)
	case boil.BeforeUpdateHook:
		posDatumBeforeUpdateHooks = append(posDatumBeforeUpdateHooks, posDatumHook)
	case boil.BeforeDeleteHook:
		posDatumBeforeDeleteHooks = append(posDatumBeforeDeleteHooks, posDatumHook)
	case boil.BeforeUpsertHook:
		posDatumBeforeUpsertHooks = append(posDatumBeforeUpsertHooks, posDatumHook)
	case boil.AfterInsertHook:
		posDatumAfterInsertHooks = append(posDatumAfterInsertHooks, posDatumHook)
	case boil.AfterSelectHook:
		posDatumAfterSelectHooks = append(posDatumAfterSelectHooks, posDatumHook)
	case boil.AfterUpdateHook:
		posDatumAfterUpdateHooks = append(posDatumAfterUpdateHooks, posDatumHook)
	case boil.AfterDeleteHook:
		posDatumAfterDeleteHooks = append(posDatumAfterDeleteHooks, posDatumHook)
	case boil.AfterUpsertHook:
		posDatumAfterUpsertHooks = append(posDatumAfterUpsertHooks, posDatumHook)
	}
}

// OneP returns a single posDatum record from the query, and panics on error.
func (q posDatumQuery) OneP() *PosDatum {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single posDatum record from the query.
func (q posDatumQuery) One() (*PosDatum, error) {
	o := &PosDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pos_data")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all PosDatum records from the query, and panics on error.
func (q posDatumQuery) AllP() PosDatumSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all PosDatum records from the query.
func (q posDatumQuery) All() (PosDatumSlice, error) {
	var o []*PosDatum

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PosDatum slice")
	}

	if len(posDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all PosDatum records in the query, and panics on error.
func (q posDatumQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all PosDatum records in the query.
func (q posDatumQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pos_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q posDatumQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q posDatumQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pos_data exists")
	}

	return count > 0, nil
}

// PosDataG retrieves all records.
func PosDataG(mods ...qm.QueryMod) posDatumQuery {
	return PosData(boil.GetDB(), mods...)
}

// PosData retrieves all the records using an executor.
func PosData(exec boil.Executor, mods ...qm.QueryMod) posDatumQuery {
	mods = append(mods, qm.From("\"pos_data\""))
	return posDatumQuery{NewQuery(exec, mods...)}
}

// FindPosDatumG retrieves a single record by ID.
func FindPosDatumG(id int, selectCols ...string) (*PosDatum, error) {
	return FindPosDatum(boil.GetDB(), id, selectCols...)
}

// FindPosDatumGP retrieves a single record by ID, and panics on error.
func FindPosDatumGP(id int, selectCols ...string) *PosDatum {
	retobj, err := FindPosDatum(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindPosDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPosDatum(exec boil.Executor, id int, selectCols ...string) (*PosDatum, error) {
	posDatumObj := &PosDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pos_data\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(posDatumObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pos_data")
	}

	return posDatumObj, nil
}

// FindPosDatumP retrieves a single record by ID with an executor, and panics on error.
func FindPosDatumP(exec boil.Executor, id int, selectCols ...string) *PosDatum {
	retobj, err := FindPosDatum(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *PosDatum) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *PosDatum) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *PosDatum) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *PosDatum) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no pos_data provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(posDatumColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	posDatumInsertCacheMut.RLock()
	cache, cached := posDatumInsertCache[key]
	posDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			posDatumColumns,
			posDatumColumnsWithDefault,
			posDatumColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(posDatumType, posDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(posDatumType, posDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"pos_data\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"pos_data\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into pos_data")
	}

	if !cached {
		posDatumInsertCacheMut.Lock()
		posDatumInsertCache[key] = cache
		posDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single PosDatum record. See Update for
// whitelist behavior description.
func (o *PosDatum) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single PosDatum record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *PosDatum) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the PosDatum, and panics on error.
// See Update for whitelist behavior description.
func (o *PosDatum) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the PosDatum.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *PosDatum) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	posDatumUpdateCacheMut.RLock()
	cache, cached := posDatumUpdateCache[key]
	posDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			posDatumColumns,
			posDatumPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update pos_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pos_data\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, posDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(posDatumType, posDatumMapping, append(wl, posDatumPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update pos_data row")
	}

	if !cached {
		posDatumUpdateCacheMut.Lock()
		posDatumUpdateCache[key] = cache
		posDatumUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q posDatumQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q posDatumQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for pos_data")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PosDatumSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o PosDatumSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o PosDatumSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PosDatumSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), posDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"pos_data\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, posDatumPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in posDatum slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *PosDatum) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *PosDatum) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *PosDatum) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *PosDatum) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no pos_data provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(posDatumColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
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
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	posDatumUpsertCacheMut.RLock()
	cache, cached := posDatumUpsertCache[key]
	posDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			posDatumColumns,
			posDatumColumnsWithDefault,
			posDatumColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			posDatumColumns,
			posDatumPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert pos_data, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(posDatumPrimaryKeyColumns))
			copy(conflict, posDatumPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"pos_data\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(posDatumType, posDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(posDatumType, posDatumMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert pos_data")
	}

	if !cached {
		posDatumUpsertCacheMut.Lock()
		posDatumUpsertCache[key] = cache
		posDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single PosDatum record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *PosDatum) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single PosDatum record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *PosDatum) DeleteG() error {
	if o == nil {
		return errors.New("models: no PosDatum provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single PosDatum record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *PosDatum) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single PosDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PosDatum) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no PosDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), posDatumPrimaryKeyMapping)
	sql := "DELETE FROM \"pos_data\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from pos_data")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q posDatumQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q posDatumQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no posDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from pos_data")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o PosDatumSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o PosDatumSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no PosDatum slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o PosDatumSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PosDatumSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no PosDatum slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(posDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), posDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"pos_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, posDatumPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from posDatum slice")
	}

	if len(posDatumAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *PosDatum) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *PosDatum) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *PosDatum) ReloadG() error {
	if o == nil {
		return errors.New("models: no PosDatum provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PosDatum) Reload(exec boil.Executor) error {
	ret, err := FindPosDatum(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PosDatumSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PosDatumSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PosDatumSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty PosDatumSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PosDatumSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	posData := PosDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), posDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"pos_data\".* FROM \"pos_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, posDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&posData)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PosDatumSlice")
	}

	*o = posData

	return nil
}

// PosDatumExists checks if the PosDatum row exists.
func PosDatumExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"pos_data\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pos_data exists")
	}

	return exists, nil
}

// PosDatumExistsG checks if the PosDatum row exists.
func PosDatumExistsG(id int) (bool, error) {
	return PosDatumExists(boil.GetDB(), id)
}

// PosDatumExistsGP checks if the PosDatum row exists. Panics on error.
func PosDatumExistsGP(id int) bool {
	e, err := PosDatumExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// PosDatumExistsP checks if the PosDatum row exists. Panics on error.
func PosDatumExistsP(exec boil.Executor, id int) bool {
	e, err := PosDatumExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
