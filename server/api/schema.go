package api

import (
	"errors"
	"fmt"
	"log"

	"github.com/cayleygraph/cayley/quad"
	"github.com/emwalker/digraph/server/types"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
)

type Resource interface {
	Init()
	IRI() quad.IRI
	ParentTopicIDs() []quad.IRI
}

var (
	linkType         *Type
	nodeDefinitions  *relay.NodeDefinitions
	organizationType *Type
	topicType        *Type
	userType         *Type
)

func NewLink(node *types.Link, conn Connection) *types.Link {
	var useTitle string

	if node.Title == "" {
		var err error
		useTitle, err = conn.FetchTitle(node.URL)
		if err != nil {
			useTitle = node.URL
		}
		node.Title = useTitle
	}

	node.ResourceID = generateIDForType("link")
	node.Init()
	return node
}

func NewTopic(node *types.Topic) *types.Topic {
	node.ResourceID = generateIDForType("topic")
	node.Init()
	return node
}

func resolveType(p graphql.ResolveTypeParams) *graphql.Object {
	switch p.Value.(type) {
	case *types.Link:
		return linkType.NodeType
	case *types.Organization:
		return organizationType.NodeType
	case *types.Topic:
		return topicType.NodeType
	case *types.User:
		return userType.NodeType
	default:
		panic("unknown type")
	}
}

func (config *Config) fetcher() relay.IDFetcherFn {
	return func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
		resolvedID := relay.FromGlobalID(id)
		orgId := quad.IRI("organization:tyrell")

		switch resolvedID.Type {
		case "Organization":
			return config.Connection.FetchOrganization(resolvedID.ID)
		case "User":
			return config.Connection.FetchUser(resolvedID.ID)
		case "Link":
			return config.Connection.FetchLink(orgId, quad.IRI(resolvedID.ID))
		case "Topic":
			return config.Connection.FetchTopic(orgId, resolvedID.ID)
		default:
			return nil, errors.New(fmt.Sprintf("unknown node type: %s", resolvedID.Type))
		}
	}
}

func (config *Config) createTopicMutation(edgeType graphql.Output) *graphql.Field {
	return relay.MutationWithClientMutationID(relay.MutationConfig{
		Name: "CreateTopic",

		InputFields: graphql.InputObjectConfigFieldMap{
			"organizationId": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"topicIds": &graphql.InputObjectFieldConfig{
				Type:         graphql.NewList(graphql.String),
				DefaultValue: []interface{}{},
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"description": &graphql.InputObjectFieldConfig{
				Type:         graphql.String,
				DefaultValue: nil,
			},
		},

		OutputFields: graphql.Fields{
			"topicEdge": &graphql.Field{
				Type: edgeType,

				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					payload := p.Source.(map[string]interface{})
					orgId := payload["organizationId"].(quad.IRI)
					node, err := config.Connection.FetchTopic(orgId, payload["topicId"].(string))
					checkErr(err)
					return &relay.Edge{Node: node}, nil
				},
			},
		},

		MutateAndGetPayload: func(input map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
			topicIds := stringList(input["topicIds"])
			orgId := quad.IRI(input["organizationId"].(string))
			node := NewTopic(&types.Topic{
				Name:        input["name"].(string),
				Description: maybeString(input["description"]),
				TopicIDs:    *topicIds,
			})
			checkErr(config.Connection.CreateTopic(orgId, node))

			return map[string]interface{}{
				"topicId":        node.ID,
				"organizationId": orgId,
			}, nil
		},
	})
}

func (config *Config) upsertLinkMutation(edgeType graphql.Output) *graphql.Field {
	return relay.MutationWithClientMutationID(relay.MutationConfig{
		Name: "UpsertLink",

		InputFields: graphql.InputObjectConfigFieldMap{
			"organizationId": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"resourceId": &graphql.InputObjectFieldConfig{
				Type:         graphql.String,
				DefaultValue: nil,
			},
			"addTopicIds": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(graphql.String),
			},
			"title": &graphql.InputObjectFieldConfig{
				Type:         graphql.String,
				DefaultValue: nil,
			},
			"url": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},

		OutputFields: graphql.Fields{
			"linkEdge": &graphql.Field{
				Type: edgeType,

				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					payload := p.Source.(map[string]interface{})
					orgId := payload["organizationId"].(quad.IRI)
					linkId := quad.IRI(payload["linkId"].(string))
					node, err := config.Connection.FetchLink(orgId, linkId)
					checkErr(err)
					return &relay.Edge{Node: node}, nil
				},
			},
		},

		MutateAndGetPayload: func(input map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
			resourceId := input["resourceId"]
			title := input["title"]
			topicIds := stringList(input["addTopicIds"])
			orgId := quad.IRI(input["organizationId"].(string))
			url := input["url"].(string)

			var node interface{}
			var err error

			if resourceId, ok := resourceId.(string); ok {
				node, err = config.Connection.FetchLink(orgId, quad.IRI(resourceId))
			} else {
				log.Println("Determining whether url is already present:", url)
				node, err = config.Connection.FetchLinkByURL(orgId, url)
				checkErr(err)
			}

			if node == nil {
				log.Println("Link is new:", url)
				node = NewLink(&types.Link{
					URL:      url,
					Title:    stringOr("", title),
					TopicIDs: *topicIds,
				}, config.Connection)
				checkErr(config.Connection.UpsertLink(orgId, node.(*types.Link), true))
			} else {
				log.Println("Link is already in datastore:", url)
				link := node.(*types.Link)
				link.TopicIDs = *topicIds
				link.URL = url
				if title != nil {
					link.Title = title.(string)
				}
				checkErr(config.Connection.UpsertLink(orgId, link, false))
			}

			return map[string]interface{}{
				"linkId":         node.(*types.Link).ID,
				"organizationId": orgId,
			}, nil
		},
	})
}

func (config *Config) viewerField(userType *Type) *graphql.Field {
	return &graphql.Field{
		Type: userType.NodeType,

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return config.Connection.Viewer()
		},
	}
}

func (config *Config) newSchema() (*graphql.Schema, error) {
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher:   config.fetcher(),
		TypeResolve: resolveType,
	})

	topicType = NewType(&TypeConfig{
		Name: "Topic",

		FetchNode: func(orgId quad.IRI, resourceId string) (interface{}, error) {
			return config.Connection.FetchTopic(orgId, resourceId)
		},

		FetchConnection: func(orgId quad.IRI, out *[]interface{}) {
			checkErr(config.Connection.FetchTopics(orgId, out))
		},

		NodeFields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Name of the topic",
			},
			"description": &graphql.Field{
				Type:        graphql.String,
				Description: "Description of the topic",
			},
		},

		NodeDefinitions: nodeDefinitions,
	})

	userType = NewType(&TypeConfig{
		Name: "User",

		NodeFields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Name of the user",
			},

			"email": &graphql.Field{
				Type:        graphql.String,
				Description: "Email address of the user",
			},
		},

		NodeDefinitions: nodeDefinitions,
	})

	linkType = NewType(&TypeConfig{
		Name: "Link",

		FetchNode: func(orgId quad.IRI, resourceId string) (interface{}, error) {
			return config.Connection.FetchLink(orgId, quad.IRI(resourceId))
		},

		FetchConnection: func(orgId quad.IRI, out *[]interface{}) {
			checkErr(config.Connection.FetchLinks(orgId, out))
		},

		NodeFields: graphql.Fields{
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "Title of the page",
			},

			"url": &graphql.Field{
				Type:        graphql.String,
				Description: "Url of the page",
			},
		},

		NodeDefinitions: nodeDefinitions,
	})

	linkType.NodeType.AddFieldConfig("topics",
		makeConnection(topicType, func(orgId quad.IRI, parent Resource, out *[]interface{}) {
			config.Connection.FetchTopicsForLink(orgId, parent.IRI(), out)
		}),
	)

	topicType.NodeType.AddFieldConfig("links",
		makeConnection(linkType, func(orgId quad.IRI, parent Resource, out *[]interface{}) {
			config.Connection.FetchLinksForTopic(orgId, parent.IRI(), out)
		}),
	)

	topicType.NodeType.AddFieldConfig("parentTopics",
		makeConnection(topicType, func(orgId quad.IRI, parent Resource, out *[]interface{}) {
			config.Connection.FetchParentTopicsForTopic(orgId, parent.IRI(), out)
		}),
	)

	topicType.NodeType.AddFieldConfig("childTopics",
		makeConnection(topicType, func(orgId quad.IRI, parent Resource, out *[]interface{}) {
			config.Connection.FetchChildTopicsForTopic(orgId, parent.IRI(), out)
		}),
	)

	organizationType = NewType(&TypeConfig{
		Name: "Organization",

		FetchNode: func(orgId quad.IRI, resourceId string) (interface{}, error) {
			return config.Connection.FetchOrganization(resourceId)
		},

		NodeFields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Name of the organization",
			},
			"topic":  topicType.NodeField,
			"topics": topicType.ConnectionField,
			"links":  linkType.ConnectionField,
		},

		NodeDefinitions: nodeDefinitions,
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"viewer":       config.viewerField(userType),
			"organization": organizationType.NodeField,
			"user":         userType.NodeField,
			"node":         nodeDefinitions.NodeField,
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createTopic": config.createTopicMutation(topicType.Definitions.EdgeType),
			"upsertLink":  config.upsertLinkMutation(linkType.Definitions.EdgeType),
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	return &schema, err
}
