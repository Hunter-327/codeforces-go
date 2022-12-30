// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/abc251/submit?taskScreenName=abc251_f
func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`6 8
5 1
4 3
1 4
3 5
1 2
2 6
1 6
4 2`,
			`1 4
4 3
5 3
4 2
6 2
1 5
5 3
1 4
2 1
1 6`,
		},
		{
			`4 3
3 1
1 2
1 4`,
			`1 2
1 3
1 4
1 4
1 3
1 2`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc251/tasks/abc251_f
