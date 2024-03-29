package saddle

// A Manifest defines mappings between configurations and files
#Manifest: {
	files: {[string]: #File}
}

// A File represents how `saddle` generates files at a path encoded in a file format.
#File: {
	_resource: null
	path:      string
	format:    "JSON" | "YAML"
	content:   string
}

// Format defines available encoding options `saddle` can generate files for
#Format: {
	JSON: "JSON"
	YAML: "YAML"
}

// Root is the current repository root. This is an opaque value that `saddle`
// will supply at runtime.
#Root: *"" | string
