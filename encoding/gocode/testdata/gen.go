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

package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/wylswz/cue-se/cue"
	"github.com/wylswz/cue-se/cue/load"
	"github.com/wylswz/cue-se/encoding/gocode"
	_ "github.com/wylswz/cue-se/pkg"
)

func main() {
	dirs, err := ioutil.ReadDir("testdata")
	if err != nil {
		log.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		dir := filepath.Join(cwd, "testdata")
		pkg := "." + string(filepath.Separator) + d.Name()
		inst := cue.Build(load.Instances([]string{pkg}, &load.Config{
			Dir:        dir,
			ModuleRoot: dir,
			Module:     "github.com/wylswz/cue-se/encoding/gocode/testdata",
		}))[0]
		if err := inst.Err; err != nil {
			log.Fatal(err)
		}

		goPkg := "./testdata/" + d.Name()
		b, err := gocode.Generate(goPkg, inst, nil)
		if err != nil {
			log.Fatal(err)
		}

		goFile := filepath.Join("testdata", d.Name(), "cue_gen.go")
		if err := os.WriteFile(goFile, b, 0644); err != nil {
			log.Fatal(err)
		}
	}
}
