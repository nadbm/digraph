package dgraph

import (
	"fmt"
	"reflect"
	"testing"
)

type TestFragment struct {
	name   string
	query  string
	fields []string
}

type builderTestCase struct {
	description string
	fragments   []TestFragment
	expected    string
}

func runBuilderTests(t *testing.T, tables *[]builderTestCase) {
	for _, table := range *tables {
		builder := NewBuilder()
		for _, fragment := range table.fragments {
			builder.Fragment(fragment.name, fragment.query)
			for _, field := range fragment.fields {
				builder.Field(field)
			}
		}
		actual := builder.Stringify()
		if squash(actual) != squash(table.expected) {
			t.Errorf(`
				Output did not match expected output for test: %s:
				%s
			`, table.description, actual)
		}
	}
}

func TestBuilder(t *testing.T) {
	tables := []builderTestCase{
		{
			"test fragments",
			[]TestFragment{
				TestFragment{
					"TestFragment",
					"links(query string)",
					[]string{"title", "url"},
				},
			},
			`{
				links(query string) {
					title
					url
				}
			}`,
		},
	}

	runBuilderTests(t, &tables)
}

func TestErrors(t *testing.T) {
	message1 := "Something happened: ..."
	message2 := "Something else happened: ..."

	b := NewBuilder()
	b.Error(message1)
	b.Error(message2)
	expected := []string{message1, message2}

	if !reflect.DeepEqual(*b.Errors, expected) {
		t.Errorf("error not saved: %+v", b.Errors)
	}

	if b.ErrorString() != fmt.Sprintf("%s; %s", message1, message2) {
		t.Errorf("unexpected error string: %s", b.ErrorString())
	}
}
