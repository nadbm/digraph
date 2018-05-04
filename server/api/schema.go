package api

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/cayleygraph/cayley/quad"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
)

type Organization struct {
	_          struct{} `quad:"@type > foaf:Organization"`
	ID         string   `json:"id" quad:",optional"`
	ResourceID quad.IRI `json:"@id"`
	Name       string   `json:"name" quad:"di:name"`
}

type User struct {
	_          struct{} `quad:"@type > foaf:Person"`
	ID         string   `json:"id" quad:",optional"`
	ResourceID quad.IRI `json:"@id"`
	Name       string   `json:"name" quad:"di:name"`
	Email      string   `json:"email" quad:"di:email"`
}

type Topic struct {
	_           struct{} `quad:"@type > foaf:topic"`
	ID          string   `json:"id" quad:",optional"`
	ResourceID  quad.IRI `json:"@id"`
	Name        string   `json:"name" quad:"di:name"`
	Description *string  `json:"description" quad:"description,optional"`
}

type Link struct {
	_          struct{} `quad:"@type > di:link"`
	ID         string   `json:"id" quad:",optional"`
	ResourceID quad.IRI `json:"@id"`
	Title      string   `json:"title" quad:"di:title"`
	URL        string   `json:"url" quad:"di:url"`
}

type Resource interface {
	Init()
	IRI() quad.IRI
}

var (
	linkType         *Type
	nodeDefinitions  *relay.NodeDefinitions
	organizationType *Type
	topicType        *Type
	userType         *Type
)

var replacer = strings.NewReplacer("<", "", ">", "")

func isomorphicID(id quad.IRI) string {
	return replacer.Replace(id.Short().String())
}

func resourcePath(id quad.IRI) string {
	return replacer.Replace(id.Full().String())
}

func (o *User) Init() {
	o.ID = isomorphicID(o.ResourceID)
}

func (o *User) IRI() quad.IRI {
	return o.ResourceID
}

func (o *Organization) Init() {
	o.ID = isomorphicID(o.ResourceID)
}

func (o *Organization) IRI() quad.IRI {
	return o.ResourceID
}

func (o *Topic) Init() {
	o.ID = isomorphicID(o.ResourceID)
}

func (o *Topic) IRI() quad.IRI {
	return o.ResourceID
}

func (o *Link) Init() {
	o.ID = isomorphicID(o.ResourceID)
}

func (o *Link) IRI() quad.IRI {
	return o.ResourceID
}

func resolveType(p graphql.ResolveTypeParams) *graphql.Object {
	switch p.Value.(type) {
	case *Link:
		return linkType.NodeType
	case *Organization:
		return organizationType.NodeType
	case *Topic:
		return topicType.NodeType
	case *User:
		return userType.NodeType
	default:
		panic("unknown type")
	}
}

func (config *Config) fetcher() relay.IDFetcherFn {
	return func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
		resolvedID := relay.FromGlobalID(id)

		switch resolvedID.Type {
		case "Link":
			return config.Connection.FetchLink(resolvedID.ID)
		case "Organization":
			return config.Connection.FetchOrganization(resolvedID.ID)
		case "Topic":
			return config.Connection.FetchTopic(resolvedID.ID)
		case "User":
			return config.Connection.FetchUser(resolvedID.ID)
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
					if payload, ok := p.Source.(map[string]interface{}); ok {
						node, err := config.Connection.FetchTopic(payload["topicId"].(string))
						checkErr(err)
						return &relay.Edge{Node: node}, nil
					}
					return nil, nil
				},
			},
		},

		MutateAndGetPayload: func(input map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
			orgIri := input["organizationId"].(string)
			name := input["name"].(string)
			description := maybeString(input["description"])
			node, err := config.Connection.CreateTopic(orgIri, name, description)
			checkErr(err)

			return map[string]interface{}{
				"topicId": node.ID,
			}, nil
		},
	})
}

func (config *Config) createLinkMutation(edgeType graphql.Output) *graphql.Field {
	return relay.MutationWithClientMutationID(relay.MutationConfig{
		Name: "CreateLink",

		InputFields: graphql.InputObjectConfigFieldMap{
			"organizationId": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"title": &graphql.InputObjectFieldConfig{
				Type:         graphql.String,
				DefaultValue: nil,
			},
			"url": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"topicIds": &graphql.InputObjectFieldConfig{
				Type:         graphql.NewList(graphql.String),
				DefaultValue: []interface{}{},
			},
		},

		OutputFields: graphql.Fields{
			"linkEdge": &graphql.Field{
				Type: edgeType,

				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					payload := p.Source.(map[string]interface{})
					node, err := config.Connection.FetchLink(payload["linkId"].(string))
					checkErr(err)
					return &relay.Edge{Node: node}, nil
				},
			},
		},

		MutateAndGetPayload: func(input map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
			orgId := input["organizationId"].(string)
			url := input["url"].(string)
			topicIds := input["topicIds"].([]interface{})

			var useTitle string
			if title, ok := input["title"].(string); ok {
				useTitle = title
			}

			node, err := config.Connection.CreateLink(orgId, url, useTitle, topicIds)
			checkErr(err)

			return map[string]interface{}{
				"linkId": node.ID,
			}, nil
		},
	})
}

func (config *Config) selectTopicMutation(topicType *Type) *graphql.Field {
	return relay.MutationWithClientMutationID(relay.MutationConfig{
		Name: "SelectTopic",

		InputFields: graphql.InputObjectConfigFieldMap{
			"topicId": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},

		OutputFields: graphql.Fields{
			"topic": &graphql.Field{
				Type: topicType.NodeType,

				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if payload, ok := p.Source.(map[string]interface{}); ok {
						if topicId, ok := payload["topicId"]; ok {
							node, err := config.Connection.FetchTopic(topicId.(string))
							if err != nil {
								return nil, err
							}
							return node, nil
						}
					}
					return nil, nil
				},
			},
		},

		MutateAndGetPayload: func(input map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
			viewer, err := config.Connection.Viewer()
			checkErr(err)
			topicId := input["topicId"].(string)

			node, err := config.Connection.SelectTopic(viewer.(*User).ID, topicId)
			if err != nil {
				log.Println("there was a problem:", err)
				return map[string]interface{}{}, nil
			}

			return map[string]interface{}{
				"topicId": node.ID,
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

		FetchNode: func(resourceId string) (interface{}, error) {
			return config.Connection.FetchTopic(resourceId)
		},

		FetchConnection: func(out *[]interface{}, org *Organization) {
			checkErr(config.Connection.FetchTopics(out, org))
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

			"selectedTopic": &graphql.Field{
				Type:        topicType.NodeType,
				Description: "Topic selected by the user",

				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if res, ok := p.Source.(*User); ok {
						node, err := config.Connection.SelectedTopic(res.ID)
						if err != nil {
							return nil, err
						}
						return node, nil
					}
					return nil, errors.New("unable to provide IRI")
				},
			},
		},

		NodeDefinitions: nodeDefinitions,
	})

	linkType = NewType(&TypeConfig{
		Name: "Link",

		FetchNode: func(resourceId string) (interface{}, error) {
			return config.Connection.FetchLink(resourceId)
		},

		FetchConnection: func(out *[]interface{}, org *Organization) {
			checkErr(config.Connection.FetchLinks(out, org))
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

			"topics": &graphql.Field{
				Type: topicType.Definitions.ConnectionType,

				Args: relay.ConnectionArgs,

				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := relay.NewConnectionArguments(p.Args)
					dest := []interface{}{}
					link := p.Source.(*Link)
					config.Connection.FetchTopicsForLink(&dest, link)
					return relay.ConnectionFromArray(dest, args), nil
				},
			},
		},

		NodeDefinitions: nodeDefinitions,
	})

	topicType.NodeType.AddFieldConfig("links", &graphql.Field{
		Type: linkType.Definitions.ConnectionType,

		Args: relay.ConnectionArgs,

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			args := relay.NewConnectionArguments(p.Args)
			dest := []interface{}{}
			topic := p.Source.(*Topic)
			config.Connection.FetchLinksForTopic(&dest, topic)
			return relay.ConnectionFromArray(dest, args), nil
		},
	})

	organizationType = NewType(&TypeConfig{
		Name: "Organization",

		FetchNode: func(resourceId string) (interface{}, error) {
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
			"createLink":  config.createLinkMutation(linkType.Definitions.EdgeType),
			"selectTopic": config.selectTopicMutation(topicType),
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	return &schema, err
}
