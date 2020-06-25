# sortstringer
extended sort package for fmt.Stringer in golang

![ci](https://github.com/mashiike/sortstringer/workflows/Test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/mashiike/sortstringer)](https://goreportcard.com/report/github.com/mashiike/sortstringer)
[![Documentation](https://godoc.org/github.com/mashiike/sortstringer?status.svg)](http://godoc.org/github.com/mashiike/sortstringer)
[![license](https://img.shields.io/github/license/mashiike/sortstringer.svg?maxAge=2592000)](https://github.com/mashiike/sortstringer/LICENSE)
[![Version](https://img.shields.io/github/v/tag/mashiike/sortstringer.svg?label=Version)](https://github.com/mashiike/sortstringer/tags)

## Simple usecase

```
package main

import (
	"fmt"

	"github.com/mashiike/sortstringer"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return p.Name
}

func main() {
	persons := []*Person{
		&Person{Name: "Bob", Age: 15},
		&Person{Name: "Alice", Age: 15},
		&Person{Name: "Charlie", Age: 20},
		&Person{Name: "Carol", Age: 22},
	}
	sortstringer.Slice(persons, sortstringer.Ascending)
	fmt.Println(persons)
	//Output: [Alice Bob Carol Charlie]
}
```
