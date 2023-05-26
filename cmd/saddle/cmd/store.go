package cmd

import (
	"errors"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
)

const (
	// defaultPkg is the default path to look for as a saddle store
	defaultPkg = ".saddle/"
)

type store struct {
	// pkg is the path to the CUE package representing the store
	pkg string

	// root is the top level CUE value of the package
	root cue.Value

	runtime *cue.Context
}

func newStore(pkg string, runtime *cue.Context) *store {
	cfg := &load.Config{
		Dir: pkg,
	}
	val := runtime.BuildInstance(load.Instances(nil, cfg)[0])
	return &store{
		pkg:     pkg,
		root:    val,
		runtime: runtime,
	}
}

func (st *store) manifest() (cue.Value, error) {
	target := cue.MakePath(cue.Str("manifest"))

	if c := st.root.LookupPath(target); c.Exists() {
		return c, nil
	}

	return cue.Value{}, errors.New("No 'manifest' in top level")
}

func (st *store) collect(name, pkg string) []cue.Value {
	var all []cue.Value
	target := cue.MakePath(cue.Hid(name, pkg))

	walker := func(v cue.Value) bool {
		if v.Kind() == cue.StructKind {
			if c := v.LookupPath(target); c.Exists() {
				// We want the enclosing struct
				all = append(all, v)
			}
		}

		return true
	}

	// Walk only evaluates emitted fields
	st.root.Walk(walker, nil)

	return all
}
