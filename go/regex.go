// Port of Rob Pike's regex implementation from "The Practice of Programming"
// to the Go language.
package main

import "fmt"

func match(regexp string, text string) (rv int) {
	if regexp[0] == '^' {
		return matchhere(regexp[1:],text)
	}

	// Go doesn't support pointer access to individual elements of an array (http://golang.org/doc/go_spec.html#String_types).
	// The net result of the do loop + pointer arithmetic in the C code is to enumerate all substrings of text whose last
	// character is the last character of text.  We can do the same thing via Go's slice operations.
	for key,_ := range text {
		if matchhere(regexp,text[key:]) == 1 {
			return 1
		}
	}
	return 0
}

func matchhere(regexp string, text string) (rv int) {

	if len(regexp) == 0 { return 1 }

	// The first part of the text below isn't necessary in the C version; by checking for a first character of 
	// \0 before ever arriving at this test we're guaranteed that we've got at least one character.
	if len(regexp) > 1 && regexp[1] == '*' { return matchstar(regexp[0], regexp[2:], text) }
	if regexp[0] == '$' && len(regexp) == 1 { 

		// Ternary operator would be very useful here
		if len(text) == 0 { return 1 } else { return 0 }
	}
	if len(text) > 0 && (regexp[0] == '.' || regexp[0] == text[0]) { return matchhere(regexp[1:],text[1:]) }
	return 0
}

func matchstar(c byte, regexp string, text string) (rv int) {
	return 0
}

func main() {

  fmt.Printf("Basic match, beginning of string [1 expected]: %d\n",match("foo","foobar"));
  fmt.Printf("Basic match, middle of string [1 expected]: %d\n",match("oba","foobar"));
  fmt.Printf("Basic match, no match [0 expected]: %d\n",match("obo","foobar"));
  fmt.Printf("Match with start qualifier [1 expected]: %d\n",match("^fo","foobar"));
  fmt.Printf("Match with start qualifier in body [0 expected]: %d\n",match("^bar","foobar"));
  fmt.Printf("Match with end qualifier [1 expected]: %d\n",match("bar$","foobar"));
  fmt.Printf("Match with end qualifier in body [0 expected]: %d\n",match("foo$","foobar"));
  fmt.Printf("Match with optional qualifier [1 expected]: %d\n",match("fo*b","foobar"));
}