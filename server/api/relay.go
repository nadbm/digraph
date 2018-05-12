package api

import (
	"github.com/cayleygraph/cayley/quad"
	"github.com/emwalker/digraph/server/types"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

type Type struct {
	ConnectionField *graphql.Field
	Definitions     *relay.GraphQLConnectionDefinitions
	NodeField       *graphql.Field
	NodeType        *graphql.Object
}

type ConnectionFetcher func(quad.IRI, *[]interface{})
type NodeFetcher func(quad.IRI, string) (interface{}, error)
type ConnectionResolver func(quad.IRI, Resource, *[]interface{})

type TypeConfig struct {
	FetchNode       NodeFetcher
	FetchConnection ConnectionFetcher
	Name            string
	NodeDefinitions *relay.NodeDefinitions
	NodeFields      graphql.Fields
}

var (
	resourceIdentifiableInterface *graphql.Interface
	resourceIdentifierField       *graphql.Field
	resourcePathField             *graphql.Field
)

func init() {
	resourceIdentifiableInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name: "ResourceIdentifiable",
		Fields: graphql.Fields{
			"resourceId": &graphql.Field{
				Type: graphql.String,
			},
			"resourcePath": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	resourceIdentifierField = &graphql.Field{
		Type:        graphql.String,
		Description: "The international resource identifier (IRI).",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			node := p.Source.(Resource)
			return types.IsomorphicID(node.IRI()), nil
		},
	}

	resourcePathField = &graphql.Field{
		Type:        graphql.String,
		Description: "The relative path to the resource.",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			node := p.Source.(Resource)
			return types.ResourcePath(node.IRI()), nil
		},
	}
}

func makeConnection(typ *Type, loadConnection ConnectionResolver) *graphql.Field {
	return &graphql.Field{
		Type: typ.Definitions.ConnectionType,

		Args: relay.ConnectionArgs,

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			args := relay.NewConnectionArguments(p.Args)
			out := []interface{}{}
			orgId := quad.IRI("organization:tyrell")
			loadConnection(orgId, p.Source.(Resource), &out)
			return relay.ConnectionFromArray(out, args), nil
		},
	}
}

func (config *TypeConfig) nodeType() *graphql.Object {
	fields := graphql.Fields{
		"id":           relay.GlobalIDField(config.Name, nil),
		"resourceId":   resourceIdentifierField,
		"resourcePath": resourcePathField,
	}

	for name, field := range config.NodeFields {
		fields[name] = field
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   config.Name,
		Fields: fields,
		Interfaces: []*graphql.Interface{
			config.NodeDefinitions.NodeInterface,
			resourceIdentifiableInterface,
		},
	})
}

func (config *TypeConfig) nodeField(nodeType graphql.Output) *graphql.Field {
	return &graphql.Field{
		Type: nodeType,

		Args: graphql.FieldConfigArgument{
			"resourceId": &graphql.ArgumentConfig{
				Description: "International resource identifier",
				Type:        graphql.NewNonNull(graphql.String),
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return config.FetchNode(quad.IRI("organization:tyrell"), p.Args["resourceId"].(string))
		},
	}
}

func (config *TypeConfig) connectionField(connectionType graphql.Output) *graphql.Field {
	return &graphql.Field{
		Type: connectionType,

		Args: relay.ConnectionArgs,

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			args := relay.NewConnectionArguments(p.Args)
			dest := []interface{}{}
			config.FetchConnection(quad.IRI("organization:tyrell"), &dest)
			return relay.ConnectionFromArray(dest, args), nil
		},
	}
}

func NewType(config *TypeConfig) *Type {
	nodeType := config.nodeType()

	defs := relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name:     config.Name,
		NodeType: nodeType,
	})

	return &Type{
		NodeType:        nodeType,
		Definitions:     defs,
		NodeField:       config.nodeField(nodeType),
		ConnectionField: config.connectionField(defs.ConnectionType),
	}
}
