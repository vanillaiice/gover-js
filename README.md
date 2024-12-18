# gover-js [![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/github.com/vanillaiice/gover-js) [![Go Report Card](https://goreportcard.com/badge/github.com/vanillaiice/gover-js)](https://goreportcard.com/report/github.com/vanillaiice/gover-js)

> This project is forked from [gover](https://github.com/vanillaiice/gover), which is a package version management tool for Go projects.

gover-js is package version management tool for JavaScript projects.

Instead of manually incrementing the version number in your code like a - 🗿,
you can simply use `gover-js` to automatically do it.
Also, you can use `gover-js` to commit changes to your `package.json` file,
and tag branches.

Under the hood, `gover-js` reads the `package.json` file in project root
and updates the version number accordingly.

# Installation

```sh
$ go install github.com/vanillaiice/gover-js@latest
```

# Usage

First, initialize your js project with your favorite package manager:

```sh
# using bun
$ bun init
# npm
$ npm init
```

## `bump`

You can increment the version string in the `package.json` file using the `bump` command:

```sh
# bump to major version (e.g. v1.0.0 -> v2.0.0)
$ gover-js bump --major
# bump to minor version (e.g. v1.0.0 -> v1.1.0) with verbose log from custom package.json file
$ gover-js -V -f pkg.json bump --minor
```

## `commit`

You can commit the `package.json` file using the `commit` command:

```sh
# commit using default git template
$ gover-js commit
# commit with custom template and custom package.json file
$ gover-js -f pkg.json commit --command "git commit {{ .File }} -m 'bump to {{ .Version }}'"
```

## `tag`

You can tag the current branch using the `tag` command:

```sh
# tag using default git command ("git tag {{ .Version }}")
$ gover-js tag
```

## `get`

The `get` commands returns the current version of the package:

```sh
$ gover-js get
# with custom file
$ gover-js -f pkg.json get
```

> You can also set some arguments with environment variables:

> - `package.json` file: VERSION_FILE
> - commit command: COMMIT_COMMAND
> - tag command: TAG_COMMAND

# Help

```sh
NAME:
   gover-js - package version management tool for JavaScript projects

USAGE:
   gover-js [global options] command [command options]

VERSION:
   0.2.0

AUTHOR:
   vanillaiice <vanillaiice1@proton.me>

COMMANDS:
   get, e     get the current version
   bump, b    bump version
   commit, c  commit version
   tag, t     tag branch with the current version
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file FILE, -f FILE  load version from FILE (default: "package.json") [$VERSION_FILE]
   --verbose, -V         show verbose log (default: false)
   --help, -h            show help
   --version, -v         print the version
```

# License

GPLv3

# Author

[vanillaiice](https://github.com/vanillaiice)
