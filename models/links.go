// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Link is an object representing the database table.
type Link struct {
	OrganizationID string    `boil:"organization_id" json:"organization_id" toml:"organization_id" yaml:"organization_id"`
	ID             string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	URL            string    `boil:"url" json:"url" toml:"url" yaml:"url"`
	Title          string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	Sha1           string    `boil:"sha1" json:"sha1" toml:"sha1" yaml:"sha1"`
	CreatedAt      time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *linkR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L linkL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LinkColumns = struct {
	OrganizationID string
	ID             string
	URL            string
	Title          string
	Sha1           string
	CreatedAt      string
	UpdatedAt      string
}{
	OrganizationID: "organization_id",
	ID:             "id",
	URL:            "url",
	Title:          "title",
	Sha1:           "sha1",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// LinkRels is where relationship names are stored.
var LinkRels = struct {
	Organization string
	ParentTopics string
}{
	Organization: "Organization",
	ParentTopics: "ParentTopics",
}

// linkR is where relationships are stored.
type linkR struct {
	Organization *Organization
	ParentTopics TopicSlice
}

// NewStruct creates a new relationship struct
func (*linkR) NewStruct() *linkR {
	return &linkR{}
}

// linkL is where Load methods for each relationship are stored.
type linkL struct{}

var (
	linkColumns               = []string{"organization_id", "id", "url", "title", "sha1", "created_at", "updated_at"}
	linkColumnsWithoutDefault = []string{"organization_id", "url", "sha1"}
	linkColumnsWithDefault    = []string{"id", "title", "created_at", "updated_at"}
	linkPrimaryKeyColumns     = []string{"id"}
)

type (
	// LinkSlice is an alias for a slice of pointers to Link.
	// This should generally be used opposed to []Link.
	LinkSlice []*Link
	// LinkHook is the signature for custom Link hook methods
	LinkHook func(context.Context, boil.ContextExecutor, *Link) error

	linkQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	linkType                 = reflect.TypeOf(&Link{})
	linkMapping              = queries.MakeStructMapping(linkType)
	linkPrimaryKeyMapping, _ = queries.BindMapping(linkType, linkMapping, linkPrimaryKeyColumns)
	linkInsertCacheMut       sync.RWMutex
	linkInsertCache          = make(map[string]insertCache)
	linkUpdateCacheMut       sync.RWMutex
	linkUpdateCache          = make(map[string]updateCache)
	linkUpsertCacheMut       sync.RWMutex
	linkUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var linkBeforeInsertHooks []LinkHook
var linkBeforeUpdateHooks []LinkHook
var linkBeforeDeleteHooks []LinkHook
var linkBeforeUpsertHooks []LinkHook

var linkAfterInsertHooks []LinkHook
var linkAfterSelectHooks []LinkHook
var linkAfterUpdateHooks []LinkHook
var linkAfterDeleteHooks []LinkHook
var linkAfterUpsertHooks []LinkHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Link) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range linkBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Link) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range linkBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Link) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range linkBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Link) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range linkBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Link) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range linkAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Link) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range linkAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Link) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range linkAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Link) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range linkAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Link) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range linkAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLinkHook registers your hook function for all future operations.
func AddLinkHook(hookPoint boil.HookPoint, linkHook LinkHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		linkBeforeInsertHooks = append(linkBeforeInsertHooks, linkHook)
	case boil.BeforeUpdateHook:
		linkBeforeUpdateHooks = append(linkBeforeUpdateHooks, linkHook)
	case boil.BeforeDeleteHook:
		linkBeforeDeleteHooks = append(linkBeforeDeleteHooks, linkHook)
	case boil.BeforeUpsertHook:
		linkBeforeUpsertHooks = append(linkBeforeUpsertHooks, linkHook)
	case boil.AfterInsertHook:
		linkAfterInsertHooks = append(linkAfterInsertHooks, linkHook)
	case boil.AfterSelectHook:
		linkAfterSelectHooks = append(linkAfterSelectHooks, linkHook)
	case boil.AfterUpdateHook:
		linkAfterUpdateHooks = append(linkAfterUpdateHooks, linkHook)
	case boil.AfterDeleteHook:
		linkAfterDeleteHooks = append(linkAfterDeleteHooks, linkHook)
	case boil.AfterUpsertHook:
		linkAfterUpsertHooks = append(linkAfterUpsertHooks, linkHook)
	}
}

// One returns a single link record from the query.
func (q linkQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Link, error) {
	o := &Link{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for links")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Link records from the query.
func (q linkQuery) All(ctx context.Context, exec boil.ContextExecutor) (LinkSlice, error) {
	var o []*Link

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Link slice")
	}

	if len(linkAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Link records in the query.
func (q linkQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count links rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q linkQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if links exists")
	}

	return count > 0, nil
}

// Organization pointed to by the foreign key.
func (o *Link) Organization(mods ...qm.QueryMod) organizationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.OrganizationID),
	}

	queryMods = append(queryMods, mods...)

	query := Organizations(queryMods...)
	queries.SetFrom(query.Query, "\"organizations\"")

	return query
}

// ParentTopics retrieves all the topic's Topics with an executor via id column.
func (o *Link) ParentTopics(mods ...qm.QueryMod) topicQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"link_topics\" on \"topics\".\"id\" = \"link_topics\".\"parent_id\""),
		qm.Where("\"link_topics\".\"child_id\"=?", o.ID),
	)

	query := Topics(queryMods...)
	queries.SetFrom(query.Query, "\"topics\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"topics\".*"})
	}

	return query
}

// LoadOrganization allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (linkL) LoadOrganization(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLink interface{}, mods queries.Applicator) error {
	var slice []*Link
	var object *Link

	if singular {
		object = maybeLink.(*Link)
	} else {
		slice = *maybeLink.(*[]*Link)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &linkR{}
		}
		args = append(args, object.OrganizationID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &linkR{}
			}

			for _, a := range args {
				if a == obj.OrganizationID {
					continue Outer
				}
			}

			args = append(args, obj.OrganizationID)

		}
	}

	query := NewQuery(qm.From(`organizations`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Organization")
	}

	var resultSlice []*Organization
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Organization")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for organizations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for organizations")
	}

	if len(linkAfterSelectHooks) != 0 {
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
		object.R.Organization = foreign
		if foreign.R == nil {
			foreign.R = &organizationR{}
		}
		foreign.R.Links = append(foreign.R.Links, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrganizationID == foreign.ID {
				local.R.Organization = foreign
				if foreign.R == nil {
					foreign.R = &organizationR{}
				}
				foreign.R.Links = append(foreign.R.Links, local)
				break
			}
		}
	}

	return nil
}

// LoadParentTopics allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (linkL) LoadParentTopics(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLink interface{}, mods queries.Applicator) error {
	var slice []*Link
	var object *Link

	if singular {
		object = maybeLink.(*Link)
	} else {
		slice = *maybeLink.(*[]*Link)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &linkR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &linkR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(
		qm.Select("\"topics\".*, \"a\".\"child_id\""),
		qm.From("\"topics\""),
		qm.InnerJoin("\"link_topics\" as \"a\" on \"topics\".\"id\" = \"a\".\"parent_id\""),
		qm.WhereIn("\"a\".\"child_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load topics")
	}

	var resultSlice []*Topic

	var localJoinCols []string
	for results.Next() {
		one := new(Topic)
		var localJoinCol string

		err = results.Scan(&one.OrganizationID, &one.ID, &one.Name, &one.Description, &one.CreatedAt, &one.UpdatedAt, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for topics")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice topics")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on topics")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for topics")
	}

	if len(topicAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ParentTopics = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &topicR{}
			}
			foreign.R.ChildLinks = append(foreign.R.ChildLinks, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.ParentTopics = append(local.R.ParentTopics, foreign)
				if foreign.R == nil {
					foreign.R = &topicR{}
				}
				foreign.R.ChildLinks = append(foreign.R.ChildLinks, local)
				break
			}
		}
	}

	return nil
}

// SetOrganization of the link to the related item.
// Sets o.R.Organization to related.
// Adds o to related.R.Links.
func (o *Link) SetOrganization(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Organization) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"links\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"organization_id"}),
		strmangle.WhereClause("\"", "\"", 2, linkPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrganizationID = related.ID
	if o.R == nil {
		o.R = &linkR{
			Organization: related,
		}
	} else {
		o.R.Organization = related
	}

	if related.R == nil {
		related.R = &organizationR{
			Links: LinkSlice{o},
		}
	} else {
		related.R.Links = append(related.R.Links, o)
	}

	return nil
}

// AddParentTopics adds the given related objects to the existing relationships
// of the link, optionally inserting them as new records.
// Appends related to o.R.ParentTopics.
// Sets related.R.ChildLinks appropriately.
func (o *Link) AddParentTopics(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Topic) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"link_topics\" (\"child_id\", \"parent_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, query)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &linkR{
			ParentTopics: related,
		}
	} else {
		o.R.ParentTopics = append(o.R.ParentTopics, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &topicR{
				ChildLinks: LinkSlice{o},
			}
		} else {
			rel.R.ChildLinks = append(rel.R.ChildLinks, o)
		}
	}
	return nil
}

// SetParentTopics removes all previously related items of the
// link replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.ChildLinks's ParentTopics accordingly.
// Replaces o.R.ParentTopics with related.
// Sets related.R.ChildLinks's ParentTopics accordingly.
func (o *Link) SetParentTopics(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Topic) error {
	query := "delete from \"link_topics\" where \"child_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeParentTopicsFromChildLinksSlice(o, related)
	if o.R != nil {
		o.R.ParentTopics = nil
	}
	return o.AddParentTopics(ctx, exec, insert, related...)
}

// RemoveParentTopics relationships from objects passed in.
// Removes related items from R.ParentTopics (uses pointer comparison, removal does not keep order)
// Sets related.R.ChildLinks.
func (o *Link) RemoveParentTopics(ctx context.Context, exec boil.ContextExecutor, related ...*Topic) error {
	var err error
	query := fmt.Sprintf(
		"delete from \"link_topics\" where \"child_id\" = $1 and \"parent_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeParentTopicsFromChildLinksSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ParentTopics {
			if rel != ri {
				continue
			}

			ln := len(o.R.ParentTopics)
			if ln > 1 && i < ln-1 {
				o.R.ParentTopics[i] = o.R.ParentTopics[ln-1]
			}
			o.R.ParentTopics = o.R.ParentTopics[:ln-1]
			break
		}
	}

	return nil
}

func removeParentTopicsFromChildLinksSlice(o *Link, related []*Topic) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.ChildLinks {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.ChildLinks)
			if ln > 1 && i < ln-1 {
				rel.R.ChildLinks[i] = rel.R.ChildLinks[ln-1]
			}
			rel.R.ChildLinks = rel.R.ChildLinks[:ln-1]
			break
		}
	}
}

// Links retrieves all the records using an executor.
func Links(mods ...qm.QueryMod) linkQuery {
	mods = append(mods, qm.From("\"links\""))
	return linkQuery{NewQuery(mods...)}
}

// FindLink retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLink(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Link, error) {
	linkObj := &Link{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"links\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, linkObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from links")
	}

	return linkObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Link) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no links provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(linkColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	linkInsertCacheMut.RLock()
	cache, cached := linkInsertCache[key]
	linkInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			linkColumns,
			linkColumnsWithDefault,
			linkColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(linkType, linkMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(linkType, linkMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"links\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"links\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into links")
	}

	if !cached {
		linkInsertCacheMut.Lock()
		linkInsertCache[key] = cache
		linkInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Link.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Link) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	linkUpdateCacheMut.RLock()
	cache, cached := linkUpdateCache[key]
	linkUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			linkColumns,
			linkPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update links, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"links\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, linkPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(linkType, linkMapping, append(wl, linkPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update links row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for links")
	}

	if !cached {
		linkUpdateCacheMut.Lock()
		linkUpdateCache[key] = cache
		linkUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q linkQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for links")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for links")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LinkSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), linkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"links\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, linkPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in link slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all link")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Link) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no links provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(linkColumnsWithDefault, o)

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

	linkUpsertCacheMut.RLock()
	cache, cached := linkUpsertCache[key]
	linkUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			linkColumns,
			linkColumnsWithDefault,
			linkColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			linkColumns,
			linkPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert links, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(linkPrimaryKeyColumns))
			copy(conflict, linkPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"links\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(linkType, linkMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(linkType, linkMapping, ret)
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
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert links")
	}

	if !cached {
		linkUpsertCacheMut.Lock()
		linkUpsertCache[key] = cache
		linkUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Link record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Link) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Link provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), linkPrimaryKeyMapping)
	sql := "DELETE FROM \"links\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from links")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for links")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q linkQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no linkQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from links")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for links")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LinkSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Link slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(linkBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), linkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"links\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, linkPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from link slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for links")
	}

	if len(linkAfterDeleteHooks) != 0 {
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
func (o *Link) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLink(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LinkSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LinkSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), linkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"links\".* FROM \"links\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, linkPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LinkSlice")
	}

	*o = slice

	return nil
}

// LinkExists checks if the Link row exists.
func LinkExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"links\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if links exists")
	}

	return exists, nil
}
