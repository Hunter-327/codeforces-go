// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	testutil2 "github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

var examples = [][]string{
	{
		`[[3,3,1],[3,2,1]]`,
		`[[4,3]]`,
		`2`,
		`1`,
	},
	{
		`[[1,3,2],[4,3,1],[7,1,2]]`,
		`[[1,0],[3,3]]`,
		`4`,
		`2`,
	},
	{
		`[[2,2,1]]`,
		`[[1,3],[1,4]]`,
		`3`,
		`1`,
	},
}

func Test(t *testing.T) {
	t.Log("Current test is [d]")

	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, circleGame, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode-cn.com/contest/season/2021-fall/problems/vFjcfV/

func TestCompareInf(t *testing.T) {
	inputGenerator := func() (toys [][]int, a [][]int, r int) {
		rg := testutil2.NewRandGenerator()
		n := rg.Int(1, 5)
		for i := 0; i < n; i++ {
			toys = append(toys, []int{rg.Int(1, 5), rg.Int(1, 5), rg.Int(1, 5)})
		}
		m := rg.Int(1, 5)
		for i := 0; i < m; i++ {
			a = append(a, []int{rg.Int(1, 5), rg.Int(1, 5)})
		}
		r = rg.Int(1, 5)
		return
	}

	runAC := func(toys [][]int, a [][]int, r int) (ans int) {
		for _, p := range toys {
			x, y, rp := p[0], p[1], p[2]
			if rp > r {
				continue
			}
			for _, q := range a {
				qx, qy := q[0], q[1]
				if (x-qx)*(x-qx)+(y-qy)*(y-qy) <= (r-rp)*(r-rp) {
					ans++
					break
				}
			}
		}
		return
	}

	// test examples first (or make it global)
	//if err := testutil.RunLeetCodeFuncWithExamples(t, runAC, examples, 0); err != nil {
	//	t.Fatal(err)
	//}
	//return

	testutil.CompareInf(t, inputGenerator, runAC, circleGame)
}