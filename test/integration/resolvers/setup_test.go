package resolvers_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/emwalker/digraph/internal/loaders"
	"github.com/emwalker/digraph/internal/models"
	"github.com/emwalker/digraph/internal/resolvers"
	"github.com/emwalker/digraph/internal/services"
	"github.com/emwalker/digraph/internal/services/pageinfo"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

const orgID = "45dc89a6-e6f0-11e8-8bc1-6f4d565e3ddb"

var (
	testDB     *sql.DB
	testActor  *models.User
	testActor2 *models.User
)

type testFetcher struct{}

type mutator struct {
	actor    *models.User
	ctx      context.Context
	db       *sql.DB
	resolver models.MutationResolver
	t        *testing.T
}

func (m mutator) repo(login string) *models.Repository {
	repo, err := m.actor.OwnerRepositories(
		qm.InnerJoin("organizations o on o.id = repositories.organization_id"),
		qm.Where("repositories.system and o.login = ?", login),
	).One(m.ctx, testDB)
	if err != nil {
		panic(err)
	}
	return repo
}

func (m mutator) defaultRepo() *models.Repository {
	return m.repo(m.actor.Login)
}

func TestMain(m *testing.M) {
	services.Fetcher = &testFetcher{}
	testDB = newTestDb()
	defer testDB.Close()

	var err error

	testActor, err = models.Users(
		qm.Load("SelectedRepository"),
		qm.Where("users.selected_repository_id is not null"),
	).One(context.Background(), testDB)
	if err != nil {
		panic(err)
	}

	testActor2, err = models.Users(
		qm.Where("users.id != ?", testActor.ID),
	).One(context.Background(), testDB)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func newTestDb() *sql.DB {
	var err error
	if testDB, err = sql.Open("postgres", "dbname=digraph_dev user=postgres sslmode=disable"); err != nil {
		log.Fatal("Unable to connect to the database", err)
	}
	return testDB
}

func newMutator(t *testing.T, actor *models.User) mutator {
	resolver := &resolvers.MutationResolver{
		&resolvers.Resolver{DB: testDB, Actor: actor},
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, resolvers.CurrentUserKey, actor)
	ctx = loaders.AddToContext(ctx, testDB, 1*time.Millisecond)

	return mutator{
		actor:    actor,
		ctx:      ctx,
		db:       testDB,
		resolver: resolver,
		t:        t,
	}
}

func (f *testFetcher) FetchPage(url string) (*pageinfo.PageInfo, error) {
	title := "Gnusto's blog"
	return &pageinfo.PageInfo{
		URL:   url,
		Title: &title,
	}, nil
}

func (m mutator) addParentTopicToTopic(child, parent *models.TopicValue) {
	everything, err := models.Topics(qm.Where("name like 'Everything'")).One(context.Background(), testDB)
	if err != nil {
		m.t.Fatal(err)
	}

	input := models.UpdateTopicParentTopicsInput{
		TopicID:        child.ID,
		ParentTopicIds: []string{everything.ID, parent.ID},
	}

	if _, err := m.resolver.UpdateTopicParentTopics(m.ctx, input); err != nil {
		m.t.Fatal(err)
	}
}

func (m mutator) addParentTopicToLink(link *models.LinkValue, topic *models.TopicValue) {
	input := models.UpdateLinkTopicsInput{
		LinkID:         link.ID,
		ParentTopicIds: []string{topic.ID},
	}

	if _, err := m.resolver.UpdateLinkTopics(m.ctx, input); err != nil {
		m.t.Fatal(err)
	}
}

func (m mutator) deleteTopic(topic models.TopicValue) {
	count, err := topic.Delete(m.ctx, m.db)
	if err != nil {
		m.t.Fatal(err)
	}

	if count != int64(1) {
		m.t.Fatal("Expected a single row to be deleted")
	}
}

func (m mutator) createTopic(orgLogin, repoName, name string) (*models.TopicValue, services.CleanupFunc) {
	parentTopic, err := models.Topics(qm.Where("name like 'Everything'")).One(m.ctx, m.db)
	if err != nil {
		m.t.Fatal(err)
	}

	input := models.UpsertTopicInput{
		Name:              name,
		OrganizationLogin: orgLogin,
		RepositoryName:    repoName,
		TopicIds:          []string{parentTopic.ID},
	}

	payload, err := m.resolver.UpsertTopic(m.ctx, input)
	if err != nil {
		m.t.Fatal(err)
	}

	topic := payload.TopicEdge.Node

	cleanup := func() error {
		m.deleteTopic(topic)
		return nil
	}

	return &topic, cleanup
}

func (m mutator) createLink(orgLogin, repoName, title, url string) (*models.LinkValue, services.CleanupFunc) {
	payload1, err := m.resolver.UpsertLink(m.ctx, models.UpsertLinkInput{
		AddParentTopicIds: []string{},
		OrganizationLogin: orgLogin,
		RepositoryName:    repoName,
		Title:             &title,
		URL:               url,
	})
	if err != nil {
		m.t.Fatal(err)
	}

	link := payload1.LinkEdge.Node

	cleanup := func() error {
		_, err := models.UserLinks(qm.Where("link_id = ?", link.ID)).DeleteAll(m.ctx, testDB)
		if err != nil {
			m.t.Fatal(err)
		}

		count, err := link.Delete(m.ctx, testDB)
		if err != nil {
			m.t.Fatal(err)
		}

		if count != int64(1) {
			log.Printf("Expected at least one updated record, but none was updated")
		}
		return nil
	}

	return &link, cleanup
}
