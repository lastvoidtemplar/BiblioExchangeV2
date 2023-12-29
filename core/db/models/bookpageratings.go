// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Bookpagerating is an object representing the database table.
type Bookpagerating struct {
	BookRatingID string      `boil:"book_rating_id" json:"book_rating_id" toml:"book_rating_id" yaml:"book_rating_id"`
	BookID       null.String `boil:"book_id" json:"book_id,omitempty" toml:"book_id" yaml:"book_id,omitempty"`
	RatingType   null.Int    `boil:"rating_type" json:"rating_type,omitempty" toml:"rating_type" yaml:"rating_type,omitempty"`
	UserID       null.String `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`

	R *bookpageratingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bookpageratingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BookpageratingColumns = struct {
	BookRatingID string
	BookID       string
	RatingType   string
	UserID       string
}{
	BookRatingID: "book_rating_id",
	BookID:       "book_id",
	RatingType:   "rating_type",
	UserID:       "user_id",
}

var BookpageratingTableColumns = struct {
	BookRatingID string
	BookID       string
	RatingType   string
	UserID       string
}{
	BookRatingID: "bookpageratings.book_rating_id",
	BookID:       "bookpageratings.book_id",
	RatingType:   "bookpageratings.rating_type",
	UserID:       "bookpageratings.user_id",
}

// Generated where

var BookpageratingWhere = struct {
	BookRatingID whereHelperstring
	BookID       whereHelpernull_String
	RatingType   whereHelpernull_Int
	UserID       whereHelpernull_String
}{
	BookRatingID: whereHelperstring{field: "\"bookpageratings\".\"book_rating_id\""},
	BookID:       whereHelpernull_String{field: "\"bookpageratings\".\"book_id\""},
	RatingType:   whereHelpernull_Int{field: "\"bookpageratings\".\"rating_type\""},
	UserID:       whereHelpernull_String{field: "\"bookpageratings\".\"user_id\""},
}

// BookpageratingRels is where relationship names are stored.
var BookpageratingRels = struct {
	Book string
}{
	Book: "Book",
}

// bookpageratingR is where relationships are stored.
type bookpageratingR struct {
	Book *Book `boil:"Book" json:"Book" toml:"Book" yaml:"Book"`
}

// NewStruct creates a new relationship struct
func (*bookpageratingR) NewStruct() *bookpageratingR {
	return &bookpageratingR{}
}

func (r *bookpageratingR) GetBook() *Book {
	if r == nil {
		return nil
	}
	return r.Book
}

// bookpageratingL is where Load methods for each relationship are stored.
type bookpageratingL struct{}

var (
	bookpageratingAllColumns            = []string{"book_rating_id", "book_id", "rating_type", "user_id"}
	bookpageratingColumnsWithoutDefault = []string{}
	bookpageratingColumnsWithDefault    = []string{"book_rating_id", "book_id", "rating_type", "user_id"}
	bookpageratingPrimaryKeyColumns     = []string{"book_rating_id"}
	bookpageratingGeneratedColumns      = []string{}
)

type (
	// BookpageratingSlice is an alias for a slice of pointers to Bookpagerating.
	// This should almost always be used instead of []Bookpagerating.
	BookpageratingSlice []*Bookpagerating
	// BookpageratingHook is the signature for custom Bookpagerating hook methods
	BookpageratingHook func(context.Context, boil.ContextExecutor, *Bookpagerating) error

	bookpageratingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bookpageratingType                 = reflect.TypeOf(&Bookpagerating{})
	bookpageratingMapping              = queries.MakeStructMapping(bookpageratingType)
	bookpageratingPrimaryKeyMapping, _ = queries.BindMapping(bookpageratingType, bookpageratingMapping, bookpageratingPrimaryKeyColumns)
	bookpageratingInsertCacheMut       sync.RWMutex
	bookpageratingInsertCache          = make(map[string]insertCache)
	bookpageratingUpdateCacheMut       sync.RWMutex
	bookpageratingUpdateCache          = make(map[string]updateCache)
	bookpageratingUpsertCacheMut       sync.RWMutex
	bookpageratingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bookpageratingAfterSelectHooks []BookpageratingHook

var bookpageratingBeforeInsertHooks []BookpageratingHook
var bookpageratingAfterInsertHooks []BookpageratingHook

var bookpageratingBeforeUpdateHooks []BookpageratingHook
var bookpageratingAfterUpdateHooks []BookpageratingHook

var bookpageratingBeforeDeleteHooks []BookpageratingHook
var bookpageratingAfterDeleteHooks []BookpageratingHook

var bookpageratingBeforeUpsertHooks []BookpageratingHook
var bookpageratingAfterUpsertHooks []BookpageratingHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Bookpagerating) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookpageratingAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Bookpagerating) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookpageratingBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Bookpagerating) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookpageratingAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Bookpagerating) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookpageratingBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Bookpagerating) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookpageratingAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Bookpagerating) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookpageratingBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Bookpagerating) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookpageratingAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Bookpagerating) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookpageratingBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Bookpagerating) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookpageratingAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBookpageratingHook registers your hook function for all future operations.
func AddBookpageratingHook(hookPoint boil.HookPoint, bookpageratingHook BookpageratingHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		bookpageratingAfterSelectHooks = append(bookpageratingAfterSelectHooks, bookpageratingHook)
	case boil.BeforeInsertHook:
		bookpageratingBeforeInsertHooks = append(bookpageratingBeforeInsertHooks, bookpageratingHook)
	case boil.AfterInsertHook:
		bookpageratingAfterInsertHooks = append(bookpageratingAfterInsertHooks, bookpageratingHook)
	case boil.BeforeUpdateHook:
		bookpageratingBeforeUpdateHooks = append(bookpageratingBeforeUpdateHooks, bookpageratingHook)
	case boil.AfterUpdateHook:
		bookpageratingAfterUpdateHooks = append(bookpageratingAfterUpdateHooks, bookpageratingHook)
	case boil.BeforeDeleteHook:
		bookpageratingBeforeDeleteHooks = append(bookpageratingBeforeDeleteHooks, bookpageratingHook)
	case boil.AfterDeleteHook:
		bookpageratingAfterDeleteHooks = append(bookpageratingAfterDeleteHooks, bookpageratingHook)
	case boil.BeforeUpsertHook:
		bookpageratingBeforeUpsertHooks = append(bookpageratingBeforeUpsertHooks, bookpageratingHook)
	case boil.AfterUpsertHook:
		bookpageratingAfterUpsertHooks = append(bookpageratingAfterUpsertHooks, bookpageratingHook)
	}
}

// OneG returns a single bookpagerating record from the query using the global executor.
func (q bookpageratingQuery) OneG(ctx context.Context) (*Bookpagerating, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single bookpagerating record from the query.
func (q bookpageratingQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Bookpagerating, error) {
	o := &Bookpagerating{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for bookpageratings")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Bookpagerating records from the query using the global executor.
func (q bookpageratingQuery) AllG(ctx context.Context) (BookpageratingSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Bookpagerating records from the query.
func (q bookpageratingQuery) All(ctx context.Context, exec boil.ContextExecutor) (BookpageratingSlice, error) {
	var o []*Bookpagerating

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to Bookpagerating slice")
	}

	if len(bookpageratingAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Bookpagerating records in the query using the global executor
func (q bookpageratingQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Bookpagerating records in the query.
func (q bookpageratingQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count bookpageratings rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q bookpageratingQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q bookpageratingQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if bookpageratings exists")
	}

	return count > 0, nil
}

// Book pointed to by the foreign key.
func (o *Bookpagerating) Book(mods ...qm.QueryMod) bookQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"book_id\" = ?", o.BookID),
	}

	queryMods = append(queryMods, mods...)

	return Books(queryMods...)
}

// LoadBook allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (bookpageratingL) LoadBook(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBookpagerating interface{}, mods queries.Applicator) error {
	var slice []*Bookpagerating
	var object *Bookpagerating

	if singular {
		var ok bool
		object, ok = maybeBookpagerating.(*Bookpagerating)
		if !ok {
			object = new(Bookpagerating)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBookpagerating)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBookpagerating))
			}
		}
	} else {
		s, ok := maybeBookpagerating.(*[]*Bookpagerating)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBookpagerating)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBookpagerating))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bookpageratingR{}
		}
		if !queries.IsNil(object.BookID) {
			args = append(args, object.BookID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bookpageratingR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.BookID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.BookID) {
				args = append(args, obj.BookID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`books`),
		qm.WhereIn(`books.book_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Book")
	}

	var resultSlice []*Book
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Book")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for books")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for books")
	}

	if len(bookAfterSelectHooks) != 0 {
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
		object.R.Book = foreign
		if foreign.R == nil {
			foreign.R = &bookR{}
		}
		foreign.R.Bookpageratings = append(foreign.R.Bookpageratings, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.BookID, foreign.BookID) {
				local.R.Book = foreign
				if foreign.R == nil {
					foreign.R = &bookR{}
				}
				foreign.R.Bookpageratings = append(foreign.R.Bookpageratings, local)
				break
			}
		}
	}

	return nil
}

// SetBookG of the bookpagerating to the related item.
// Sets o.R.Book to related.
// Adds o to related.R.Bookpageratings.
// Uses the global database handle.
func (o *Bookpagerating) SetBookG(ctx context.Context, insert bool, related *Book) error {
	return o.SetBook(ctx, boil.GetContextDB(), insert, related)
}

// SetBook of the bookpagerating to the related item.
// Sets o.R.Book to related.
// Adds o to related.R.Bookpageratings.
func (o *Bookpagerating) SetBook(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Book) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"bookpageratings\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"book_id"}),
		strmangle.WhereClause("\"", "\"", 2, bookpageratingPrimaryKeyColumns),
	)
	values := []interface{}{related.BookID, o.BookRatingID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.BookID, related.BookID)
	if o.R == nil {
		o.R = &bookpageratingR{
			Book: related,
		}
	} else {
		o.R.Book = related
	}

	if related.R == nil {
		related.R = &bookR{
			Bookpageratings: BookpageratingSlice{o},
		}
	} else {
		related.R.Bookpageratings = append(related.R.Bookpageratings, o)
	}

	return nil
}

// RemoveBookG relationship.
// Sets o.R.Book to nil.
// Removes o from all passed in related items' relationships struct.
// Uses the global database handle.
func (o *Bookpagerating) RemoveBookG(ctx context.Context, related *Book) error {
	return o.RemoveBook(ctx, boil.GetContextDB(), related)
}

// RemoveBook relationship.
// Sets o.R.Book to nil.
// Removes o from all passed in related items' relationships struct.
func (o *Bookpagerating) RemoveBook(ctx context.Context, exec boil.ContextExecutor, related *Book) error {
	var err error

	queries.SetScanner(&o.BookID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("book_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Book = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Bookpageratings {
		if queries.Equal(o.BookID, ri.BookID) {
			continue
		}

		ln := len(related.R.Bookpageratings)
		if ln > 1 && i < ln-1 {
			related.R.Bookpageratings[i] = related.R.Bookpageratings[ln-1]
		}
		related.R.Bookpageratings = related.R.Bookpageratings[:ln-1]
		break
	}
	return nil
}

// Bookpageratings retrieves all the records using an executor.
func Bookpageratings(mods ...qm.QueryMod) bookpageratingQuery {
	mods = append(mods, qm.From("\"bookpageratings\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"bookpageratings\".*"})
	}

	return bookpageratingQuery{q}
}

// FindBookpageratingG retrieves a single record by ID.
func FindBookpageratingG(ctx context.Context, bookRatingID string, selectCols ...string) (*Bookpagerating, error) {
	return FindBookpagerating(ctx, boil.GetContextDB(), bookRatingID, selectCols...)
}

// FindBookpagerating retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBookpagerating(ctx context.Context, exec boil.ContextExecutor, bookRatingID string, selectCols ...string) (*Bookpagerating, error) {
	bookpageratingObj := &Bookpagerating{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"bookpageratings\" where \"book_rating_id\"=$1", sel,
	)

	q := queries.Raw(query, bookRatingID)

	err := q.Bind(ctx, exec, bookpageratingObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from bookpageratings")
	}

	if err = bookpageratingObj.doAfterSelectHooks(ctx, exec); err != nil {
		return bookpageratingObj, err
	}

	return bookpageratingObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Bookpagerating) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Bookpagerating) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no bookpageratings provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookpageratingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bookpageratingInsertCacheMut.RLock()
	cache, cached := bookpageratingInsertCache[key]
	bookpageratingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bookpageratingAllColumns,
			bookpageratingColumnsWithDefault,
			bookpageratingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bookpageratingType, bookpageratingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bookpageratingType, bookpageratingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"bookpageratings\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"bookpageratings\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "dbmodels: unable to insert into bookpageratings")
	}

	if !cached {
		bookpageratingInsertCacheMut.Lock()
		bookpageratingInsertCache[key] = cache
		bookpageratingInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Bookpagerating record using the global executor.
// See Update for more documentation.
func (o *Bookpagerating) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Bookpagerating.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Bookpagerating) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bookpageratingUpdateCacheMut.RLock()
	cache, cached := bookpageratingUpdateCache[key]
	bookpageratingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bookpageratingAllColumns,
			bookpageratingPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update bookpageratings, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"bookpageratings\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, bookpageratingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bookpageratingType, bookpageratingMapping, append(wl, bookpageratingPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodels: unable to update bookpageratings row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for bookpageratings")
	}

	if !cached {
		bookpageratingUpdateCacheMut.Lock()
		bookpageratingUpdateCache[key] = cache
		bookpageratingUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q bookpageratingQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q bookpageratingQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for bookpageratings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for bookpageratings")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BookpageratingSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BookpageratingSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodels: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookpageratingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"bookpageratings\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, bookpageratingPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in bookpagerating slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all bookpagerating")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Bookpagerating) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Bookpagerating) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no bookpageratings provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookpageratingColumnsWithDefault, o)

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

	bookpageratingUpsertCacheMut.RLock()
	cache, cached := bookpageratingUpsertCache[key]
	bookpageratingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bookpageratingAllColumns,
			bookpageratingColumnsWithDefault,
			bookpageratingColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			bookpageratingAllColumns,
			bookpageratingPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert bookpageratings, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bookpageratingPrimaryKeyColumns))
			copy(conflict, bookpageratingPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"bookpageratings\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bookpageratingType, bookpageratingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bookpageratingType, bookpageratingMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to upsert bookpageratings")
	}

	if !cached {
		bookpageratingUpsertCacheMut.Lock()
		bookpageratingUpsertCache[key] = cache
		bookpageratingUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Bookpagerating record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Bookpagerating) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Bookpagerating record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Bookpagerating) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no Bookpagerating provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bookpageratingPrimaryKeyMapping)
	sql := "DELETE FROM \"bookpageratings\" WHERE \"book_rating_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete from bookpageratings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for bookpageratings")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q bookpageratingQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q bookpageratingQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no bookpageratingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from bookpageratings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for bookpageratings")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o BookpageratingSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BookpageratingSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bookpageratingBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookpageratingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"bookpageratings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bookpageratingPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from bookpagerating slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for bookpageratings")
	}

	if len(bookpageratingAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Bookpagerating) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("dbmodels: no Bookpagerating provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Bookpagerating) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBookpagerating(ctx, exec, o.BookRatingID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BookpageratingSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("dbmodels: empty BookpageratingSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BookpageratingSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BookpageratingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookpageratingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"bookpageratings\".* FROM \"bookpageratings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bookpageratingPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in BookpageratingSlice")
	}

	*o = slice

	return nil
}

// BookpageratingExistsG checks if the Bookpagerating row exists.
func BookpageratingExistsG(ctx context.Context, bookRatingID string) (bool, error) {
	return BookpageratingExists(ctx, boil.GetContextDB(), bookRatingID)
}

// BookpageratingExists checks if the Bookpagerating row exists.
func BookpageratingExists(ctx context.Context, exec boil.ContextExecutor, bookRatingID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"bookpageratings\" where \"book_rating_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, bookRatingID)
	}
	row := exec.QueryRowContext(ctx, sql, bookRatingID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if bookpageratings exists")
	}

	return exists, nil
}

// Exists checks if the Bookpagerating row exists.
func (o *Bookpagerating) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BookpageratingExists(ctx, exec, o.BookRatingID)
}
