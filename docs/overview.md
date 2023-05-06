# About Saddle

Saddle is a tool to scaffold and manage configuration for code repositories. Saddle uses **configuration libraries** and **generators** to quickly add definitions for common developers tools in the **saddle store** then write the configuration to file paths various tools expect (eg. `.pretttierrc`, `.github/ci.yaml`, and `tsconfig.json`).

[CUE](https://cuelang.org/) drives much of `saddle`'s behavior as the configuration language and the document will reference CUE features in the design.

## Saddle Components

Saddle is composed of some logical abstractions:

- **Store** - A CUE package (conventionally `.saddle/`) that stores configuration definitions and mappings to the repository file structure.
<!-- - **Resource** - A CUE definition with a hidden `_resource` field that `saddle` will search for in the store. Resources drive actions like mapping configuration to files. -->
- **Generator** - A manifest that tells saddle how to generate and wire up configuration in the `Store`.

### Store

A package is a sensible and well understood method for organizing code. For `saddle`, we treat a package as something closer to a database that we actively read and write to for the state of the repository. This distinction matters for multiple reasons:

1. With a package, we can define a complete configuration across multiple files instead of designating fixed paths. A common issue is that data such as the application name can be repeated across various configurations and multiple formats with limited reusability. We can consolidate core facts of an application independently of what tools demand and create novel ways of working with configuration data.

2. In the future, we may need manage configuration in a package used by multiple tools. As CUE becomes an attractive choice for configuration, there can be relevant data to multiple tools in a single CUE package. We can get ahead of this by querying CUE packages for only data relevant to our functionality.

3. Generators actively modify package state by adding new definitions and fields, and we want to do this in a way that does not break user defined configuration and customizations. Think adding a new row in a database rather than drilling down and updating a data structure.

Command line tools generally approach configuration as a file to parse for initializing state. `saddle` changes the frame, willing to approach the repository as state that we can actively query and add augment with new data. In the future, we expect to analyze multiple sources internal and external to the codebase to optimize project tooling.

Currently [querying](https://cuelang.org/docs/usecases/query/) a package will be a linear graph walk and search, but as the CUE team improves query functionality and code generation we expect faster and novel ways to work with CUE code as a data store.

### Generator

`Generators` are the evolution of traditional code scaffolding. The usual flow for code generators is to generate source files (optionally customized with text templating) and leave further changes to developers. Bootstrapping repositories is useful for getting started, but leaves much to be desired in helping developers managing configuration over the application lifetime, especially as tools add features or change best practices.

`Generators` introduce step change improvements in granularity and long term utility. Instead of a single generator that offers little customization, `saddle` generators are instructions to make smaller transformations in a repository in a way that can be layered and only include what is needed. The `Store` is a recording of configuration data and how it maps to the repository, so developers know where files come from and the source of the recommendations.

Granularity is a critical feature for code scaffolding. The number of tools and platforms a codebase needs to integrate with is growing over time, and supporting templates that cover the matrix of options between languages (Go, Javascript, Python), tools (build systems, linters, code formatters), and platforms (cloud infrastructure, code repositories, package managers) is impractical. `saddle` takes the approach of targeting "features" that can detected from existing configuration or added as needed.

## The Saddle Stays On

We expect `saddle` to be a tool developers will use for the long run. Configuring the universe of tooling needed today for effective productivity can be overwhelming. `saddle` is opportunity to socialize best practices through configuration libraries, discover new tools that increase developer happiness, and provide the fastest experience working with cutting edge software development practices.
