package dgraph

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	indent string = "  "
)

type Fragment struct {
	Name   string
	Query  string
	Fields *[]string
}

type QueryBuilder struct {
	Fragments *[]*Fragment
	Current   *Fragment
	Children  *[]*QueryBuilder
	Errors    *[]string
}

func NewBuilder() *QueryBuilder {
	return &QueryBuilder{
		Fragments: &[]*Fragment{},
		Children:  &[]*QueryBuilder{},
		Errors:    &[]string{},
	}
}

func (b *QueryBuilder) Fragment(name, query string) *QueryBuilder {
	b.Current = &Fragment{
		Name:   name,
		Query:  query,
		Fields: &[]string{},
	}
	*b.Fragments = append(*b.Fragments, b.Current)
	return b
}

func (b *QueryBuilder) Push() *QueryBuilder {
	child := NewBuilder()
	*b.Children = append(*b.Children, child)
	return child
}

func (b *QueryBuilder) Field(field string) {
	if b.Current == nil {
		b.Fragment("TestFragment", "test")
	}
	*b.Current.Fields = append(*b.Current.Fields, field)
}

func (b *QueryBuilder) Error(message string) {
	*b.Errors = append(*b.Errors, message)
}

func (b *QueryBuilder) ErrorString() string {
	return strings.Join(*b.Errors, "; ")
}

func (b *QueryBuilder) IsValid() bool {
	return len(*b.Errors) == 0
}

func (b *QueryBuilder) String() string {
	return fmt.Sprintf("QueryBuilder{Fragments: %#v, Current: %#v}", b.Fragments, b.Current)
}

func (b *QueryBuilder) writeTo(buf *bytes.Buffer, depth int) {
	prefix := strings.Repeat("\t", depth)
	for _, fragment := range *b.Fragments {
		buf.WriteString(prefix + fragment.Query + " {\n")
		for _, field := range *fragment.Fields {
			buf.WriteString(prefix + "\t" + field + "\n")
		}
		for _, child := range *b.Children {
			child.writeTo(buf, depth+1)
		}
		buf.WriteString(prefix + "}\n")
	}
}

func (b *QueryBuilder) Stringify() string {
	var buf bytes.Buffer
	buf.WriteString("{\n")
	b.writeTo(&buf, 1)
	buf.WriteString("}\n")
	return buf.String()
}
