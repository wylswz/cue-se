// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.

package csv

import (
	"github.com/wylswz/cue-se/internal/core/adt"
	"github.com/wylswz/cue-se/internal/pkg"
)

func init() {
	pkg.Register("encoding/csv", p)
}

var _ = adt.TopKind // in case the adt package isn't used

var p = &pkg.Package{
	Native: []*pkg.Builtin{{
		Name: "Encode",
		Params: []pkg.Param{
			{Kind: adt.TopKind},
		},
		Result: adt.StringKind,
		Func: func(c *pkg.CallCtxt) {
			x := c.Value(0)
			if c.Do() {
				c.Ret, c.Err = Encode(x)
			}
		},
	}, {
		Name: "Decode",
		Params: []pkg.Param{
			{Kind: adt.BytesKind | adt.StringKind},
		},
		Result: adt.ListKind,
		Func: func(c *pkg.CallCtxt) {
			r := c.Reader(0)
			if c.Do() {
				c.Ret, c.Err = Decode(r)
			}
		},
	}},
}
