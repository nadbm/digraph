// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type Alert struct {
	Text string    `json:"text"`
	Type AlertType `json:"type"`
	ID   string    `json:"id"`
}

type Alertable interface {
	IsAlertable()
}

type LinkConnection struct {
	Edges    []*LinkEdge `json:"edges"`
	PageInfo PageInfo    `json:"pageInfo"`
}

type LinkEdge struct {
	Cursor string `json:"cursor"`
	Node   Link   `json:"node"`
}

type Namespaceable interface {
	IsNamespaceable()
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type ResourceIdentifiable interface {
	IsResourceIdentifiable()
}

type SelectTopicInput struct {
	OrganizationID string `json:"organizationId"`
	TopicID        string `json:"topicId"`
}

type SelectTopicPayload struct {
	Topic *Topic `json:"topic"`
}

type TopicConnection struct {
	Edges    []*TopicEdge `json:"edges"`
	PageInfo PageInfo     `json:"pageInfo"`
}

type TopicEdge struct {
	Cursor string `json:"cursor"`
	Node   Topic  `json:"node"`
}

type UpdateLinkTopicsInput struct {
	ClientMutationID *string  `json:"clientMutationId"`
	LinkID           string   `json:"linkId"`
	ParentTopicIds   []string `json:"parentTopicIds"`
}

type UpdateLinkTopicsPayload struct {
	Link Link `json:"link"`
}

type UpdateTopicInput struct {
	ClientMutationID *string  `json:"clientMutationId"`
	Description      *string  `json:"description"`
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	OrganizationID   string   `json:"organizationId"`
	TopicIds         []string `json:"topicIds"`
}

type UpdateTopicParentTopicsInput struct {
	ClientMutationID *string  `json:"clientMutationId"`
	TopicID          string   `json:"topicId"`
	ParentTopicIds   []string `json:"parentTopicIds"`
}

type UpdateTopicParentTopicsPayload struct {
	Alerts []Alert `json:"alerts"`
	Topic  Topic   `json:"topic"`
}

func (UpdateTopicParentTopicsPayload) IsAlertable() {}

type UpdateTopicPayload struct {
	Topic Topic `json:"topic"`
}

type UpsertLinkInput struct {
	AddParentTopicIds []string `json:"addParentTopicIds"`
	ClientMutationID  *string  `json:"clientMutationId"`
	OrganizationID    string   `json:"organizationId"`
	Title             *string  `json:"title"`
	URL               string   `json:"url"`
}

type UpsertLinkPayload struct {
	Alerts   []Alert   `json:"alerts"`
	LinkEdge *LinkEdge `json:"linkEdge"`
}

func (UpsertLinkPayload) IsAlertable() {}

type UpsertTopicInput struct {
	ClientMutationID *string  `json:"clientMutationId"`
	Description      *string  `json:"description"`
	Name             string   `json:"name"`
	OrganizationID   string   `json:"organizationId"`
	TopicIds         []string `json:"topicIds"`
}

type UpsertTopicPayload struct {
	Alerts    []Alert    `json:"alerts"`
	TopicEdge *TopicEdge `json:"topicEdge"`
}

func (UpsertTopicPayload) IsAlertable() {}

type AlertType string

const (
	AlertTypeSuccess AlertType = "SUCCESS"
	AlertTypeWarn    AlertType = "WARN"
	AlertTypeError   AlertType = "ERROR"
)

func (e AlertType) IsValid() bool {
	switch e {
	case AlertTypeSuccess, AlertTypeWarn, AlertTypeError:
		return true
	}
	return false
}

func (e AlertType) String() string {
	return string(e)
}

func (e *AlertType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AlertType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AlertType", str)
	}
	return nil
}

func (e AlertType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
