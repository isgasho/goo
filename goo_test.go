package goo

import (
	"fmt"
	"testing"
	"github.com/googege/goo/app"
)

func TestInset(t *testing.T) {
	i := &IntSet{
		[]uint64{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		},
	}

	if i.Len() != 10 {
		t.Error("the i.Len() method is wrong.")
	}

	d := i.Copy()
	for k, v := range d.Words {
		if d.Len() != i.Len() {
			t.Error("Copy is wrong")
		}
		if v != i.Words[k] {
			t.Error("copy is wrong")
		}
	}

	i.Remove(1)
	if i.Words[1] == 2 {
		t.Error("Remove is wrong")
	}
	i.Clear()
	if len(i.Words) != 0 {
		t.Error("Clear is wrong")
	}
}

//  60.3 ns/op
func BenchmarkInset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := &IntSet{
			[]uint64{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			},
		}
		i.Len()
		i.Remove(1)
		i.Copy()
		i.Clear()
	}
}

// ip

func TestWhichCountry(t *testing.T) {
	WhichCountry("173.82.115.125")
}

func BenchmarkWhichCountry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WhichCountry("173.82.115.125", "173.82.115.125", "173.82.115.125", "173.82.115.125", "173.82.115.125")
	}
}

func ExampleWhichCountry() {
	v := WhichCountry("173.82.115.125")
	for _, b := range v {
		fmt.Println(b.Country, b.City)
		//output: 美国 洛杉矶
	}
}

// test app

func Testapp(t *testing.T) {
	app.Join(nil,"1")
}
