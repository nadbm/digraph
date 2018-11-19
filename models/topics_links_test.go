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

func testTopicsLinks(t *testing.T) {
	t.Parallel()

	query := TopicsLinks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTopicsLinksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
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

	count, err := TopicsLinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTopicsLinksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TopicsLinks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TopicsLinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTopicsLinksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TopicsLinkSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TopicsLinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTopicsLinksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TopicsLinkExists(ctx, tx, o.OrganizationID, o.ParentID, o.ChildID)
	if err != nil {
		t.Errorf("Unable to check if TopicsLink exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TopicsLinkExists to return true, but got false.")
	}
}

func testTopicsLinksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	topicsLinkFound, err := FindTopicsLink(ctx, tx, o.OrganizationID, o.ParentID, o.ChildID)
	if err != nil {
		t.Error(err)
	}

	if topicsLinkFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTopicsLinksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TopicsLinks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTopicsLinksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TopicsLinks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTopicsLinksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	topicsLinkOne := &TopicsLink{}
	topicsLinkTwo := &TopicsLink{}
	if err = randomize.Struct(seed, topicsLinkOne, topicsLinkDBTypes, false, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}
	if err = randomize.Struct(seed, topicsLinkTwo, topicsLinkDBTypes, false, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = topicsLinkOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = topicsLinkTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TopicsLinks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTopicsLinksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	topicsLinkOne := &TopicsLink{}
	topicsLinkTwo := &TopicsLink{}
	if err = randomize.Struct(seed, topicsLinkOne, topicsLinkDBTypes, false, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}
	if err = randomize.Struct(seed, topicsLinkTwo, topicsLinkDBTypes, false, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = topicsLinkOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = topicsLinkTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TopicsLinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func topicsLinkBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *TopicsLink) error {
	*o = TopicsLink{}
	return nil
}

func topicsLinkAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *TopicsLink) error {
	*o = TopicsLink{}
	return nil
}

func topicsLinkAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *TopicsLink) error {
	*o = TopicsLink{}
	return nil
}

func topicsLinkBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TopicsLink) error {
	*o = TopicsLink{}
	return nil
}

func topicsLinkAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TopicsLink) error {
	*o = TopicsLink{}
	return nil
}

func topicsLinkBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TopicsLink) error {
	*o = TopicsLink{}
	return nil
}

func topicsLinkAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TopicsLink) error {
	*o = TopicsLink{}
	return nil
}

func topicsLinkBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TopicsLink) error {
	*o = TopicsLink{}
	return nil
}

func topicsLinkAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TopicsLink) error {
	*o = TopicsLink{}
	return nil
}

func testTopicsLinksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &TopicsLink{}
	o := &TopicsLink{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TopicsLink object: %s", err)
	}

	AddTopicsLinkHook(boil.BeforeInsertHook, topicsLinkBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	topicsLinkBeforeInsertHooks = []TopicsLinkHook{}

	AddTopicsLinkHook(boil.AfterInsertHook, topicsLinkAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	topicsLinkAfterInsertHooks = []TopicsLinkHook{}

	AddTopicsLinkHook(boil.AfterSelectHook, topicsLinkAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	topicsLinkAfterSelectHooks = []TopicsLinkHook{}

	AddTopicsLinkHook(boil.BeforeUpdateHook, topicsLinkBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	topicsLinkBeforeUpdateHooks = []TopicsLinkHook{}

	AddTopicsLinkHook(boil.AfterUpdateHook, topicsLinkAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	topicsLinkAfterUpdateHooks = []TopicsLinkHook{}

	AddTopicsLinkHook(boil.BeforeDeleteHook, topicsLinkBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	topicsLinkBeforeDeleteHooks = []TopicsLinkHook{}

	AddTopicsLinkHook(boil.AfterDeleteHook, topicsLinkAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	topicsLinkAfterDeleteHooks = []TopicsLinkHook{}

	AddTopicsLinkHook(boil.BeforeUpsertHook, topicsLinkBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	topicsLinkBeforeUpsertHooks = []TopicsLinkHook{}

	AddTopicsLinkHook(boil.AfterUpsertHook, topicsLinkAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	topicsLinkAfterUpsertHooks = []TopicsLinkHook{}
}

func testTopicsLinksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TopicsLinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTopicsLinksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(topicsLinkColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TopicsLinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTopicsLinkToOneLinkUsingChild(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local TopicsLink
	var foreign Link

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, topicsLinkDBTypes, false, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, linkDBTypes, false, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ChildID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Child().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TopicsLinkSlice{&local}
	if err = local.L.LoadChild(ctx, tx, false, (*[]*TopicsLink)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Child == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Child = nil
	if err = local.L.LoadChild(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Child == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTopicsLinkToOneOrganizationUsingOrganization(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local TopicsLink
	var foreign Organization

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, topicsLinkDBTypes, false, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
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

	slice := TopicsLinkSlice{&local}
	if err = local.L.LoadOrganization(ctx, tx, false, (*[]*TopicsLink)(&slice), nil); err != nil {
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

func testTopicsLinkToOneTopicUsingParent(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local TopicsLink
	var foreign Topic

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, topicsLinkDBTypes, false, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, topicDBTypes, false, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ParentID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Parent().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TopicsLinkSlice{&local}
	if err = local.L.LoadParent(ctx, tx, false, (*[]*TopicsLink)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Parent == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Parent = nil
	if err = local.L.LoadParent(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Parent == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTopicsLinkToOneSetOpLinkUsingChild(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a TopicsLink
	var b, c Link

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, topicsLinkDBTypes, false, strmangle.SetComplement(topicsLinkPrimaryKeyColumns, topicsLinkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, linkDBTypes, false, strmangle.SetComplement(linkPrimaryKeyColumns, linkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, linkDBTypes, false, strmangle.SetComplement(linkPrimaryKeyColumns, linkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Link{&b, &c} {
		err = a.SetChild(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Child != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ChildTopicsLinks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ChildID != x.ID {
			t.Error("foreign key was wrong value", a.ChildID)
		}

		if exists, err := TopicsLinkExists(ctx, tx, a.OrganizationID, a.ParentID, a.ChildID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testTopicsLinkToOneSetOpOrganizationUsingOrganization(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a TopicsLink
	var b, c Organization

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, topicsLinkDBTypes, false, strmangle.SetComplement(topicsLinkPrimaryKeyColumns, topicsLinkColumnsWithoutDefault)...); err != nil {
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

		if x.R.TopicsLinks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrganizationID != x.ID {
			t.Error("foreign key was wrong value", a.OrganizationID)
		}

		if exists, err := TopicsLinkExists(ctx, tx, a.OrganizationID, a.ParentID, a.ChildID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testTopicsLinkToOneSetOpTopicUsingParent(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a TopicsLink
	var b, c Topic

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, topicsLinkDBTypes, false, strmangle.SetComplement(topicsLinkPrimaryKeyColumns, topicsLinkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, topicDBTypes, false, strmangle.SetComplement(topicPrimaryKeyColumns, topicColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, topicDBTypes, false, strmangle.SetComplement(topicPrimaryKeyColumns, topicColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Topic{&b, &c} {
		err = a.SetParent(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Parent != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ParentTopicsLinks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ParentID != x.ID {
			t.Error("foreign key was wrong value", a.ParentID)
		}

		if exists, err := TopicsLinkExists(ctx, tx, a.OrganizationID, a.ParentID, a.ChildID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testTopicsLinksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
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

func testTopicsLinksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TopicsLinkSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTopicsLinksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TopicsLinks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	topicsLinkDBTypes = map[string]string{`ChildID`: `uuid`, `OrganizationID`: `uuid`, `ParentID`: `uuid`}
	_                 = bytes.MinRead
)

func testTopicsLinksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(topicsLinkPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(topicsLinkColumns) == len(topicsLinkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TopicsLinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTopicsLinksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(topicsLinkColumns) == len(topicsLinkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TopicsLink{}
	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TopicsLinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, topicsLinkDBTypes, true, topicsLinkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(topicsLinkColumns, topicsLinkPrimaryKeyColumns) {
		fields = topicsLinkColumns
	} else {
		fields = strmangle.SetComplement(
			topicsLinkColumns,
			topicsLinkPrimaryKeyColumns,
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

	slice := TopicsLinkSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTopicsLinksUpsert(t *testing.T) {
	t.Parallel()

	if len(topicsLinkColumns) == len(topicsLinkPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TopicsLink{}
	if err = randomize.Struct(seed, &o, topicsLinkDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TopicsLink: %s", err)
	}

	count, err := TopicsLinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, topicsLinkDBTypes, false, topicsLinkPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TopicsLink struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TopicsLink: %s", err)
	}

	count, err = TopicsLinks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
