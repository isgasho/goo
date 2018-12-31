# googege network library
[![GoDoc](https://godoc.org/github.com/googege/goo?status.svg)](https://godoc.org/github.com/googege/goo)
### What?
this is a network library which created by googege.This library contains a lot of wrapper functions in common web processing.
### Install
```bash
go get github.com/googege/goo
```
### Usage
```go
package main

import (
  "fmt"
	"github.com/googege/goo"
)

func main() {
	value := goo.Join(nil, "a", "b", "c")
  fmt.Println(value)
}
//abc
```
> [Document in md format](./doc.md)
