// Code generated by copypasta/template/acwing/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`6
4 8 6 3 12 9`,
			`9 3 6 12 4 8`,
		},
		{
			`4
42 28 84 126`,
			`126 42 84 28`,
		},
		{
			`2
1000000000000000000 3000000000000000000`,
			`3000000000000000000 1000000000000000000`,
		},
	}
	target := 0 // -1
	testutil.AssertEqualStringCase(t, testCases, target, run)
}