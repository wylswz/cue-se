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
	_ "github.com/wylswz/cue-se/pkg/crypto/ed25519"
	_ "github.com/wylswz/cue-se/pkg/crypto/hmac"
	_ "github.com/wylswz/cue-se/pkg/crypto/md5"
	_ "github.com/wylswz/cue-se/pkg/crypto/sha1"
	_ "github.com/wylswz/cue-se/pkg/crypto/sha256"
	_ "github.com/wylswz/cue-se/pkg/crypto/sha512"
	_ "github.com/wylswz/cue-se/pkg/encoding/base64"
	_ "github.com/wylswz/cue-se/pkg/encoding/csv"
	_ "github.com/wylswz/cue-se/pkg/encoding/hex"
	_ "github.com/wylswz/cue-se/pkg/encoding/json"
	_ "github.com/wylswz/cue-se/pkg/encoding/yaml"
	_ "github.com/wylswz/cue-se/pkg/html"

	_ "github.com/wylswz/cue-se/pkg/list"
	_ "github.com/wylswz/cue-se/pkg/math"
	_ "github.com/wylswz/cue-se/pkg/math/bits"
	_ "github.com/wylswz/cue-se/pkg/net"
	_ "github.com/wylswz/cue-se/pkg/path"
	_ "github.com/wylswz/cue-se/pkg/regexp"
	_ "github.com/wylswz/cue-se/pkg/strconv"
	_ "github.com/wylswz/cue-se/pkg/strings"
	_ "github.com/wylswz/cue-se/pkg/struct"
	_ "github.com/wylswz/cue-se/pkg/text/tabwriter"
	_ "github.com/wylswz/cue-se/pkg/text/template"
	_ "github.com/wylswz/cue-se/pkg/time"
	_ "github.com/wylswz/cue-se/pkg/tool"
	_ "github.com/wylswz/cue-se/pkg/tool/cli"
	_ "github.com/wylswz/cue-se/pkg/tool/exec"
	_ "github.com/wylswz/cue-se/pkg/tool/file"
	_ "github.com/wylswz/cue-se/pkg/tool/http"
	_ "github.com/wylswz/cue-se/pkg/tool/os"
	_ "github.com/wylswz/cue-se/pkg/uuid"
)
