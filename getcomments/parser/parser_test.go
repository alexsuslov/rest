package parser_test

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"rest/getcomments/parser"
	"rest/getcomments/parser/tests/anonymous"
	"rest/getcomments/parser/tests/chans"
	"rest/getcomments/parser/tests/docs"
	"rest/getcomments/parser/tests/enum"
	"rest/getcomments/parser/tests/functions"
	"rest/getcomments/parser/tests/functiontypes"
	"rest/getcomments/parser/tests/pointers"
	"rest/getcomments/parser/tests/privatetypes"
	"rest/getcomments/parser/tests/publictypes"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name     string
		pkg      string
		expected string
	}{
		{
			name:     "private structs are ignored",
			pkg:      "github.com/a-h/rest/getcomments/parser/tests/privatetypes",
			expected: privatetypes.Expected,
		},
		{
			name:     "public structs are included",
			pkg:      "github.com/a-h/rest/getcomments/parser/tests/publictypes",
			expected: publictypes.Expected,
		},
		{
			name:     "string and integer enums are supported",
			pkg:      "github.com/a-h/rest/getcomments/parser/tests/enum",
			expected: enum.Expected,
		},
		{
			name:     "pointers to pointers become a single pointer",
			pkg:      "github.com/a-h/rest/getcomments/parser/tests/pointers",
			expected: pointers.Expected,
		},
		{
			name:     "functions and method receivers are ignored",
			pkg:      "github.com/a-h/rest/getcomments/parser/tests/functions",
			expected: functions.Expected,
		},
		{
			name:     "fields of type channel are ignored",
			pkg:      "github.com/a-h/rest/getcomments/parser/tests/chans",
			expected: chans.Expected,
		},
		{
			name:     "anonymous structs are ignored",
			pkg:      "github.com/a-h/rest/getcomments/parser/tests/anonymous",
			expected: anonymous.Expected,
		},
		{
			name:     "function fields and function types are ignored",
			pkg:      "github.com/a-h/rest/getcomments/parser/tests/functiontypes",
			expected: functiontypes.Expected,
		},
		{
			name:     "stuct, field and constant comments are extracted",
			pkg:      "github.com/a-h/rest/getcomments/parser/tests/docs",
			expected: docs.Expected,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			m, err := parser.Get(test.pkg)
			if err != nil {
				t.Fatalf("failed to get model %q: %v", test.pkg, err)
			}

			var expected map[string]string
			err = json.Unmarshal([]byte(test.expected), &expected)
			if err != nil {
				t.Fatalf("snapshot load failed: %v", err)
			}

			if diff := cmp.Diff(expected, m); diff != "" {
				t.Error(diff)
			}
		})
	}
}
