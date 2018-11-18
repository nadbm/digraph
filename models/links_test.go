// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testLinks(t *testing.T) {
	t.Parallel()

	query := Links()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLinksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
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

	count, err := Links().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLinksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Links().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Links().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLinksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LinkSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Links().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLinksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LinkExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Link exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LinkExists to return true, but got false.")
	}
}

func testLinksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	linkFound, err := FindLink(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if linkFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLinksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Links().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLinksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Links().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLinksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	linkOne := &Link{}
	linkTwo := &Link{}
	if err = randomize.Struct(seed, linkOne, linkDBTypes, false, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}
	if err = randomize.Struct(seed, linkTwo, linkDBTypes, false, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = linkOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = linkTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Links().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLinksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	linkOne := &Link{}
	linkTwo := &Link{}
	if err = randomize.Struct(seed, linkOne, linkDBTypes, false, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}
	if err = randomize.Struct(seed, linkTwo, linkDBTypes, false, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = linkOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = linkTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Links().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func linkBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Link) error {
	*o = Link{}
	return nil
}

func linkAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Link) error {
	*o = Link{}
	return nil
}

func linkAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Link) error {
	*o = Link{}
	return nil
}

func linkBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Link) error {
	*o = Link{}
	return nil
}

func linkAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Link) error {
	*o = Link{}
	return nil
}

func linkBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Link) error {
	*o = Link{}
	return nil
}

func linkAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Link) error {
	*o = Link{}
	return nil
}

func linkBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Link) error {
	*o = Link{}
	return nil
}

func linkAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Link) error {
	*o = Link{}
	return nil
}

func testLinksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Link{}
	o := &Link{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, linkDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Link object: %s", err)
	}

	AddLinkHook(boil.BeforeInsertHook, linkBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	linkBeforeInsertHooks = []LinkHook{}

	AddLinkHook(boil.AfterInsertHook, linkAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	linkAfterInsertHooks = []LinkHook{}

	AddLinkHook(boil.AfterSelectHook, linkAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	linkAfterSelectHooks = []LinkHook{}

	AddLinkHook(boil.BeforeUpdateHook, linkBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	linkBeforeUpdateHooks = []LinkHook{}

	AddLinkHook(boil.AfterUpdateHook, linkAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	linkAfterUpdateHooks = []LinkHook{}

	AddLinkHook(boil.BeforeDeleteHook, linkBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	linkBeforeDeleteHooks = []LinkHook{}

	AddLinkHook(boil.AfterDeleteHook, linkAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	linkAfterDeleteHooks = []LinkHook{}

	AddLinkHook(boil.BeforeUpsertHook, linkBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	linkBeforeUpsertHooks = []LinkHook{}

	AddLinkHook(boil.AfterUpsertHook, linkAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	linkAfterUpsertHooks = []LinkHook{}
}

func testLinksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Links().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLinksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(linkColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Links().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLinkToManyChildTopicsLinks(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Link
	var b, c TopicsLink

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, topicsLinkDBTypes, false, topicsLinkColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, topicsLinkDBTypes, false, topicsLinkColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ChildID = a.ID
	c.ChildID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	topicsLink, err := a.ChildTopicsLinks().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range topicsLink {
		if v.ChildID == b.ChildID {
			bFound = true
		}
		if v.ChildID == c.ChildID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := LinkSlice{&a}
	if err = a.L.LoadChildTopicsLinks(ctx, tx, false, (*[]*Link)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChildTopicsLinks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ChildTopicsLinks = nil
	if err = a.L.LoadChildTopicsLinks(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChildTopicsLinks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", topicsLink)
	}
}

func testLinkToManyAddOpChildTopicsLinks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Link
	var b, c, d, e TopicsLink

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, linkDBTypes, false, strmangle.SetComplement(linkPrimaryKeyColumns, linkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*TopicsLink{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, topicsLinkDBTypes, false, strmangle.SetComplement(topicsLinkPrimaryKeyColumns, topicsLinkColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*TopicsLink{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddChildTopicsLinks(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ChildID {
			t.Error("foreign key was wrong value", a.ID, first.ChildID)
		}
		if a.ID != second.ChildID {
			t.Error("foreign key was wrong value", a.ID, second.ChildID)
		}

		if first.R.Child != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Child != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ChildTopicsLinks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ChildTopicsLinks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ChildTopicsLinks().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testLinkToOneOrganizationUsingOrganization(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Link
	var foreign Organization

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, linkDBTypes, false, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, organizationDBTypes, false, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.OrganizationID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Organization().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := LinkSlice{&local}
	if err = local.L.LoadOrganization(ctx, tx, false, (*[]*Link)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Organization == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Organization = nil
	if err = local.L.LoadOrganization(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Organization == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testLinkToOneSetOpOrganizationUsingOrganization(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Link
	var b, c Organization

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, linkDBTypes, false, strmangle.SetComplement(linkPrimaryKeyColumns, linkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organizationDBTypes, false, strmangle.SetComplement(organizationPrimaryKeyColumns, organizationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, organizationDBTypes, false, strmangle.SetComplement(organizationPrimaryKeyColumns, organizationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Organization{&b, &c} {
		err = a.SetOrganization(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Organization != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Links[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrganizationID != x.ID {
			t.Error("foreign key was wrong value", a.OrganizationID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OrganizationID))
		reflect.Indirect(reflect.ValueOf(&a.OrganizationID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrganizationID != x.ID {
			t.Error("foreign key was wrong value", a.OrganizationID, x.ID)
		}
	}
}

func testLinksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
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

func testLinksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LinkSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLinksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Links().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	linkDBTypes = map[string]string{`ID`: `uuid`, `OrganizationID`: `uuid`, `Title`: `text`, `URL`: `text`}
	_           = bytes.MinRead
)

func testLinksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(linkPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(linkColumns) == len(linkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Links().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, linkDBTypes, true, linkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLinksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(linkColumns) == len(linkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Link{}
	if err = randomize.Struct(seed, o, linkDBTypes, true, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Links().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, linkDBTypes, true, linkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(linkColumns, linkPrimaryKeyColumns) {
		fields = linkColumns
	} else {
		fields = strmangle.SetComplement(
			linkColumns,
			linkPrimaryKeyColumns,
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

	slice := LinkSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testLinksUpsert(t *testing.T) {
	t.Parallel()

	if len(linkColumns) == len(linkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Link{}
	if err = randomize.Struct(seed, &o, linkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Link: %s", err)
	}

	count, err := Links().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, linkDBTypes, false, linkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Link: %s", err)
	}

	count, err = Links().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
