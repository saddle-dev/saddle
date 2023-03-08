package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/build"
	"cuelang.org/go/cue/cuecontext"
	"golang.org/x/tools/txtar"

	"github.com/saddle-dev/saddle/internal/cuetxtar"
)

func loadTxTar(file string) []*build.Instance {
	dir, _ := ioutil.TempDir("", "*")
	defer os.RemoveAll(dir)

	insts := cuetxtar.Load(txtar.Parse([]byte(file)), dir)

	return insts
}

func TestCollect(t *testing.T) {
	conf := `
-- cue.mod/module.cue --
module: "example.com/test"

-- x.cue --
package x

import "banana.com/p"

foo: p.#Def & {
	y: 4
}

-- cue.mod/pkg/banana.com/p/p.cue --
package p

#Def: {
	_y: 4
	y: int
}
`
	ctx := cuecontext.New()
	bps := loadTxTar(conf)
	v := ctx.BuildInstance(bps[0]).LookupPath(cue.ParsePath(""))

	found := collect(v, "_y", "banana.com/p")
	if len(found) != 1 {
		t.Fatalf("expected number of found structs to be 1")
	}
}
