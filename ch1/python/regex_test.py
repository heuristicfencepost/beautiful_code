import unittest

from regex import match

class regex_test (unittest.TestCase):

    def testRegex(self):

        # Basic match, beginning of string
        self.assertEqual(1,match("foo","foobar"))

        # Basic match, middle of string
        self.assertEqual(1,match("oba","foobar"))

        # Basic match, no match 
	self.assertEqual(0,match("obo","foobar"))

        # Match with start qualifier
        self.assertEqual(1,match("^fo","foobar"))

        # Match with start qualifier in body
        self.assertEqual(0,match("^bar","foobar"))

        # Match with end qualifier 
        self.assertEqual(1,match("bar$","foobar"))

        # Match with end qualifier in body
        self.assertEqual(0,match("foo$","foobar"))

        # Match with optional qualifier
        self.assertEqual(1,match("fo*b","foobar"))

        # Match with optional qualifier 2
        self.assertEqual(1,match("fooa*b","foobar"))

        # Match with optional qualifier 3
        self.assertEqual(1,match("a*foo","foobar"))


if __name__ == '__main__':
    unittest.main()
