package saddle

// A File represents how `saddle` generates files at a path encoded in a file format.
#File: {
	path:    string
	format:  "JSON" | "YAML"
	content: string
}

// Format defines available encoding options `saddle` can generate files for
#Format: {
	JSON: "JSON"
	YAML: "YAML"
}

// Root is the current repository root. This is an opaque value that `saddle`
// will supply at runtime.
#Root: *"" | string
