package app

import (
	"strings"
	"testing"
)

// test Join function
func TestJoin(t *testing.T) {
	if testJoin() {
		t.Log("ok")
	} else {
		t.Error("Join存在问题")
	}
}

func testJoin() bool {
	sep, err := Join(nil, "", "1", "2", "3", "4")
	if sep != "1234" || err != nil {
		return false
	}
	return true
}

//  565 ns/op
func BenchmarkJoin1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, "1", "2", "3", "4", "5", "6", "7", "8", "9", "10")
	}
}

// 1041 ns/op
func BenchmarkJoin2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "1",
			"2", "3", "4", "5", "6", "7", "8", "9", "10",
			"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		)
	}
}

// 64.2 ns/op
func BencharkJoinGoLibrary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a1 := strings.Join([]string{"1"}, "")
		a2 := strings.Join([]string{a1, "2"}, "")
		a3 := strings.Join([]string{a2, "3"}, "")
		a4 := strings.Join([]string{a3, "4"}, "")
		a5 := strings.Join([]string{a4, "5"}, "")
		a6 := strings.Join([]string{a5, "6"}, "")
		a7 := strings.Join([]string{a6, "7"}, "")
		a8 := strings.Join([]string{a7, "8"}, "")
		a9 := strings.Join([]string{a8, "9"}, "")
		strings.Join([]string{a9, "10"}, "")
	}
}

// 0.37 ns/op
func BenchmarkOnePlusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "1" + "2" + "3" + "4" + "5" + "6" + "7" + "8" + "9" + "10"
	}
}
