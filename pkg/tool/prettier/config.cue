package prettier

// Config is a valid `.prettierrc` configuration
#Config: {
  trailingComma:    *"es5" | string
  tabWidth:         *4 | int
  semi:             *false | bool
  singleQuote:      *true | bool
}
