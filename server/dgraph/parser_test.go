package dgraph

import (
	"context"
	"testing"

	graphql "github.com/graph-gophers/graphql-go"
)

var schema *graphql.Schema

func init() {
	schema = LoadSchema("../../schema.graphql")
}

type parserTestCase struct {
	description string
	input       string
	expected    string
}

func runTests(t *testing.T, tables *[]parserTestCase) {
	for _, table := range *tables {
		builder := ParseQuery(context.Background(), schema, QueryParams{Query: table.input})
		actual := builder.Stringify()
		if squash(actual) != squash(table.expected) {
			t.Errorf(`
Output did not match expected output for test: %s:

%s
			`, table.description, actual)
		}
	}
}

func TestLinkQueries(t *testing.T) {
	tables := []parserTestCase{
		{
			"an organization",
			`{
				organization(externalId: "3") {
					name
				}
			}`,
			`{
				organizationScalar(func: has(externalId)) @filter(eq(externalId, "tyrell")) {
					name
				}
			}`,
		},
		{
			"links in an organization",
			`{
				organization(externalId: "3") {
					links {
						edges {
							node {
								title
								url
							}
						}
					}
				}
			}`,
			`{
				organizationScalar(func: has(externalId)) @filter(eq(externalId, "tyrell")) {
					linkConnection: owns @filter(eq(type, "link")) {
						url
						title
					}
				}
			}`,
		},
		{
			`links in an organization with a "first:" arg`,
			`{
				organization(externalId: "3") {
					links(first: 100) {
						edges {
							node {
								title
								url
							}
						}
					}
				}
			}`,
			`{
				organizationScalar(func: has(externalId)) @filter(eq(externalId, "tyrell")) {
					linkConnection: owns @filter(eq(type, "link")) {
						url
						title
					}
				}
			}`,
		},
	}

	runTests(t, &tables)
}

func TestTopicQueries(t *testing.T) {
	tables := []parserTestCase{
		{
			"topics in an organization",
			`{
				organization(externalId: "3") {
					topics {
						edges {
							node {
								name
							}
						}
					}
				}
			}`,
			`{
				organizationScalar(func: has(externalId)) @filter(eq(externalId, "tyrell")) {
					topicConnection: owns @filter(eq(type, "topic")) {
						name
					}
				}
			}`,
		},
		{
			"subtopics of a topic",
			`{
				organization(externalId: "3") {
					topic(externalId: "4") {
						childTopics(first: 100) {
							edges {
								node {
									name
									uid
								}
							}
						}
					}
				}
			}`,
			`{
				organizationScalar(func: has(externalId)) @filter(eq(externalId, "tyrell")) {
					topicScalar: owns @filter(eq(type, "topic") AND eq(externalId, "github")) {
						childTopicConnection: includes @filter(eq(type, "topic")) {
							uid
							name
						}
					}
				}
			}`,
		},
	}

	runTests(t, &tables)
}

func TestGraphqlParsing(t *testing.T) {
	tables := []parserTestCase{
		{
			"information about the current viewer",
			`query router_Query {
				viewer {
					name
				}
			}`,
			`{
				viewerScalar(func: uid(0x5d)) {
					name
				}
			}`,
		},
	}

	runTests(t, &tables)
}
