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

func testListings(t *testing.T) {
	t.Parallel()

	query := Listings()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testListingsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
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

	count, err := Listings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testListingsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Listings().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Listings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testListingsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ListingSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Listings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testListingsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ListingExists(ctx, tx, o.ListingID)
	if err != nil {
		t.Errorf("Unable to check if Listing exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ListingExists to return true, but got false.")
	}
}

func testListingsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	listingFound, err := FindListing(ctx, tx, o.ListingID)
	if err != nil {
		t.Error(err)
	}

	if listingFound == nil {
		t.Error("want a record, got nil")
	}
}

func testListingsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Listings().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testListingsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Listings().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testListingsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	listingOne := &Listing{}
	listingTwo := &Listing{}
	if err = randomize.Struct(seed, listingOne, listingDBTypes, false, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}
	if err = randomize.Struct(seed, listingTwo, listingDBTypes, false, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = listingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = listingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Listings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testListingsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	listingOne := &Listing{}
	listingTwo := &Listing{}
	if err = randomize.Struct(seed, listingOne, listingDBTypes, false, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}
	if err = randomize.Struct(seed, listingTwo, listingDBTypes, false, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = listingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = listingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Listings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func listingBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Listing) error {
	*o = Listing{}
	return nil
}

func listingAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Listing) error {
	*o = Listing{}
	return nil
}

func listingAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Listing) error {
	*o = Listing{}
	return nil
}

func listingBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Listing) error {
	*o = Listing{}
	return nil
}

func listingAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Listing) error {
	*o = Listing{}
	return nil
}

func listingBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Listing) error {
	*o = Listing{}
	return nil
}

func listingAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Listing) error {
	*o = Listing{}
	return nil
}

func listingBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Listing) error {
	*o = Listing{}
	return nil
}

func listingAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Listing) error {
	*o = Listing{}
	return nil
}

func testListingsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Listing{}
	o := &Listing{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, listingDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Listing object: %s", err)
	}

	AddListingHook(boil.BeforeInsertHook, listingBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	listingBeforeInsertHooks = []ListingHook{}

	AddListingHook(boil.AfterInsertHook, listingAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	listingAfterInsertHooks = []ListingHook{}

	AddListingHook(boil.AfterSelectHook, listingAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	listingAfterSelectHooks = []ListingHook{}

	AddListingHook(boil.BeforeUpdateHook, listingBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	listingBeforeUpdateHooks = []ListingHook{}

	AddListingHook(boil.AfterUpdateHook, listingAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	listingAfterUpdateHooks = []ListingHook{}

	AddListingHook(boil.BeforeDeleteHook, listingBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	listingBeforeDeleteHooks = []ListingHook{}

	AddListingHook(boil.AfterDeleteHook, listingAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	listingAfterDeleteHooks = []ListingHook{}

	AddListingHook(boil.BeforeUpsertHook, listingBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	listingBeforeUpsertHooks = []ListingHook{}

	AddListingHook(boil.AfterUpsertHook, listingAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	listingAfterUpsertHooks = []ListingHook{}
}

func testListingsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Listings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testListingsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(listingColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Listings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testListingToManyBooks(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Listing
	var b, c Book

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, bookDBTypes, false, bookColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookDBTypes, false, bookColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"listingsbooks\" (\"resourse_id\", \"book_id\") values ($1, $2)", a.ListingID, b.BookID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"listingsbooks\" (\"resourse_id\", \"book_id\") values ($1, $2)", a.ListingID, c.BookID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Books().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BookID == b.BookID {
			bFound = true
		}
		if v.BookID == c.BookID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ListingSlice{&a}
	if err = a.L.LoadBooks(ctx, tx, false, (*[]*Listing)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Books); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Books = nil
	if err = a.L.LoadBooks(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Books); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testListingToManyListingsurls(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Listing
	var b, c Listingsurl

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, listingsurlDBTypes, false, listingsurlColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, listingsurlDBTypes, false, listingsurlColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.ListingID, a.ListingID)
	queries.Assign(&c.ListingID, a.ListingID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Listingsurls().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ListingID, b.ListingID) {
			bFound = true
		}
		if queries.Equal(v.ListingID, c.ListingID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ListingSlice{&a}
	if err = a.L.LoadListingsurls(ctx, tx, false, (*[]*Listing)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Listingsurls); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Listingsurls = nil
	if err = a.L.LoadListingsurls(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Listingsurls); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testListingToManyAddOpBooks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Listing
	var b, c, d, e Book

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, listingDBTypes, false, strmangle.SetComplement(listingPrimaryKeyColumns, listingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Book{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Book{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBooks(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.ResourseListings[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.ResourseListings[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Books[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Books[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Books().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testListingToManySetOpBooks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Listing
	var b, c, d, e Book

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, listingDBTypes, false, strmangle.SetComplement(listingPrimaryKeyColumns, listingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Book{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetBooks(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Books().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetBooks(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Books().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.ResourseListings) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.ResourseListings) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.ResourseListings[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.ResourseListings[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Books[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Books[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testListingToManyRemoveOpBooks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Listing
	var b, c, d, e Book

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, listingDBTypes, false, strmangle.SetComplement(listingPrimaryKeyColumns, listingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Book{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddBooks(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Books().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveBooks(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Books().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.ResourseListings) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.ResourseListings) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.ResourseListings[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.ResourseListings[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Books) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Books[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Books[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testListingToManyAddOpListingsurls(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Listing
	var b, c, d, e Listingsurl

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, listingDBTypes, false, strmangle.SetComplement(listingPrimaryKeyColumns, listingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Listingsurl{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, listingsurlDBTypes, false, strmangle.SetComplement(listingsurlPrimaryKeyColumns, listingsurlColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Listingsurl{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddListingsurls(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ListingID, first.ListingID) {
			t.Error("foreign key was wrong value", a.ListingID, first.ListingID)
		}
		if !queries.Equal(a.ListingID, second.ListingID) {
			t.Error("foreign key was wrong value", a.ListingID, second.ListingID)
		}

		if first.R.Listing != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Listing != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Listingsurls[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Listingsurls[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Listingsurls().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testListingToManySetOpListingsurls(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Listing
	var b, c, d, e Listingsurl

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, listingDBTypes, false, strmangle.SetComplement(listingPrimaryKeyColumns, listingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Listingsurl{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, listingsurlDBTypes, false, strmangle.SetComplement(listingsurlPrimaryKeyColumns, listingsurlColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetListingsurls(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Listingsurls().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetListingsurls(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Listingsurls().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ListingID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ListingID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ListingID, d.ListingID) {
		t.Error("foreign key was wrong value", a.ListingID, d.ListingID)
	}
	if !queries.Equal(a.ListingID, e.ListingID) {
		t.Error("foreign key was wrong value", a.ListingID, e.ListingID)
	}

	if b.R.Listing != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Listing != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Listing != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Listing != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Listingsurls[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Listingsurls[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testListingToManyRemoveOpListingsurls(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Listing
	var b, c, d, e Listingsurl

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, listingDBTypes, false, strmangle.SetComplement(listingPrimaryKeyColumns, listingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Listingsurl{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, listingsurlDBTypes, false, strmangle.SetComplement(listingsurlPrimaryKeyColumns, listingsurlColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddListingsurls(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Listingsurls().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveListingsurls(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Listingsurls().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ListingID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ListingID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Listing != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Listing != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Listing != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Listing != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Listingsurls) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Listingsurls[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Listingsurls[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testListingsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
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

func testListingsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ListingSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testListingsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Listings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	listingDBTypes = map[string]string{`ListingID`: `uuid`, `OwnerID`: `uuid`, `Title`: `character varying`, `Price`: `integer`, `Currency`: `character varying`, `Description`: `character varying`}
	_              = bytes.MinRead
)

func testListingsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(listingPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(listingAllColumns) == len(listingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Listings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, listingDBTypes, true, listingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testListingsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(listingAllColumns) == len(listingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Listing{}
	if err = randomize.Struct(seed, o, listingDBTypes, true, listingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Listings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, listingDBTypes, true, listingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(listingAllColumns, listingPrimaryKeyColumns) {
		fields = listingAllColumns
	} else {
		fields = strmangle.SetComplement(
			listingAllColumns,
			listingPrimaryKeyColumns,
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

	slice := ListingSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testListingsUpsert(t *testing.T) {
	t.Parallel()

	if len(listingAllColumns) == len(listingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Listing{}
	if err = randomize.Struct(seed, &o, listingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Listing: %s", err)
	}

	count, err := Listings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, listingDBTypes, false, listingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Listing struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Listing: %s", err)
	}

	count, err = Listings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
