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

package runtime

import (
	"github.com/wylswz/cue-se/cue/ast"
	"github.com/wylswz/cue-se/cue/errors"
	"github.com/wylswz/cue-se/cue/token"
)

var _ errors.Error = &nodeError{}

// A nodeError is an error associated with processing an AST node.
type nodeError struct {
	path []string // optional
	n    ast.Node

	errors.Message
}

func (n *nodeError) Error() string {
	return errors.String(n)
}

func nodeErrorf(n ast.Node, format string, args ...interface{}) *nodeError {
	return &nodeError{
		n:       n,
		Message: errors.NewMessagef(format, args...),
	}
}

func (e *nodeError) Position() token.Pos {
	return e.n.Pos()
}

func (e *nodeError) InputPositions() []token.Pos { return nil }

func (e *nodeError) Path() []string {
	return e.path
}
