package resolvers

import (
	"context"
	"fmt"

	"github.com/emwalker/digraph/internal/loaders"
	"github.com/emwalker/digraph/internal/models"
)

type repositoryResolver struct {
	*Resolver
}

func getRepositoryLoader(ctx context.Context) *loaders.RepositoryLoader {
	return ctx.Value(loaders.RepositoryLoaderKey).(*loaders.RepositoryLoader)
}

func fetchRepository(ctx context.Context, repoID string) (models.Repository, error) {
	loader := getRepositoryLoader(ctx)
	repo, err := loader.Load(repoID)
	if err != nil {
		return models.Repository{}, err
	}
	return *repo, nil
}

// DisplayColor returns a color by which to display the topic.
func (r *repositoryResolver) DisplayColor(ctx context.Context, repo *models.Repository) (string, error) {
	color := repo.DisplayColor()
	return color, nil
}

func (r *repositoryResolver) DisplayName(ctx context.Context, repo *models.Repository) (string, error) {
	if repo.IsPrivate() {
		return "Private collection", nil
	}
	return repo.Name, nil
}

// FullName returns a path-like name that can be used in lists and select options.
func (r *repositoryResolver) FullName(
	ctx context.Context, repo *models.Repository,
) (string, error) {
	var org models.Organization
	var err error

	if org, err = fetchOrganization(ctx, repo.OrganizationID); err != nil {
		return "", err
	}

	var name string
	if repo.IsPrivate() {
		name = "private"
	} else if repo.System {
		name = "general"
	} else {
		name = repo.Name
	}

	return fmt.Sprintf("%s/%s", org.Login, name), nil
}

// isPrivate indicates whether the repository is private or not.
func (r *repositoryResolver) IsPrivate(
	ctx context.Context, repo *models.Repository,
) (bool, error) {
	return repo.IsPrivate(), nil
}

// Organization returns a set of links.
func (r *repositoryResolver) Organization(
	ctx context.Context, repo *models.Repository,
) (models.Organization, error) {
	return fetchOrganization(ctx, repo.OrganizationID)
}

// Organization returns a set of links.
func (r *repositoryResolver) Owner(
	ctx context.Context, repo *models.Repository,
) (models.User, error) {
	owner, err := repo.Owner().One(ctx, r.DB)
	return *owner, err
}

// RootTopic returns the root topic of the repository.
func (r *repositoryResolver) RootTopic(
	ctx context.Context, repo *models.Repository,
) (models.TopicValue, error) {
	topic, err := repo.RootTopic(ctx, r.DB, r.Actor.DefaultView())
	if err != nil {
		return models.TopicValue{}, err
	}
	return *topic, err
}
