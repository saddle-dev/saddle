package main

import (
	"fmt"
	"io/ioutil"
	"os"

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

func main() {
	const file = `
-- cue.mod/module.cue --
module: "example.com/test"

-- x.cue --
package x

import "banana.com/p"

x: p.#Def & {
	y: 4
}

-- cue.mod/pkg/banana.com/p/p.cue --
package p

#Def: {
	_y: 4
}
`
	ctx := cuecontext.New()
	bps := loadTxTar(file)

	v := ctx.BuildInstance(bps[0]).LookupPath(cue.ParsePath("x"))

	targetPath := cue.MakePath(cue.Hid("_y", "banana.com/p"))
	if c := v.LookupPath(targetPath); c.Exists() {
		fmt.Printf("we found it: %v\n", c)
	} else {
		fmt.Println("we did not find it")
	}
}
