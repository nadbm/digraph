package models

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// LinkValue wraps a link with additional fields that are not obtained from the database.
type LinkValue struct {
	*Link
	NewlyAdded bool
	View       *View
}

// TopicValue wraps a topic with additional fields that are not obtained from the database.
type TopicValue struct {
	*Topic
	NewlyAdded bool
	View       *View
}

// IsNamespaceable tags Link as implementing the Namespaceable interface.
func (Link) IsNamespaceable() {}

// IsResourceIdentifiable tags Link as implementing the ResourceIdentifiable interface.
func (Link) IsResourceIdentifiable() {}

// IsSearchResultItem tags Link as being in the SearchResultItem union.
func (Link) IsSearchResultItem() {}

// Summary returns info on a user that can be printed to the log
func (l Link) Summary() string {
	return fmt.Sprintf("link %s (%s)", l.URL, l.ID)
}

// IsResourceIdentifiable tags Repository as implementing the ResourceIdentifiable interface.
func (Repository) IsResourceIdentifiable() {}

// Summary returns a string that can be included in logs.
func (r Repository) Summary() string {
	return fmt.Sprintf("repo %s (%s)", r.Name, r.ID)
}

// IsNamespaceable tags Topic as implementing the Namespaceable interface.
func (Topic) IsNamespaceable() {}

// IsResourceIdentifiable tags Topic as implementing the ResourceIdentifiable interface.
func (Topic) IsResourceIdentifiable() {}

// IsSearchResultItem tags Topic as being in the SearchResultItem union.
func (Topic) IsSearchResultItem() {}

// Summary returns info on a user that can be printed to the log
func (t Topic) Summary() string {
	return fmt.Sprintf("topic %s (%s)", t.Name, t.ID)
}

// IsResourceIdentifiable tags Organization as implementing the ResourceIdentifiable interface.
func (Organization) IsResourceIdentifiable() {}

// NewAlert returns an initialized alert.
func NewAlert(typ AlertType, text string) *Alert {
	return &Alert{
		ID:   uuid.New().String(),
		Text: text,
		Type: typ,
	}
}

// DefaultRepo returns the user's default repo.
func (u *User) DefaultRepo(ctx context.Context, exec boil.ContextExecutor) (*Repository, error) {
	// log.Printf("Looking for %s and %s (%s)", u.Login, u.ID, u.Name)
	return Repositories(
		qm.InnerJoin("organizations o on o.id = repositories.organization_id"),
		qm.Where("o.login = ? and repositories.system and repositories.owner_id = ?", u.Login, u.ID),
	).One(ctx, exec)
}

// RootTopic returns the root topic of the repository.
func (r *Repository) RootTopic(
	ctx context.Context, exec boil.ContextExecutor, view *View,
) (*TopicValue, error) {
	topic, err := r.Topics(qm.Where("root")).One(ctx, exec)
	return &TopicValue{Topic: topic, View: view}, err
}

// IsPrivate is true if the repository is a private repo.
func (r *Repository) IsPrivate() bool {
	return r.System && r.Name == "system:default"
}

// DisplayColor returns a string of the hex color to use for the topic.
func (r *Repository) DisplayColor() string {
	return "#dbedff"
}

// IsGuest returns true if the user is not backed by a row in the database.
func (u User) IsGuest() bool {
	return u.ID == ""
}

// Summary returns info on a user that can be printed to the log
func (u User) Summary() string {
	if u.Name == "" {
		return fmt.Sprintf("user no name (%s)", u.PrimaryEmail)
	}
	return fmt.Sprintf("user %s (%s)", u.Name, u.PrimaryEmail)
}

// DefaultView returns a view that can be used in return values for mutations and similar situations
func (u User) DefaultView() *View {
	return &View{ViewerID: u.ID}
}

// Filter filters a query according to the repos that have been selected and that the user can
// access.
func (v View) Filter(mods []qm.QueryMod) []qm.QueryMod {
	if v.ViewerID == "" {
		return append(mods,
			qm.InnerJoin("organizations o on o.id = r.organization_id"),
			qm.Where("r.system and o.public"),
		)
	}

	if len(v.RepositoryIds) > 0 {
		ids := v.RepositoryIdsForQuery()
		return append(mods, qm.WhereIn("r.id in ?", ids...))
	}

	return append(mods,
		qm.InnerJoin("organization_members om on r.organization_id = om.organization_id"),
		qm.WhereIn("om.user_id = ? ", v.ViewerID),
	)
}
