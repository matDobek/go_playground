package pointers

import (
	"testing"
)

type Person struct {
	firstName string
	lastName  string
}

// To access logs, use the '-v' verbose flag
// go test  -v ./internal/pointers/pointers_test.go
func TestPointers1(t *testing.T) {
	//
	// &variable - access address
	// *pointer - access value
	// *Type - part of type declaration, var will hold pointer to a Type
	//

	var str string = "Hello"
	var ptr *string = &str

	t.Log(str)  // Hello
	t.Log(ptr)  // 0x....
	t.Log(*ptr) // Hello

	str = "Hello World"
	t.Log(str)  // Hello World
	t.Log(*ptr) // Hello World

	*ptr = "Hello"
	t.Log(str)  // Hello
	t.Log(*ptr) // Hello
}

func pointers20(v int) {
	v += 1
	return
}
func pointers21(v string) {
	v = "bar"
	return
}
func pointers22(v [3]int) {
	v[0] = 99
	return
}
func pointers23(v []int) {
	v[0] = 99
	return
}
func pointers24(m map[string]string) {
	m["foo"] = "bar"
	return
}
func pointers25(v Person) {
	v.firstName = "Foo"
	v.lastName = "Bar"
	return
}

// In functions, do we pass by value, or by reference
func TestPointers2(t *testing.T) {
	var i int = 0
	pointers20(i)
	t.Log(i) // 0

	var str string = "foo"
	pointers21(str)
	t.Log(str) // "foo"

	var arr [3]int = [3]int{1, 2, 3}
	pointers22(arr)
	t.Log(arr) // [1, 2, 3]

	var slice []int = []int{1, 2, 3}
	pointers23(slice)
	t.Log(slice) // [99, 2, 3]

	var m map[string]string = map[string]string{"foo": "foo"}
	pointers24(m)
	t.Log(m) // map[foo: bar]

	//
	// Basic Data Type
	// int, float, string, bool, byte, rune, Array, Structs
	//
	// Referenced Data Type
	// slice, map
	//
}

func (p Person) changeName(first string, last string) {
	p.firstName = first
	p.lastName = last
}

func (p *Person) changeNamePtr(first string, last string) {
	p.firstName = first
	p.lastName = last
}

func changeNamePtr2(p *Person, first string, last string) {
	p.firstName = first
	p.lastName = last
}

func TestPointers3(t *testing.T) {
	var p Person = Person{"Mat", "Tam"}
	p.changeName("Foo", "Foo")
	t.Log(p) // {Mat, Tam}
	p.changeNamePtr("Bar", "Bar")
	t.Log(p) // {Bar, Bar}
	changeNamePtr2(&p, "Baz", "Baz")
	t.Log(p) // {Baz, Baz}
}
