package load_test

import (
	"fmt"
	"io/fs"
	"testing"

	"cuelabs.dev/go/oci/ociregistry/ociclient"
	"github.com/wylswz/cue-se/cue/cuecontext"
	"github.com/wylswz/cue-se/cue/errors"
	"github.com/wylswz/cue-se/internal/cuetxtar"
	"github.com/wylswz/cue-se/internal/registrytest"
)

func TestModuleFetch(t *testing.T) {
	test := cuetxtar.TxTarTest{
		Root: "./testdata/testfetch",
		Name: "modfetch",
	}
	test.Run(t, func(t *cuetxtar.Test) {
		rfs, err := fs.Sub(registrytest.TxtarFS(t.Archive), "_registry")
		if err != nil {
			t.Fatal(err)
		}
		r, err := registrytest.New(rfs, "")
		if err != nil {
			t.Fatal(err)
		}
		defer r.Close()
		reg, err := ociclient.New(r.Host(), &ociclient.Options{
			Insecure: true,
		})
		if err != nil {
			t.Fatal(err)
		}
		t.LoadConfig.Registry = reg
		ctx := cuecontext.New()
		insts := t.RawInstances()
		if len(insts) != 1 {
			t.Fatalf("wrong instance count; got %d want 1", len(insts))
		}
		inst := insts[0]
		if inst.Err != nil {
			errors.Print(t.Writer("error"), inst.Err, &errors.Config{
				ToSlash: true,
			})
			return
		}
		v := ctx.BuildInstance(inst)
		if err := v.Validate(); err != nil {
			t.Fatal(err)
		}
		fmt.Fprintf(t, "%v\n", v)
	})
}
