package types

import (
	"strings"

	"github.com/cayleygraph/cayley/quad"
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
	Description *string  `json:"description" quad:"di:description,optional"`
	TopicIDs    []string `quad:",optional"`
}

type Link struct {
	_          struct{} `quad:"@type > di:link"`
	ID         string   `json:"id" quad:",optional"`
	ResourceID quad.IRI `json:"@id"`
	Title      string   `json:"title" quad:"di:title"`
	URL        string   `json:"url" quad:"di:url"`
	TopicIDs   []string `quad:",optional"`
}

var replacer = strings.NewReplacer("<", "", ">", "")

func IsomorphicID(id quad.IRI) string {
	return replacer.Replace(id.Short().String())
}

func ResourcePath(id quad.IRI) string {
	return replacer.Replace(id.Full().String())
}

func (o *User) Init() {
	o.ID = IsomorphicID(o.ResourceID)
}

func (o *User) IRI() quad.IRI {
	return o.ResourceID
}

func (o *Organization) Init() {
	o.ID = IsomorphicID(o.ResourceID)
}

func (o *Organization) IRI() quad.IRI {
	return o.ResourceID
}

func (o *Organization) ParentTopicIDs() []quad.IRI {
	return []quad.IRI{}
}

func (o *Topic) Init() {
	o.ID = IsomorphicID(o.ResourceID)
}

func (o *Topic) IRI() quad.IRI {
	return o.ResourceID
}

func (o *Topic) ParentTopicIDs() []quad.IRI {
	ids := make([]quad.IRI, len(o.TopicIDs))
	for i, topicId := range o.TopicIDs {
		ids[i] = quad.IRI(topicId)
	}
	return ids
}

func (o *Link) Init() {
	o.ID = IsomorphicID(o.ResourceID)
}

func (o *Link) ParentTopicIDs() []quad.IRI {
	ids := make([]quad.IRI, len(o.TopicIDs))
	for i, topicId := range o.TopicIDs {
		ids[i] = quad.IRI(topicId)
	}
	return ids
}

func (o *Link) IRI() quad.IRI {
	return o.ResourceID
}
