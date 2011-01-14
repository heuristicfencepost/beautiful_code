// Port of Rob Pike's regex implementation from "The Practice of Programming"
// to the Go language.

package main

import "fmt"

func match(regexp string, text string) (rv int) {
	return 0
}

func matchhere(regexp string, text string) (rv int) {
	return 0
}

func matchstar(c int, regexp string, text string) (rv int) {
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
