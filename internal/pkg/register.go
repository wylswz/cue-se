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

package pkg

import (
	"github.com/wylswz/cue-se/cue/errors"
	"github.com/wylswz/cue-se/internal/core/adt"
	"github.com/wylswz/cue-se/internal/core/eval"
	"github.com/wylswz/cue-se/internal/core/runtime"
)

func Register(importPath string, p *Package) {
	f := func(r adt.Runtime) (*adt.Vertex, errors.Error) {
		ctx := eval.NewContext(r, nil)

		return p.MustCompile(ctx, importPath), nil
	}
	runtime.RegisterBuiltin(importPath, f)
}
