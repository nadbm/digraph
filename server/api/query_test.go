package api

import (
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
)

func init() {
	Tests = []T{
		{
			Query: `
		      query {
		        viewer {
		          name
		          email
		        }
		      }
		    `,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"viewer": map[string]interface{}{
						"name":  "Gnusto",
						"email": "gnusto@tyrell.test",
					},
				},
			},
		},
		{
			Query: `
				query {
					organization(resourceId: "organization:tyrell") {
						name

						topics(first: 100) {
							edges {
								node {
									name
								}
							}
						}
					}
				}
			`,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"organization": map[string]interface{}{
						"name": "Tyrell Corporation",
						"topics": map[string]interface{}{
							"edges": []interface{}{
								map[string]interface{}{
									"node": map[string]interface{}{
										"name": "Biology",
									},
								},
								map[string]interface{}{
									"node": map[string]interface{}{
										"name": "Chemistry",
									},
								},
								map[string]interface{}{
									"node": map[string]interface{}{
										"name": "Science",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Query: `
				query {
					topic(resourceId: "topic:science") {
						name
						description
					}
				}
			`,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"topic": map[string]interface{}{
						"name":        "Science",
						"description": nil,
					},
				},
			},
		},
	}
}

func TestQuery(t *testing.T) {
	conn := NewConnection("memstore", "")
	app, _ := New(conn, echo.New())
	conn.(*CayleyConnection).makeTestStore(simpleGraph)
	defer checkErr(conn.Close())

	for _, test := range Tests {
		params := graphql.Params{
			Schema:        *app.Schema,
			RequestString: test.Query,
		}
		testGraphql(test, params, t)
	}
}
