// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc241/tasks/abc241_d
// 提交：https://atcoder.jp/contests/abc241/submit?taskScreenName=abc241_d
// 对拍：https://atcoder.jp/contests/abc241/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc241_d&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc241/submissions?f.Status=AC&f.Task=abc241_d&orderBy=source_length
func Test_d(t *testing.T) {
	testCases := [][2]string{
		{
			`11
1 20
1 10
1 30
1 20
3 15 1
3 15 2
3 15 3
3 15 4
2 100 5
1 1
2 100 5`,
			`20
20
30
-1
-1
1`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
