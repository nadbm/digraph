package dgraph

import (
	"bytes"
	realjson "encoding/json"
	"log"
	"testing"
)

type serializerTestCase struct {
	description string
	input       string
	expected    string
}

func pretty(input []byte) string {
	var buf bytes.Buffer
	err := realjson.Indent(&buf, input, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func runSerializerTests(t *testing.T, tables *[]serializerTestCase) {
	for _, table := range *tables {
		actual := Serialize([]byte(table.input))
		formatted := pretty(actual)
		if squash(formatted) != squash(table.expected) {
			t.Errorf(`
				Json did not transform as expected: %s:
				%s
			`, table.description, formatted)
		}
	}
}

func TestSerialization(t *testing.T) {
	tables := []serializerTestCase{
		{
			"simple case",
			`{
				"organizationScalar": [
					{
						"name": "Tyrell Corporation",
						"externalId": "tyrell"
					}
				]
			}`,
			`{
				"data": {
					"organization": {
						"externalId": "tyrell",
						"name": "Tyrell Corporation"
					}
				}
			}`,
		},
		{
			"case with links",
			`{
				"organizationScalar": [
					{
						"name": "Tyrell Corporation",
						"externalId": "tyrell",
						"linkConnection": [
							{
								"title": "New York Times",
								"url": "https://www.nytimes.com"
							}
						]
					}
				]
			}`,
			` {
				"data": {
					"organization": {
						"links": {
							"edges": [
								{
									"node": {
										"title": "New York Times",
										"url": "https://www.nytimes.com"
									}
								}
							]
						},
						"externalId": "tyrell",
						"name": "Tyrell Corporation"
					}
				}
			}`,
		},
		{
			"case with topics",
			`{
				"organizationScalar": [
					{
						"topicConnection": [
							{
								"name": "New York Times"
							}
						]
					}
				]
			}`,
			` {
				"data": {
					"organization": {
						"topics": {
							"edges": [
								{
									"node": {
										"name": "New York Times"
									}
								}
							]
						}
					}
				}
			}`,
		},
		{
			"handling of the viewer",
			`{
				"viewerScalar": [
					{
						"uid": "0x5d",
						"name": "Gnusto"
					}
				]
			}`,
			` {
				"data": {
					"viewer": {
						"name": "Gnusto"
					}
				}
			}`,
		},
	}

	runSerializerTests(t, &tables)
}
