package module_test

import (
	"github.com/wylswz/cue-se/internal/mod/module"
	"github.com/wylswz/cue-se/internal/mod/mvs"
)

var _ mvs.Versions[module.Version] = module.Versions{}
