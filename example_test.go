package protobuf

import (
	"encoding/hex"
	"fmt"
)

type Test1 struct {
	A uint32
}

// This example encodes the simple message defined at the start of
// the Protocol Buffers encoding specification:
// https://developers.google.com/protocol-buffers/docs/encoding
func ExampleEncode_test1() {
	t := Test1{150}
	buf, _ := Encode(&t)
	fmt.Print(hex.Dump(buf))

	// Output:
	// 00000000  08 96 01                                          |...|
}

type Test2 struct {
	_ interface{} // = 1
	B string      // = 2
}

// This example encodes the Test2 message illustrating strings
// in the Protocol Buffers encoding specification.
func ExampleEncode_test2() {
	t := Test2{B: "testing"}
	buf, _ := Encode(&t)
	fmt.Print(hex.Dump(buf))

	// Output:
	// 00000000  12 07 74 65 73 74 69 6e  67                       |..testing|
}

type Test3 struct {
	_ interface{} // = 1
	_ interface{} // = 2
	C Test1       // = 3
}

// This example encodes the Test3 message illustrating embedded messages
// in the Protocol Buffers encoding specification.
func ExampleEncode_test3() {
	t := Test3{C: Test1{150}}
	buf, _ := Encode(&t)
	fmt.Print(hex.Dump(buf))

	// Output:
	// 00000000  1a 03 08 96 01                                    |.....|
}
