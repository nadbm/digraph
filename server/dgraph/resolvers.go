package dgraph

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type ConnectionArgs struct {
	First  *int32
	After  *string
	Last   *int32
	Before *string
}

type Connection struct {
	ids  []int
	from int
	to   int
}

type PersistentRoot struct{}

func NewRoot() (*PersistentRoot, error) {
	return &PersistentRoot{}, nil
}

func (r *PersistentRoot) Viewer(ctx context.Context) (*QueryBuilder, error) {
	builder := ctx.Value("builder").(*QueryBuilder)
	builder.Fragment(
		"ViewerFragment",
		"viewerScalar(func: uid(0x5d))",
	)
	return builder, nil
}

func (r *PersistentRoot) Organization(
	ctx context.Context,
	args struct{ ExternalId string },
) (*QueryBuilder, error) {
	builder := ctx.Value("builder").(*QueryBuilder)
	builder.Fragment(
		"OrganizationFragment",
		`organizationScalar(func: has(externalId)) @filter(eq(externalId, "tyrell"))`,
	)
	return builder, nil
}

type CreateTopicInput struct {
	ClientMutationId string
	Description      string
	Name             string
	OrganizationId   string
	TopicIds         []string
}

func (r *PersistentRoot) CreateTopic(
	ctx context.Context,
	input CreateTopicInput,
) (*QueryBuilder, error) {
	builder := ctx.Value("builder").(*QueryBuilder)
	builder.Fragment(
		"createTopicFragment",
		`some-mutation-stuff`,
	)
	return builder, nil
}

func (b *QueryBuilder) ChildTopics(args ConnectionArgs) (*QueryBuilder, error) {
	builder := b.Push().Fragment(
		"ChildTopicsFragment",
		`childTopicConnection: includes @filter(eq(type, "topic"))`,
	)
	return builder, nil
}

func (b *QueryBuilder) Cursor() (string, error) {
	return "", nil
}

func (b *QueryBuilder) Description() (*string, error) {
	b.Field("description")
	return nil, nil
}

func (b *QueryBuilder) Edges() (*[]*QueryBuilder, error) {
	return &[]*QueryBuilder{
		b,
	}, nil
}

func (b *QueryBuilder) Email() (string, error) {
	b.Field("email")
	return "", nil
}

func (b *QueryBuilder) EndCursor() (*string, error) {
	var str string = ""
	return &str, nil
}

func (b *QueryBuilder) ExternalId() (string, error) {
	b.Field("externalId")
	return "", nil
}

func (b *QueryBuilder) HasNextPage() (bool, error) {
	return false, nil
}

func (b *QueryBuilder) HasPreviousPage() (bool, error) {
	return false, nil
}

func (b *QueryBuilder) ID() *graphql.ID {
	idString := graphql.ID("1")
	return &idString
}

func (b *QueryBuilder) Links(args ConnectionArgs) (*QueryBuilder, error) {
	builder := b.Push().Fragment(
		"LinksFragment",
		`linkConnection: owns @filter(eq(type, "link"))`,
	)
	return builder, nil
}

func (b *QueryBuilder) Name() (string, error) {
	b.Field("name")
	return "", nil
}

func (b *QueryBuilder) Node() (*QueryBuilder, error) {
	return b, nil
}

func (b *QueryBuilder) PageInfo() (*QueryBuilder, error) {
	return b, nil
}

func (b *QueryBuilder) ParentTopics(args ConnectionArgs) (*QueryBuilder, error) {
	builder := b.Push().Fragment(
		"ParentTopicsFragment",
		`parentTopicConnection: ~includes @filter(eq(type, "topic"))`,
	)
	return builder, nil
}

func (b *QueryBuilder) StartCursor() (*string, error) {
	var str string = ""
	return &str, nil
}

func (b *QueryBuilder) Title() (string, error) {
	b.Field("title")
	return "", nil
}

func (b *QueryBuilder) Topic(args struct{ ExternalId string }) (*QueryBuilder, error) {
	builder := b.Push().Fragment(
		"TopicFragment",
		`topicScalar: owns @filter(eq(type, "topic") AND eq(externalId, "github"))`,
	)
	return builder, nil
}

func (b *QueryBuilder) Topics(args ConnectionArgs) (*QueryBuilder, error) {
	builder := b.Push().Fragment(
		"TopicsFragment",
		`topicConnection: owns @filter(eq(type, "topic"))`,
	)
	return builder, nil
}

func (b *QueryBuilder) Uid() (string, error) {
	b.Field("uid")
	return "", nil
}

func (b *QueryBuilder) Url() (string, error) {
	b.Field("url")
	return "", nil
}

func (b *QueryBuilder) User() (*QueryBuilder, error) {
	return b, nil
}
