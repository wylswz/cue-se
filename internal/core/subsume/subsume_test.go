// Copyright 2020 CUE Authors
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

package subsume

import (
	"testing"

	"github.com/wylswz/cue-se/cue/parser"
	"github.com/wylswz/cue-se/internal/core/adt"
	"github.com/wylswz/cue-se/internal/core/compile"
	"github.com/wylswz/cue-se/internal/core/eval"
	"github.com/wylswz/cue-se/internal/core/runtime"
)

// For debugging purposes, do not remove.
func TestX(t *testing.T) {
	t.Skip()

	r := runtime.New()
	ctx := eval.NewContext(r, nil)

	const gt = `a: *1 | int`
	const lt = `a: (*1 | int) & 1`

	a := parse(t, ctx, gt)
	b := parse(t, ctx, lt)

	p := Profile{Defaults: true}
	err := p.Check(ctx, a, b)
	t.Error(err)
}

func parse(t *testing.T, ctx *adt.OpContext, str string) *adt.Vertex {
	t.Helper()

	file, err := parser.ParseFile("subsume", str)
	if err != nil {
		t.Fatal(err)
	}

	root, errs := compile.Files(nil, ctx, "", file)
	if errs != nil {
		t.Fatal(errs)
	}

	root.Finalize(ctx)

	return root
}
