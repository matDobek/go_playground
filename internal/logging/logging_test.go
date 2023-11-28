package logging

import (
	"errors"
	"fmt"
	"testing"
)

type Foo struct {
	a int
	b int
}

func TestLogging1(t *testing.T) {
	v := Foo{1, 2}

	t.Logf("%v", v)  // Default format // {1 2}
	t.Logf("%#v", v) // Go-syntax foramt // logging.Foo{a:1, b:2}
	t.Logf("%T", v)  // The Type of the value // logging.Foo

	t.Logf("----------")

	str1 := "Foobar"
	str2 := "Bit longer Foobar"
	t.Logf("aaa %30s", str1) // justify right
	t.Logf("aaa %30s", str2)
	t.Logf("%-30s aaa", str1) // justify left
	t.Logf("%-30s aaa", str2)
	t.Logf("%q", str1) // quoted string

	t.Logf("----------")

	t.Logf("%d", 10)
	t.Logf("%t", true)

	t.Logf("----------")

	t.Logf("%e", 100.99)          // scientific notation
	t.Logf("%f", 100.99)          // decimal point
	t.Logf("%.1f", 100.99)        // precision 1
	t.Logf("%10.1f", 100.99)      // width 10, justify right, precision 1
	t.Logf("%-10.1f aaa", 100.99) // width 10, justify left, precision 1

	t.Logf("----------")

	t.Logf("foo\abar") // alert or bell
	t.Logf("foo\bbar") // backspace
	t.Logf("foo\\bar") // backslash
	t.Logf("foo\tbar") // horizontal tab
	t.Logf("foo\vbar") // vertical tab
	t.Logf("foo\rbar") // carriage return
	t.Logf("foo\nbar") // newline
	t.Logf("foo\fbar") // form feed

	t.Logf("----------")

	// %w, %v, %s for error formatting
	// %w most of the time
	// %v and %s may be useful for `nil` error values

	var err error = errors.New("Division by 0")
	var wrappedErr error = fmt.Errorf("Wrapped Error in Logging: %w", err)
	t.Logf(wrappedErr.Error())
}
