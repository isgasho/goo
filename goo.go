//goo.go this file incloud the common functions,maps slices, structs or anything else.
package goo

import "os"

// InSet is a bit array struct.
type IntSet struct {
	Words []uint64 // bit array
}

//Len return a int type value about the bit array's length.
func (i *IntSet) Len() int {
	return len(i.Words)
}

// remove this array's one bit.
func (i *IntSet) Remove(x int) {
	if i.Len() <= x {
		os.Exit(1)
	}
	i.Words = append(i.Words[:x], i.Words[x+1:]...)
}

// clear the array,and let  i.Words point a new slice.
func (i *IntSet) Clear() {
	i.Words = make([]uint64, 0)
}

// copy form the old array to a new array,return a new instance's pointer.
func (i *IntSet) Copy() *IntSet {
	t := *i
	return &t
}
