package dgraph

import (
	"strings"
)

func squash(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
