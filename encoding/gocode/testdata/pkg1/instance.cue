package pkg1

import (
	"math"
	"strings"
	"github.com/wylswz/cue-se/encoding/gocode/testdata/pkg2"
)

MyStruct: {
	A:  <=10
	B:  =~"cat" | *"dog"
	O?: OtherStruct
	I:  pkg2.ImportMe
} @go(,complete=Complete)

OtherStruct: {
	A: strings.ContainsAny("X")
	P: pkg2.PickMe
}

String: !="" @go(,validate=ValidateCUE)

SpecialString: =~"special" @go(,type=string)

IgnoreThis: =~"foo" // No corresponding Go type

Omit: int @go(-)

// NonExisting will be omitted as there is no equivalent Go type.
NonExisting: {
	B: string
} @go(-)

// ignore unexported unless explicitly enabled.
foo: int

Ptr: {
	A: math.MultipleOf(10)
}
