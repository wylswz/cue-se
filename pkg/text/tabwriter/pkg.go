// Code generated by github.com/wylswz/cue-se/pkg/gen. DO NOT EDIT.

package tabwriter

import (
	"github.com/wylswz/cue-se/internal/core/adt"
	"github.com/wylswz/cue-se/internal/pkg"
)

func init() {
	pkg.Register("text/tabwriter", p)
}

var _ = adt.TopKind // in case the adt package isn't used

var p = &pkg.Package{
	Native: []*pkg.Builtin{{
		Name: "Write",
		Params: []pkg.Param{
			{Kind: adt.TopKind},
		},
		Result: adt.StringKind,
		Func: func(c *pkg.CallCtxt) {
			data := c.Value(0)
			if c.Do() {
				c.Ret, c.Err = Write(data)
			}
		},
	}},
}
