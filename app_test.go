package googege

import "testing"

// test Join function
func TestJoin(t *testing.T) {
	if testJoin() {
		t.Log("ok")
	} else {
		t.Error("Join存在问题")
	}
}

func testJoin() bool {
	sep, err := Join(nil, "1", "2", "3", "4")
	if sep != "1234" || err != nil {
		return false
	}
	return true
}

//  228 ns/op
func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join(nil, "1", "2", "3", "4", "5", "6", "7", "8", "9", "10")
	}
}
