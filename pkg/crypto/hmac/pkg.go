// Code generated by github.com/wylswz/cue-se/pkg/gen. DO NOT EDIT.

package hmac

import (
	"github.com/wylswz/cue-se/internal/core/adt"
	"github.com/wylswz/cue-se/internal/pkg"
)

func init() {
	pkg.Register("crypto/hmac", p)
}

var _ = adt.TopKind // in case the adt package isn't used

var p = &pkg.Package{
	Native: []*pkg.Builtin{{
		Name:  "MD5",
		Const: "\"MD5\"",
	}, {
		Name:  "SHA1",
		Const: "\"SHA1\"",
	}, {
		Name:  "SHA224",
		Const: "\"SHA224\"",
	}, {
		Name:  "SHA256",
		Const: "\"SHA256\"",
	}, {
		Name:  "SHA384",
		Const: "\"SHA384\"",
	}, {
		Name:  "SHA512",
		Const: "\"SHA512\"",
	}, {
		Name:  "SHA512_224",
		Const: "\"SHA512_224\"",
	}, {
		Name:  "SHA512_256",
		Const: "\"SHA512_256\"",
	}, {
		Name: "Sign",
		Params: []pkg.Param{
			{Kind: adt.StringKind},
			{Kind: adt.BytesKind | adt.StringKind},
			{Kind: adt.BytesKind | adt.StringKind},
		},
		Result: adt.BytesKind | adt.StringKind,
		Func: func(c *pkg.CallCtxt) {
			hashName, key, data := c.String(0), c.Bytes(1), c.Bytes(2)
			if c.Do() {
				c.Ret, c.Err = Sign(hashName, key, data)
			}
		},
	}},
}
