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

func walker(val cue.Value) bool {
	targetPath := cue.MakePath(cue.Hid("_y", "banana.com/p"))

	fmt.Println(val.Path(), "-", val.Kind())

	if val.Kind() == cue.StructKind {
		if c := val.LookupPath(targetPath); c.Exists() {
			fmt.Printf("we found _y: %v\n", c)
		} else {
			fmt.Println("we did not find _y")
		}
	}

	return true
}

func main() {
	const file = `
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
	bps := loadTxTar(file)

	v := ctx.BuildInstance(bps[0]).LookupPath(cue.ParsePath(""))
	fmt.Println("Instance -", bps[0].ID())
	fmt.Println("Concrete?", v.IsConcrete())
	fmt.Println("Error?", v.Validate())
	v.Walk(walker, nil)
}
