package regex

import "testing"

func TestRegex(test *testing.T) {

	// Basic match, beginning of string
	if match("foo","foobar") != 1 { test.Fail() }

	// Basic match, middle of string
	if match("oba","foobar") != 1 { test.Fail() }

	// Basic match, no match 
	if match("obo","foobar") != 0 { test.Fail() }

	// Match with start qualifier
	if match("^fo","foobar") != 1 { test.Fail() }

	// Match with start qualifier in body
	if match("^bar","foobar") != 0 { test.Fail() }

	// Match with end qualifier 
	if match("bar$","foobar") != 1 { test.Fail() }

	// Match with end qualifier in body
	if match("foo$","foobar") != 0 { test.Fail() }

	// Match with optional qualifier
	if match("fo*b","foobar") != 1 { test.Fail() }

	// Match with optional qualifier 2
	if match("fooa*b","foobar") != 1 { test.Fail() }

	// Match with optional qualifier 3
	if match("a*foo","foobar") != 1 { test.Fail() }
}