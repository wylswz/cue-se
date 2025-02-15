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

package adt_test

import (
	"flag"
	"fmt"
	"strings"
	"testing"

	"golang.org/x/tools/txtar"

	"github.com/wylswz/cue-se/cue"
	"github.com/wylswz/cue-se/cue/cuecontext"
	"github.com/wylswz/cue-se/cue/errors"
	"github.com/wylswz/cue-se/internal/core/adt"
	"github.com/wylswz/cue-se/internal/core/debug"
	"github.com/wylswz/cue-se/internal/core/eval"
	"github.com/wylswz/cue-se/internal/core/runtime"
	"github.com/wylswz/cue-se/internal/core/validate"
	"github.com/wylswz/cue-se/internal/cuetxtar"
	_ "github.com/wylswz/cue-se/pkg"
)

var (
	todo = flag.Bool("todo", false, "run tests marked with #todo-compile")
)

// TestEval tests the default implementation of the evaluator.
func TestEval(t *testing.T) {
	test := cuetxtar.TxTarTest{
		Root: "../../../cue/testdata",
		Name: "eval",
		Skip: alwaysSkip,
		ToDo: needFix,
	}

	if *todo {
		test.ToDo = nil
	}

	test.Run(t, func(tc *cuetxtar.Test) {
		runEvalTest(tc, adt.DefaultVersion)
	})
}

var alwaysSkip = map[string]string{
	"compile/erralias": "compile error",
}

var needFix = map[string]string{
	"DIR/NAME": "reason",
}

var todoAlpha = map[string]string{
	"DIR/NAME": "reason",
}

func TestEvalAlpha(t *testing.T) {
	test := cuetxtar.TxTarTest{
		Root:     "../../../cue/testdata",
		Name:     "evalalpha",
		Fallback: "eval", // Allow eval golden files to pass these tests.
		Skip:     alwaysSkip,
		ToDo:     todoAlpha,
	}

	if *todo {
		test.ToDo = nil
	}

	test.Run(t, func(t *cuetxtar.Test) {
		runEvalTest(t, adt.DevVersion)
	})
}

func runEvalTest(t *cuetxtar.Test, version adt.EvaluatorVersion) {
	a := t.Instance()
	r := runtime.New()

	v, err := r.Build(nil, a)
	if err != nil {
		t.WriteErrors(err)
		return
	}

	e := eval.New(r)
	ctx := e.NewContext(v)
	ctx.Version = version
	v.Finalize(ctx)

	stats := ctx.Stats()
	w := t.Writer("stats")
	fmt.Fprintln(w, stats)
	// if n := stats.Leaks(); n > 0 {
	// 	t.Skipf("%d leaks reported", n)
	// }

	if b := validate.Validate(ctx, v, &validate.Config{
		AllErrors: true,
	}); b != nil {
		fmt.Fprintln(t, "Errors:")
		t.WriteErrors(b.Err)
		fmt.Fprintln(t, "")
		fmt.Fprintln(t, "Result:")
	}

	if v == nil {
		return
	}

	debug.WriteNode(t, r, v, &debug.Config{Cwd: t.Dir})
	fmt.Fprintln(t)
}

// TestX is for debugging. Do not delete.
func TestX(t *testing.T) {
	var verbosity int
	verbosity = 1 // comment to turn logging off.

	var version adt.EvaluatorVersion
	version = adt.DevVersion // comment to use default implementation.

	in := `
-- cue.mod/module.cue --
module: "mod.test"

-- in.cue --
	`

	if strings.HasSuffix(strings.TrimSpace(in), ".cue --") {
		t.Skip()
	}

	a := txtar.Parse([]byte(in))
	instance := cuetxtar.Load(a, t.TempDir())[0]
	if instance.Err != nil {
		t.Fatal(instance.Err)
	}

	r := runtime.New()

	v, err := r.Build(nil, instance)
	if err != nil {
		t.Fatal(err)
	}

	// t.Error(debug.NodeString(r, v, nil))
	// eval.Debug = true
	adt.Verbosity = verbosity
	t.Cleanup(func() { adt.Verbosity = 0 })

	e := eval.New(r)
	ctx := e.NewContext(v)
	ctx.Version = version
	v.Finalize(ctx)
	adt.Verbosity = 0

	if b := validate.Validate(ctx, v, &validate.Config{
		AllErrors: true,
	}); b != nil {
		t.Log(errors.Details(b.Err, nil))
	}

	t.Error(debug.NodeString(r, v, nil))

	t.Log(ctx.Stats())
}

func BenchmarkUnifyAPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ctx := cuecontext.New()
		v := ctx.CompileString("")
		for j := 0; j < 500; j++ {
			if j == 400 {
				b.StartTimer()
			}
			v = v.FillPath(cue.ParsePath(fmt.Sprintf("i_%d", i)), i)
		}
	}
}

func TestIssue2293(t *testing.T) {
	ctx := cuecontext.New()
	c := `a: {}, a`
	v1 := ctx.CompileString(c)
	v2 := ctx.CompileString(c)

	v1.Unify(v2)
}
