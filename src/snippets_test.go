package pmy

import "testing"

// TestIsApplicable tests if applicability of
// snippet is correctly works
func TestIsApplicable(t *testing.T) {
	testcases := []struct {
		description  string
		path         string
		basename     string
		relpath      string
		givenRelpath string
		expectedRes  bool
	}{
		{
			description:  "valid pattern",
			path:         "./snippets/curl/test.txt",
			basename:     "test.txt",
			relpath:      "curl/test.txt",
			givenRelpath: "curl/test",
			expectedRes:  true,
		},
		{
			description:  "invalid pattern",
			path:         "./snippets/curl/test.txt",
			basename:     "test.txt",
			relpath:      "curl/test.txt",
			givenRelpath: "curl/testhoge",
			expectedRes:  false,
		},
		//////////////////////////////////////////////////
		// !!!!!!!! TODO: deprecated test !!!!!!!!!!
		//////////////////////////////////////////////////
		{
			description:  "(deprecated) valid pattern",
			path:         "./snippets/curl/test_pmy_snippet.txt",
			basename:     "test_pmy_snippet.txt",
			relpath:      "curl/test_pmy_snippet.txt",
			givenRelpath: "curl/test",
			expectedRes:  true,
		},
		{
			description:  "(deprecated) valid pattern",
			path:         "./snippets/curl/test_pmy_snippet.txt",
			basename:     "test_pmy_snippet.txt",
			relpath:      "curl/testhoge_pmy_snippet.txt",
			givenRelpath: "curl/test",
			expectedRes:  false,
		},
	}

	for _, testcase := range testcases {
		target := &SnippetFile{
			Path:     testcase.path,
			Basename: testcase.basename,
			Relpath:  testcase.relpath,
		}
		res := target.isApplicable(testcase.givenRelpath)
		if !(res == testcase.expectedRes) {
			t.Errorf(
				"test failed (%v): resulting res %v is not equal to %v",
				testcase.description,
				res,
				testcase.expectedRes,
			)
		}
	}
}
