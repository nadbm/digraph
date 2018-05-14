package dgraph

import (
	"log"
	"reflect"

	"github.com/emwalker/digraph/server/common"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Organization struct {
	common.Organization
	Links  *[]common.Link  `json:"linkConnection,omitempty"`
	Topics *[]common.Topic `json:"topicConnection,omitempty"`
}

type Viewer struct {
	common.User
}

type Query struct {
	Organization *[]*Organization `json:"organizationScalar,omitempty"`
	Viewer       *[]*Viewer       `json:"viewerScalar,omitempty"`
}

func firstOr(defaultValue interface{}, it interface{}) interface{} {
	ptr := reflect.ValueOf(it)
	array := reflect.Indirect(ptr)

	if !array.IsValid() {
		return defaultValue
	}

	if array.Len() < 1 {
		return defaultValue
	}
	return array.Index(0).Interface()
}

func (q *Query) MarshalJSON() ([]byte, error) {
	type data struct {
		Organization interface{} `json:"organization,omitempty"`
		Viewer       interface{} `json:"viewer,omitempty"`
	}
	return json.Marshal(&struct {
		Data *data `json:"data"`
	}{
		Data: &data{
			Organization: firstOr(nil, q.Organization),
			Viewer:       firstOr(nil, q.Viewer),
		},
	})
}

type edge struct {
	Node interface{} `json:"node"`
}

type connection struct {
	Edges []edge `json:"edges"`
}

func newConnection(it interface{}) *connection {
	ptr := reflect.ValueOf(it)
	array := reflect.Indirect(ptr)

	if !array.IsValid() {
		return nil
	}

	edges := make([]edge, 0)

	for i := 0; i < array.Len(); i++ {
		value := array.Index(i).Interface()
		edges = append(edges, edge{Node: value})
	}

	return &connection{Edges: edges}
}

func (o *Organization) MarshalJSON() ([]byte, error) {
	type Copy Organization

	return json.Marshal(&struct {
		Links           *connection `json:"links,omitempty"`
		Topics          *connection `json:"topics,omitempty"`
		TopicConnection interface{} `json:"topicConnection,omitempty"`
		LinkConnection  interface{} `json:"linkConnection,omitempty"`
		*Copy
	}{
		Links:           newConnection(o.Links),
		Topics:          newConnection(o.Topics),
		Copy:            (*Copy)(o),
		TopicConnection: nil,
		LinkConnection:  nil,
	})
}

func Deserialize(input []byte) *Query {
	var query Query
	err := json.Unmarshal(input, &query)
	if err != nil {
		log.Fatal(err)
	}
	return &query
}

func Serialize(input []byte) []byte {
	query := Deserialize(input)
	buf, err := json.Marshal(query)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}
