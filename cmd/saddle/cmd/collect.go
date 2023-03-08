package cmd

import "cuelang.org/go/cue"

// Collect aggregates structs matching resource definitions
func collect(val cue.Value, name, pkg string) []cue.Value {
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
	val.Walk(walker, nil)

	return all
}
