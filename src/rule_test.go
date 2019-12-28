package pmy

import "testing"

// TestSetPriority tests if setter of priority
// for RuleFile is correctly works
func TestSetPriority(t *testing.T) {
	testcases := []struct {
		description      string
		path             string
		basename         string
		cmd              string
		expectedPriority int
	}{
		{
			description:      "global yml rule is correctly judged",
			path:             "./hoge/_.yml",
			basename:         "_.yml",
			cmd:              "ls",
			expectedPriority: 1,
		},
		{
			description:      "global yaml rule is correctly judged",
			path:             "./hoge/_.yaml",
			basename:         "_.yaml",
			cmd:              "ls",
			expectedPriority: 1,
		},
		{
			description:      "global json rule is correctly judged",
			path:             "./hoge/_.json",
			basename:         "_.json",
			cmd:              "ls",
			expectedPriority: 1,
		},
		{
			description:      "command specific yml rule is correctly judged",
			path:             "./hoge/ls.yml",
			basename:         "ls.yml",
			cmd:              "ls",
			expectedPriority: 2,
		},
		{
			description:      "command specific yaml rule is correctly judged",
			path:             "./hoge/ls.yaml",
			basename:         "ls.yaml",
			cmd:              "ls",
			expectedPriority: 2,
		},
		{
			description:      "command specific json rule is correctly judged",
			path:             "./hoge/ls.json",
			basename:         "ls.json",
			cmd:              "ls",
			expectedPriority: 2,
		},
		//////////////////////////////////////////////////
		// !!!!!!!! TODO: deprecated test !!!!!!!!!!
		//////////////////////////////////////////////////
		{
			description:      "(deprecated) global yml rule is correctly judged",
			path:             "./hoge/pmy_rules.yml",
			basename:         "pmy_rules.yml",
			cmd:              "ls",
			expectedPriority: 1,
		},
		{
			description:      "(deprecated) global yaml rule is correctly judged",
			path:             "./hoge/pmy_rules.yaml",
			basename:         "pmy_rules.yaml",
			cmd:              "ls",
			expectedPriority: 1,
		},
		{
			description:      "(deprecated) global json rule is correctly judged",
			path:             "./hoge/pmy_rules.json",
			basename:         "pmy_rules.json",
			cmd:              "ls",
			expectedPriority: 1,
		},
		{
			description:      "(deprecated) command specific yml rule is correctly judged",
			path:             "./hoge/ls_pmy_rules.yml",
			basename:         "ls_pmy_rules.yml",
			cmd:              "ls",
			expectedPriority: 2,
		},
		{
			description:      "(deprecated) command specific yaml rule is correctly judged",
			path:             "./hoge/ls_pmy_rules.yaml",
			basename:         "ls_pmy_rules.yaml",
			cmd:              "ls",
			expectedPriority: 2,
		},
		{
			description:      "(deprecated) command specific json rule is correctly judged",
			path:             "./hoge/ls_pmy_rules.json",
			basename:         "ls_pmy_rules.json",
			cmd:              "ls",
			expectedPriority: 2,
		},
	}

	for _, testcase := range testcases {
		target := &RuleFile{
			Path:     testcase.path,
			Basename: testcase.basename,
			priority: 0,
		}
		target.setPriority(testcase.cmd)
		if !(target.priority == testcase.expectedPriority) {
			t.Errorf(
				"test failed (%v): resulting priority %v is not equal to %v",
				testcase.description,
				target.priority,
				testcase.expectedPriority,
			)
		}
	}
}
