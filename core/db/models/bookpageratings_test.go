// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBookpageratings(t *testing.T) {
	t.Parallel()

	query := Bookpageratings()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBookpageratingsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Bookpageratings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookpageratingsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Bookpageratings().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Bookpageratings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookpageratingsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookpageratingSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Bookpageratings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookpageratingsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BookpageratingExists(ctx, tx, o.BookRatingID)
	if err != nil {
		t.Errorf("Unable to check if Bookpagerating exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BookpageratingExists to return true, but got false.")
	}
}

func testBookpageratingsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	bookpageratingFound, err := FindBookpagerating(ctx, tx, o.BookRatingID)
	if err != nil {
		t.Error(err)
	}

	if bookpageratingFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBookpageratingsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Bookpageratings().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBookpageratingsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Bookpageratings().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBookpageratingsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	bookpageratingOne := &Bookpagerating{}
	bookpageratingTwo := &Bookpagerating{}
	if err = randomize.Struct(seed, bookpageratingOne, bookpageratingDBTypes, false, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}
	if err = randomize.Struct(seed, bookpageratingTwo, bookpageratingDBTypes, false, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookpageratingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookpageratingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Bookpageratings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBookpageratingsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	bookpageratingOne := &Bookpagerating{}
	bookpageratingTwo := &Bookpagerating{}
	if err = randomize.Struct(seed, bookpageratingOne, bookpageratingDBTypes, false, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}
	if err = randomize.Struct(seed, bookpageratingTwo, bookpageratingDBTypes, false, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookpageratingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookpageratingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bookpageratings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func bookpageratingBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Bookpagerating) error {
	*o = Bookpagerating{}
	return nil
}

func bookpageratingAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Bookpagerating) error {
	*o = Bookpagerating{}
	return nil
}

func bookpageratingAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Bookpagerating) error {
	*o = Bookpagerating{}
	return nil
}

func bookpageratingBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Bookpagerating) error {
	*o = Bookpagerating{}
	return nil
}

func bookpageratingAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Bookpagerating) error {
	*o = Bookpagerating{}
	return nil
}

func bookpageratingBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Bookpagerating) error {
	*o = Bookpagerating{}
	return nil
}

func bookpageratingAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Bookpagerating) error {
	*o = Bookpagerating{}
	return nil
}

func bookpageratingBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Bookpagerating) error {
	*o = Bookpagerating{}
	return nil
}

func bookpageratingAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Bookpagerating) error {
	*o = Bookpagerating{}
	return nil
}

func testBookpageratingsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Bookpagerating{}
	o := &Bookpagerating{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Bookpagerating object: %s", err)
	}

	AddBookpageratingHook(boil.BeforeInsertHook, bookpageratingBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	bookpageratingBeforeInsertHooks = []BookpageratingHook{}

	AddBookpageratingHook(boil.AfterInsertHook, bookpageratingAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	bookpageratingAfterInsertHooks = []BookpageratingHook{}

	AddBookpageratingHook(boil.AfterSelectHook, bookpageratingAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	bookpageratingAfterSelectHooks = []BookpageratingHook{}

	AddBookpageratingHook(boil.BeforeUpdateHook, bookpageratingBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	bookpageratingBeforeUpdateHooks = []BookpageratingHook{}

	AddBookpageratingHook(boil.AfterUpdateHook, bookpageratingAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	bookpageratingAfterUpdateHooks = []BookpageratingHook{}

	AddBookpageratingHook(boil.BeforeDeleteHook, bookpageratingBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	bookpageratingBeforeDeleteHooks = []BookpageratingHook{}

	AddBookpageratingHook(boil.AfterDeleteHook, bookpageratingAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	bookpageratingAfterDeleteHooks = []BookpageratingHook{}

	AddBookpageratingHook(boil.BeforeUpsertHook, bookpageratingBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	bookpageratingBeforeUpsertHooks = []BookpageratingHook{}

	AddBookpageratingHook(boil.AfterUpsertHook, bookpageratingAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	bookpageratingAfterUpsertHooks = []BookpageratingHook{}
}

func testBookpageratingsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bookpageratings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookpageratingsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(bookpageratingColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Bookpageratings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookpageratingToOneBookUsingBook(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Bookpagerating
	var foreign Book

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, bookDBTypes, false, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.BookID, foreign.BookID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Book().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.BookID, foreign.BookID) {
		t.Errorf("want: %v, got %v", foreign.BookID, check.BookID)
	}

	ranAfterSelectHook := false
	AddBookHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Book) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := BookpageratingSlice{&local}
	if err = local.L.LoadBook(ctx, tx, false, (*[]*Bookpagerating)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Book == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Book = nil
	if err = local.L.LoadBook(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Book == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testBookpageratingToOneSetOpBookUsingBook(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Bookpagerating
	var b, c Book

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookpageratingDBTypes, false, strmangle.SetComplement(bookpageratingPrimaryKeyColumns, bookpageratingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Book{&b, &c} {
		err = a.SetBook(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Book != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Bookpageratings[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.BookID, x.BookID) {
			t.Error("foreign key was wrong value", a.BookID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BookID))
		reflect.Indirect(reflect.ValueOf(&a.BookID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.BookID, x.BookID) {
			t.Error("foreign key was wrong value", a.BookID, x.BookID)
		}
	}
}

func testBookpageratingToOneRemoveOpBookUsingBook(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Bookpagerating
	var b Book

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookpageratingDBTypes, false, strmangle.SetComplement(bookpageratingPrimaryKeyColumns, bookpageratingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetBook(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveBook(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Book().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Book != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.BookID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Bookpageratings) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testBookpageratingsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBookpageratingsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookpageratingSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBookpageratingsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Bookpageratings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	bookpageratingDBTypes = map[string]string{`BookRatingID`: `uuid`, `BookID`: `uuid`, `RatingType`: `integer`, `UserID`: `uuid`}
	_                     = bytes.MinRead
)

func testBookpageratingsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(bookpageratingPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(bookpageratingAllColumns) == len(bookpageratingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bookpageratings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBookpageratingsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(bookpageratingAllColumns) == len(bookpageratingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Bookpagerating{}
	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Bookpageratings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookpageratingDBTypes, true, bookpageratingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(bookpageratingAllColumns, bookpageratingPrimaryKeyColumns) {
		fields = bookpageratingAllColumns
	} else {
		fields = strmangle.SetComplement(
			bookpageratingAllColumns,
			bookpageratingPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := BookpageratingSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBookpageratingsUpsert(t *testing.T) {
	t.Parallel()

	if len(bookpageratingAllColumns) == len(bookpageratingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Bookpagerating{}
	if err = randomize.Struct(seed, &o, bookpageratingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Bookpagerating: %s", err)
	}

	count, err := Bookpageratings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, bookpageratingDBTypes, false, bookpageratingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Bookpagerating struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Bookpagerating: %s", err)
	}

	count, err = Bookpageratings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
