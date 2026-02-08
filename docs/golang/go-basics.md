
# Package
- a directory of .go files
- produces one compiled package

# Module
- a collection of packages with a go.mod at the root.

# Workspace (multi-module development)
- Used when you split your project (repo) into multiple modules
```
repo/
  go.work
  api/go.mod
  node-agent/go.mod
  libs/go.mod
```

- initilize workspace
```
go work init ./api ./node-agent ./libs
```

- add a module
```
go work use ./newmodule
```

- A workspace does not “merge modules into one module” for publishing/versioning. Each module still has its own go.mod and versioning story.
- typically for monorepos or multi-repo local development

# Package and module relationship
- A module contains many packages.
- A package belongs to exactly one module (the nearest go.mod “above” it).

# Lexical tokens 
- the smallest meaningful units a language compiler reads from source code 
- Tools like formatters, linters, and syntax highlighters start from tokens.
- not much related to current topic