package main

import (
	"fmt"
	"strings"
)

type Stringer interface {
	String() string
}

type Person struct {
	First, Last string
}

func (p Person) String() string {
	return fmt.Sprintf("%s,%s", p.Last, p.First)
}

type StrList []string

func (s StrList) String() string {
	return strings.Join(s, ",")
}

func PrintStringer(s Stringer) {
	fmt.Println(s.String())
}

func main() {
	john := Person{First: "john", Last: "Doak"}
	var nameList Stringer = StrList{"David", "Sarah"}

	PrintStringer(john)
	PrintStringer(nameList)
}
