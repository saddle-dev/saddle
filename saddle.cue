package saddle

// A Manifest defines mappings between configurations and files
#Manifest: {
	paths: {[string]: #Config}
}

// A Config represents how `saddle` generates configurations at a path encoded in a Config format.
#Config: {
	_resource: null
	path:      string
	format:    #Format
	content:   string
}

// Format defines available encoding options `saddle` can generate files for
#Format: "JSON" | "YAML"

// Root is the current repository root. This is an opaque value that `saddle`
// will supply at runtime.
#Root: *"" | string
