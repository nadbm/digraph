// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Links", testLinks)
	t.Run("Organizations", testOrganizations)
	t.Run("SchemaMigrations", testSchemaMigrations)
	t.Run("Sessions", testSessions)
	t.Run("Topics", testTopics)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Links", testLinksDelete)
	t.Run("Organizations", testOrganizationsDelete)
	t.Run("SchemaMigrations", testSchemaMigrationsDelete)
	t.Run("Sessions", testSessionsDelete)
	t.Run("Topics", testTopicsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Links", testLinksQueryDeleteAll)
	t.Run("Organizations", testOrganizationsQueryDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsQueryDeleteAll)
	t.Run("Sessions", testSessionsQueryDeleteAll)
	t.Run("Topics", testTopicsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Links", testLinksSliceDeleteAll)
	t.Run("Organizations", testOrganizationsSliceDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceDeleteAll)
	t.Run("Sessions", testSessionsSliceDeleteAll)
	t.Run("Topics", testTopicsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Links", testLinksExists)
	t.Run("Organizations", testOrganizationsExists)
	t.Run("SchemaMigrations", testSchemaMigrationsExists)
	t.Run("Sessions", testSessionsExists)
	t.Run("Topics", testTopicsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Links", testLinksFind)
	t.Run("Organizations", testOrganizationsFind)
	t.Run("SchemaMigrations", testSchemaMigrationsFind)
	t.Run("Sessions", testSessionsFind)
	t.Run("Topics", testTopicsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Links", testLinksBind)
	t.Run("Organizations", testOrganizationsBind)
	t.Run("SchemaMigrations", testSchemaMigrationsBind)
	t.Run("Sessions", testSessionsBind)
	t.Run("Topics", testTopicsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Links", testLinksOne)
	t.Run("Organizations", testOrganizationsOne)
	t.Run("SchemaMigrations", testSchemaMigrationsOne)
	t.Run("Sessions", testSessionsOne)
	t.Run("Topics", testTopicsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Links", testLinksAll)
	t.Run("Organizations", testOrganizationsAll)
	t.Run("SchemaMigrations", testSchemaMigrationsAll)
	t.Run("Sessions", testSessionsAll)
	t.Run("Topics", testTopicsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Links", testLinksCount)
	t.Run("Organizations", testOrganizationsCount)
	t.Run("SchemaMigrations", testSchemaMigrationsCount)
	t.Run("Sessions", testSessionsCount)
	t.Run("Topics", testTopicsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Links", testLinksHooks)
	t.Run("Organizations", testOrganizationsHooks)
	t.Run("SchemaMigrations", testSchemaMigrationsHooks)
	t.Run("Sessions", testSessionsHooks)
	t.Run("Topics", testTopicsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Links", testLinksInsert)
	t.Run("Links", testLinksInsertWhitelist)
	t.Run("Organizations", testOrganizationsInsert)
	t.Run("Organizations", testOrganizationsInsertWhitelist)
	t.Run("SchemaMigrations", testSchemaMigrationsInsert)
	t.Run("SchemaMigrations", testSchemaMigrationsInsertWhitelist)
	t.Run("Sessions", testSessionsInsert)
	t.Run("Sessions", testSessionsInsertWhitelist)
	t.Run("Topics", testTopicsInsert)
	t.Run("Topics", testTopicsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("LinkToOrganizationUsingOrganization", testLinkToOneOrganizationUsingOrganization)
	t.Run("SessionToUserUsingUser", testSessionToOneUserUsingUser)
	t.Run("TopicToOrganizationUsingOrganization", testTopicToOneOrganizationUsingOrganization)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("LinkToParentTopics", testLinkToManyParentTopics)
	t.Run("OrganizationToLinks", testOrganizationToManyLinks)
	t.Run("OrganizationToTopics", testOrganizationToManyTopics)
	t.Run("TopicToChildLinks", testTopicToManyChildLinks)
	t.Run("TopicToParentTopics", testTopicToManyParentTopics)
	t.Run("TopicToChildTopics", testTopicToManyChildTopics)
	t.Run("UserToSessions", testUserToManySessions)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("LinkToOrganizationUsingLinks", testLinkToOneSetOpOrganizationUsingOrganization)
	t.Run("SessionToUserUsingSessions", testSessionToOneSetOpUserUsingUser)
	t.Run("TopicToOrganizationUsingTopics", testTopicToOneSetOpOrganizationUsingOrganization)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("LinkToParentTopics", testLinkToManyAddOpParentTopics)
	t.Run("OrganizationToLinks", testOrganizationToManyAddOpLinks)
	t.Run("OrganizationToTopics", testOrganizationToManyAddOpTopics)
	t.Run("TopicToChildLinks", testTopicToManyAddOpChildLinks)
	t.Run("TopicToParentTopics", testTopicToManyAddOpParentTopics)
	t.Run("TopicToChildTopics", testTopicToManyAddOpChildTopics)
	t.Run("UserToSessions", testUserToManyAddOpSessions)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("LinkToParentTopics", testLinkToManySetOpParentTopics)
	t.Run("TopicToChildLinks", testTopicToManySetOpChildLinks)
	t.Run("TopicToParentTopics", testTopicToManySetOpParentTopics)
	t.Run("TopicToChildTopics", testTopicToManySetOpChildTopics)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("LinkToParentTopics", testLinkToManyRemoveOpParentTopics)
	t.Run("TopicToChildLinks", testTopicToManyRemoveOpChildLinks)
	t.Run("TopicToParentTopics", testTopicToManyRemoveOpParentTopics)
	t.Run("TopicToChildTopics", testTopicToManyRemoveOpChildTopics)
}

func TestReload(t *testing.T) {
	t.Run("Links", testLinksReload)
	t.Run("Organizations", testOrganizationsReload)
	t.Run("SchemaMigrations", testSchemaMigrationsReload)
	t.Run("Sessions", testSessionsReload)
	t.Run("Topics", testTopicsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Links", testLinksReloadAll)
	t.Run("Organizations", testOrganizationsReloadAll)
	t.Run("SchemaMigrations", testSchemaMigrationsReloadAll)
	t.Run("Sessions", testSessionsReloadAll)
	t.Run("Topics", testTopicsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Links", testLinksSelect)
	t.Run("Organizations", testOrganizationsSelect)
	t.Run("SchemaMigrations", testSchemaMigrationsSelect)
	t.Run("Sessions", testSessionsSelect)
	t.Run("Topics", testTopicsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Links", testLinksUpdate)
	t.Run("Organizations", testOrganizationsUpdate)
	t.Run("SchemaMigrations", testSchemaMigrationsUpdate)
	t.Run("Sessions", testSessionsUpdate)
	t.Run("Topics", testTopicsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Links", testLinksSliceUpdateAll)
	t.Run("Organizations", testOrganizationsSliceUpdateAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceUpdateAll)
	t.Run("Sessions", testSessionsSliceUpdateAll)
	t.Run("Topics", testTopicsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
