package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

func main() {
	ctx := cuecontext.New()
	bps := load.Instances([]string{"."}, nil)

	v := ctx.BuildInstance(bps[0]).LookupPath(cue.ParsePath("x"))

	targetPath := cue.MakePath(cue.Hid("_y", "banana.com/p"))
	if c := v.LookupPath(targetPath); c.Exists() {
		fmt.Printf("we found it: %v\n", c)
	} else {
		fmt.Println("we did not find it")
	}
}
