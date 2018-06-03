package main

import "testing"

func TestSubstr(t *testing.T){
	tests := []struct{
		s string
		ans int
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},
		{"", 0},
		{"b", 1},
		{"bbbbbbbbbb", 1},
		{"abcabcabcabcd", 4},
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰灰化肥灰化肥挥发发黑会飞花", 5},
	}

	for _, tt := range tests {
		actual := maxlenstr(tt.s)
		if actual != tt.ans{
			t.Errorf("Got %d for input %s;" + "expected %d", actual, tt.s, tt.ans)
		}
	}
}
//性能测试
func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰灰化肥灰化肥挥发发黑会飞花"
	//因为可能原来的程序比较适合测试较长的字符串，因此这里直接扩展已有的字符串的大小
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	ans := 8
	//上面是准备数据的时间，测试时间不算上面的部分
	b.ResetTimer()
	//一次for相当于一次op
	for i := 0; i < b.N; i++ {
		actual := maxlenstr(s)
		if actual != ans{
			b.Errorf("Got %d for input %s;" + "expected %d", ans, s, actual)
		}
	}
}