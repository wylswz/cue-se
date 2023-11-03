// Code generated by github.com/wylswz/cue-se/pkg/gen. DO NOT EDIT.

package list

import (
	"github.com/wylswz/cue-se/internal/core/adt"
	"github.com/wylswz/cue-se/internal/pkg"
)

func init() {
	pkg.Register("list", p)
}

var _ = adt.TopKind // in case the adt package isn't used

var p = &pkg.Package{
	Native: []*pkg.Builtin{{
		Name: "Drop",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
			{Kind: adt.IntKind},
		},
		Result: adt.ListKind,
		Func: func(c *pkg.CallCtxt) {
			x, n := c.List(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = Drop(x, n)
			}
		},
	}, {
		Name: "FlattenN",
		Params: []pkg.Param{
			{Kind: adt.TopKind},
			{Kind: adt.IntKind},
		},
		Result: adt.ListKind,
		Func: func(c *pkg.CallCtxt) {
			xs, depth := c.Value(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = FlattenN(xs, depth)
			}
		},
	}, {
		Name: "Repeat",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
			{Kind: adt.IntKind},
		},
		Result: adt.ListKind,
		Func: func(c *pkg.CallCtxt) {
			x, count := c.List(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = Repeat(x, count)
			}
		},
	}, {
		Name: "Concat",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.ListKind,
		Func: func(c *pkg.CallCtxt) {
			a := c.List(0)
			if c.Do() {
				c.Ret, c.Err = Concat(a)
			}
		},
	}, {
		Name: "Take",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
			{Kind: adt.IntKind},
		},
		Result: adt.ListKind,
		Func: func(c *pkg.CallCtxt) {
			x, n := c.List(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = Take(x, n)
			}
		},
	}, {
		Name: "Slice",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
			{Kind: adt.IntKind},
			{Kind: adt.IntKind},
		},
		Result: adt.ListKind,
		Func: func(c *pkg.CallCtxt) {
			x, i, j := c.List(0), c.Int(1), c.Int(2)
			if c.Do() {
				c.Ret, c.Err = Slice(x, i, j)
			}
		},
	}, {
		Name: "MinItems",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
			{Kind: adt.IntKind},
		},
		Result: adt.BoolKind,
		Func: func(c *pkg.CallCtxt) {
			list, n := c.CueList(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = MinItems(list, n)
			}
		},
	}, {
		Name: "MaxItems",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
			{Kind: adt.IntKind},
		},
		Result: adt.BoolKind,
		Func: func(c *pkg.CallCtxt) {
			list, n := c.CueList(0), c.Int(1)
			if c.Do() {
				c.Ret, c.Err = MaxItems(list, n)
			}
		},
	}, {
		Name: "UniqueItems",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.BoolKind,
		Func: func(c *pkg.CallCtxt) {
			a := c.List(0)
			if c.Do() {
				c.Ret = UniqueItems(a)
			}
		},
	}, {
		Name: "Contains",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
			{Kind: adt.TopKind},
		},
		Result: adt.BoolKind,
		Func: func(c *pkg.CallCtxt) {
			a, v := c.List(0), c.Value(1)
			if c.Do() {
				c.Ret = Contains(a, v)
			}
		},
	}, {
		Name: "Avg",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.NumKind,
		Func: func(c *pkg.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Avg(xs)
			}
		},
	}, {
		Name: "Max",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.NumKind,
		Func: func(c *pkg.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Max(xs)
			}
		},
	}, {
		Name: "Min",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.NumKind,
		Func: func(c *pkg.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Min(xs)
			}
		},
	}, {
		Name: "Product",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.NumKind,
		Func: func(c *pkg.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Product(xs)
			}
		},
	}, {
		Name: "Range",
		Params: []pkg.Param{
			{Kind: adt.NumKind},
			{Kind: adt.NumKind},
			{Kind: adt.NumKind},
		},
		Result: adt.ListKind,
		Func: func(c *pkg.CallCtxt) {
			start, limit, step := c.Decimal(0), c.Decimal(1), c.Decimal(2)
			if c.Do() {
				c.Ret, c.Err = Range(start, limit, step)
			}
		},
	}, {
		Name: "Sum",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.NumKind,
		Func: func(c *pkg.CallCtxt) {
			xs := c.DecimalList(0)
			if c.Do() {
				c.Ret, c.Err = Sum(xs)
			}
		},
	}, {
		Name: "Sort",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
			{Kind: adt.TopKind},
		},
		Result: adt.ListKind,
		Func: func(c *pkg.CallCtxt) {
			list, cmp := c.List(0), c.Value(1)
			if c.Do() {
				c.Ret, c.Err = Sort(list, cmp)
			}
		},
	}, {
		Name: "SortStable",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
			{Kind: adt.TopKind},
		},
		Result: adt.ListKind,
		Func: func(c *pkg.CallCtxt) {
			list, cmp := c.List(0), c.Value(1)
			if c.Do() {
				c.Ret, c.Err = SortStable(list, cmp)
			}
		},
	}, {
		Name: "SortStrings",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.ListKind,
		Func: func(c *pkg.CallCtxt) {
			a := c.StringList(0)
			if c.Do() {
				c.Ret = SortStrings(a)
			}
		},
	}, {
		Name: "IsSorted",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
			{Kind: adt.TopKind},
		},
		Result: adt.BoolKind,
		Func: func(c *pkg.CallCtxt) {
			list, cmp := c.List(0), c.Value(1)
			if c.Do() {
				c.Ret = IsSorted(list, cmp)
			}
		},
	}, {
		Name: "IsSortedStrings",
		Params: []pkg.Param{
			{Kind: adt.ListKind},
		},
		Result: adt.BoolKind,
		Func: func(c *pkg.CallCtxt) {
			a := c.StringList(0)
			if c.Do() {
				c.Ret = IsSortedStrings(a)
			}
		},
	}},
	CUE: `{
	Comparer: {
		T:    _
		x:    T
		y:    T
		less: bool
	}
	Ascending: {
		Comparer
		T:    number | string
		x:    T
		y:    T
		less: x < y
	}
	Descending: {
		Comparer
		T:    number | string
		x:    T
		y:    T
		less: x > y
	}
}`,
}
