# Saddle: Code scaffolding for the long haul

## Caveat

**THIS IS CURRENTLY AN EXPERIMENTAL, UNCOMPLETED PROJECT**

## What's Saddle?

`saddle` is a new kind of scaffolding tool.
Instead of simply copying files to bootstrap a project, we provide **code generators** and **configuration libraries** to define a repository's configuration files in a powerful language: [CUE](https://cuelang.org/).

The unified source of truth means that core facts about an application (eg. the application name) can be reused across configurations.
Instead of searching through various files and formats, write a fact down once with confidence that the value will stay consistent in every configuration it is used.

`saddle up` syncs actual configuration files in various formats (eg. YAML, JSON, TOML) that other tools expect.
Unlike other scaffolding tools, the _saddle stays on_ and continues to manage configurations.
With CUE, we can organize our configuration in packages, import libraries of good defaults defined by the community, and rewrite and rollout configurations in a language designed from the ground up to do so.

The store cleanly composes different toolchains.
Instead of creating template `github-nextjs-typescript-vercel` for a next.js app written in typescript on GitHub and hosted on Vercel or template `gitlab-react-netlify` for a `create-react-app` managed client-side application on GitLab and hosted on Netlify, we generate configuration definitions for tools (`.github/ci.yaml` or `tsconfig.json`) in a modular, incremental way.
Once we have a clear idea of what an ideal set of tools we want on every project those toolchains can be defined together with a base set of defaults and applied to multiple repositories if needed.

## Download and Install

**Release builds**

[Download]() the latest release from GitHub.

**Install from Source**

```shell
go install github.com/saddle-dev/saddle/cmd/saddle@latest
```

This will install the `saddle` command line tool.

## Options

```
$ saddle --help
saddle manages application configuration in a single language, CUE.
Users supply configuration data provided by libraries and validated against
schemas. saddle will then generate the files in the correct format.

saddle makes it simple to manage configuration for multiple tools and apply
them consistently across one to many repositories.

For more information on using saddle to manage code see saddle.dev.

Usage:
  saddle [command]

Available Commands:
  check         Validate configuration against schemas and diff against existing files
  generate      Scaffold repository configuration for tools
  help          Help about any command
  import        Detect existing configuration and generate CUE files
  up            Sync CUE definitions to application configurations
  version       Print saddle version

Use "saddle [command] --help" for more information about a command.
```

## Quickstart

Initialize a CUE module (for package management)

```shell
cue mod init
```

Generate CUE definitions from existing application files

```shell
saddle import ./ .saddle/
```

This will look for files like `.goreleaser.yaml`, `.prettierrc`, and `tsconfig.json` and generate CUE files with imports and configuration that match your existing file locations and data.

After making changes to generated CUE definitions, we can run `saddle up` to update the actual file paths each tool is looking for.

```shell
saddle up .saddle/
```

## Tutorial: Adding More Tools

Another option is to generate definitions for specific configurations. Let's create one for `.prettierrc` to format our code for us.

```shell
saddle generate github.com/saddle-dev/saddle/pkg/tool/prettier .saddle/
```

Let's look at the generated file `.saddle/prettier.cue`:

```CUE
package saddle

import (
    "github.com/saddle-dev/saddle"
    "github.com/saddle-dev/saddle/pkg/tool/prettier"
)

#Conf: prettier.#Config

#Path: saddle.#File
#Path: {
  // Saddle provides the repository root
  path:         "\(saddle.#Root)/.prettierrc"
  // Option to explicitly define the output format
  format:       saddle.#Format.JSON
  content:      #Conf
}
```

The configuration has two definitions: The configuration (the contents of the file) for `prettier` as well as a `saddle.#File` definition to tell `saddle` where to sync the file when we run `saddle up`. A basic definition for `prettier.#Config` looks like this:

```CUE
#Config: {
  trailingComma:    *"es5" | string
  tabWidth:         *4 | int
  semi:             *false | bool
  singleQuote:      *true | bool
}
```

In this case, `prettier` itself sets default values which is mirrored in the `prettier` Saddle package and can be used as is to create a valid `.prettierrc`. Our company's style guide requires semicolons in our Javascript code, so we need to modify the base configuration:

```CUE
#Conf: prettier.#Config
#Conf: {
  semi: true
}
```

When CUE sees repeated definition declarations, it will attempt to unify the values as long as the types can be merged. At this point when we run `saddle up` it will create `.prettierrc` in our repository root with these contents:

```JSON
{
  "trailingComma": "es5",
  "tabWidth": 4,
  "semi": true,
  "singleQuote": true
}
```

With `saddle`, we can easily create application configurations that are checked for correctness and with reasonable defaults so can focus on our code.

This only scratches surface of how we can manage codebases with `saddle`. Continuing with the company example, we can define libraries and generators for how we want developers to configure tools consistently across repositories. We can also leverage `saddle` to rollout changes by creating a "configuration repository" with per-application config packages (CUE will unify values) and using `saddle check` and `saddle up` to make the code changes.

```
# saddle up can take multiple packages as input, last path is target repo 'foo'
config-repo$ saddle up ./saddle foo-config ../foo
```
