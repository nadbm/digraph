package resolvers_test

import (
	"context"
	"testing"

	"github.com/emwalker/digraph/resolvers"
)

func TestResolveView(t *testing.T) {
	ctx := context.Background()
	m := newMutator(t, testActor)

	repo := m.defaultRepo()
	org, err := repo.Organization().One(ctx, testDB)
	if err != nil {
		t.Fatal(err)
	}

	queryResolver := (&resolvers.Resolver{DB: testDB}).Query()

	cases := []struct {
		Name     string
		RepoID   *string
		RepoName *string
		OrgLogin string
	}{
		{
			Name:     "When the org login and repo name are provied",
			RepoName: &repo.Name,
			OrgLogin: org.Login,
		},
		{
			Name:     "When only the org login is provided",
			RepoName: nil,
			OrgLogin: org.Login,
		},
	}

	for _, td := range cases {
		t.Run(td.Name, func(t *testing.T) {
			view, err := queryResolver.View(
				ctx,
				td.OrgLogin,
				td.RepoName,
				[]string{},
				&testActor.ID,
			)

			if err != nil {
				t.Fatal(err)
			}

			if repo.ID != view.CurrentRepository.ID {
				t.Fatalf("Expected repo %s, got repo %s", repo.ID, view.CurrentRepository.ID)
			}
		})
	}
}
