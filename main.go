package main

import (
	"fmt"
)

type Reader interface {
	Read()
}
type Reader2 interface {
	Read()
}
type ReadWriter interface {
	Reader
	Write(str string)
}
type S struct {
	data string
}

func (s S) Read() {
	fmt.Println(s.data)
}

func (s S) setVal(str string) {
	s.data = str
	fmt.Println(s.data)
}

func (s *S) Write(str string) {
	s.data = str
}

type Ae struct {
	x int
}

func (a Ae) MethodA() {

}

type B struct {
	*Ae
}

type C struct {
	B
}

type P = *int

func main() {
	s1 := S{"a"}
	s2 := s1
	s2.data = "b"
	s1.Read()
	s2.Read()
}
