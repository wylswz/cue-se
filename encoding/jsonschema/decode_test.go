// Copyright 2019 CUE Authors
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

package jsonschema

import (
	"bytes"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-quicktest/qt"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/txtar"

	"github.com/wylswz/cue-se/cue"
	"github.com/wylswz/cue-se/cue/ast"
	"github.com/wylswz/cue-se/cue/errors"
	"github.com/wylswz/cue-se/cue/format"
	"github.com/wylswz/cue-se/cue/token"
	"github.com/wylswz/cue-se/encoding/json"
	"github.com/wylswz/cue-se/encoding/yaml"
	"github.com/wylswz/cue-se/internal/astinternal"
	"github.com/wylswz/cue-se/internal/cuetest"
	_ "github.com/wylswz/cue-se/pkg"
)

// TestDecode reads the testdata/*.txtar files, converts the contained
// JSON schema to CUE and compares it against the output.
//
// Set CUE_UPDATE=1 to update test files with the corresponding output.
func TestDecode(t *testing.T) {
	err := filepath.WalkDir("testdata", func(fullpath string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !strings.HasSuffix(fullpath, ".txtar") {
			return nil
		}

		t.Run(fullpath, func(t *testing.T) {
			a, err := txtar.ParseFile(fullpath)
			if err != nil {
				t.Fatal(err)
			}

			cfg := &Config{ID: fullpath}

			if bytes.Contains(a.Comment, []byte("openapi")) {
				cfg.Root = "#/components/schemas/"
				cfg.Map = func(p token.Pos, a []string) ([]ast.Label, error) {
					// Just for testing: does not validate the path.
					return []ast.Label{ast.NewIdent("#" + a[len(a)-1])}, nil
				}
			}

			r := &cue.Runtime{}
			var in *cue.Instance
			var out, errout []byte
			outIndex := -1
			errIndex := -1

			for i, f := range a.Files {
				switch path.Ext(f.Name) {
				case ".json":
					in, err = json.Decode(r, f.Name, f.Data)
				case ".yaml":
					in, err = yaml.Decode(r, f.Name, f.Data)
				case ".cue":
					out = f.Data
					outIndex = i
				case ".err":
					errout = f.Data
					errIndex = i
				}
			}
			if err != nil {
				t.Fatal(err)
			}

			updated := false

			expr, err := Extract(in, cfg)
			if err != nil {
				got := []byte(errors.Details(err, nil))

				got = bytes.TrimSpace(got)
				errout = bytes.TrimSpace(errout)

				switch {
				case !cmp.Equal(errout, got):
					if cuetest.UpdateGoldenFiles {
						a.Files[errIndex].Data = got
						updated = true
						break
					}
					t.Error(cmp.Diff(string(got), string(errout)))
				}
			}

			if expr != nil {
				b, err := format.Node(expr, format.Simplify())
				if err != nil {
					t.Fatal(errors.Details(err, nil))
				}

				// verify the generated CUE.
				if !bytes.Contains(a.Comment, []byte("#noverify")) {
					if _, err = r.Compile(fullpath, b); err != nil {
						t.Fatal(errors.Details(err, nil))
					}
				}

				b = bytes.TrimSpace(b)
				out = bytes.TrimSpace(out)

				switch {
				case !cmp.Equal(b, out):
					if cuetest.UpdateGoldenFiles {
						updated = true
						a.Files[outIndex].Data = b
						break
					}
					t.Error(cmp.Diff(string(out), string(b)))
				}
			}

			if updated {
				b := txtar.Format(a)
				err = os.WriteFile(fullpath, b, 0644)
				if err != nil {
					t.Fatal(err)
				}
			}
		})
		return nil
	})
	qt.Assert(t, qt.IsNil(err))
}

func TestX(t *testing.T) {
	t.Skip()
	data := `
-- schema.json --
`

	a := txtar.Parse([]byte(data))

	r := &cue.Runtime{}
	var in *cue.Instance
	var err error
	for _, f := range a.Files {
		switch path.Ext(f.Name) {
		case ".json":
			in, err = json.Decode(r, f.Name, f.Data)
			if err != nil {
				t.Fatal(err)
			}
		case ".yaml":
			in, err = yaml.Decode(r, f.Name, f.Data)
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	cfg := &Config{ID: "test"}
	expr, err := Extract(in, cfg)
	if err != nil {
		t.Fatal(err)
	}

	t.Fatal(astinternal.DebugStr(expr))
}
