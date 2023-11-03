// Copyright 2021 CUE Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jsonpb_test

import (
	"testing"

	"github.com/wylswz/cue-se/cue"
	"github.com/wylswz/cue-se/cue/ast"
	"github.com/wylswz/cue-se/cue/errors"
	"github.com/wylswz/cue-se/cue/format"
	"github.com/wylswz/cue-se/cue/parser"
	"github.com/wylswz/cue-se/encoding/protobuf/jsonpb"
	"github.com/wylswz/cue-se/internal/cuetxtar"
)

func TestEncoder(t *testing.T) {
	test := cuetxtar.TxTarTest{
		Root: "./testdata/encoder",
		Name: "jsonpb",
	}

	r := cue.Runtime{}

	test.Run(t, func(t *cuetxtar.Test) {
		// TODO: use high-level API.

		var schema cue.Value
		var file *ast.File

		for _, f := range t.Archive.Files {
			switch {
			case f.Name == "schema.cue":
				inst, err := r.Compile(f.Name, f.Data)
				if err != nil {
					t.WriteErrors(errors.Promote(err, "test"))
					return
				}
				schema = inst.Value()

			case f.Name == "value.cue":
				f, err := parser.ParseFile(f.Name, f.Data, parser.ParseComments)
				if err != nil {
					t.WriteErrors(errors.Promote(err, "test"))
					return
				}
				file = f
			}
		}

		if !schema.Exists() {
			inst, err := r.CompileFile(file)
			if err != nil {
				t.WriteErrors(errors.Promote(err, "test"))
			}
			schema = inst.Value()
		}

		err := jsonpb.NewEncoder(schema).RewriteFile(file)
		if err != nil {
			errors.Print(t, err, nil)
		}

		b, err := format.Node(file)
		if err != nil {
			t.Fatal(err)
		}
		_, _ = t.Write(b)
	})
}
