# gover-js [![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/github.com/vanillaiice/gover-js) [![Go Report Card](https://goreportcard.com/badge/github.com/vanillaiice/gover-js)](https://goreportcard.com/report/github.com/vanillaiice/gover-js)

> This project is forked from [gover](https://github.com/vanillaiice/gover), which is a package version management tool for Go projects.

gover-js is package version management tool for JavaScript projects.

Instead of manually incrementing the version number in your code (🗿),
you can simply use `gover-js` to automatically do it.
Also, you can use `gover-js` to tag the git branch to the current version of your project (using `git tag`).

Under the hood, `gover-js` will read your `package.json` file in your project
and update the version number accordingly.

# Installation

```sh
$ go install github.com/vanillaiice/gover-js@latest
```

# Usage

First, initialize your js project with your favorite package manager:

```sh
# bun
$ bun init
# npm
$ npm init
```

Then, you can simply increment the version string in the `package.json` file using the `gover-js bump` command:

```
# bump to major version (e.g. v1.0.0 -> v2.0.0)
$ gover-js bump --major
# bump to minor version (e.g. v1.0.0 -> v1.1.0) with verbose log from custom package.json file
$ gover-js -V bump --minor -f pkg.json
```
# Help

```sh
NAME:
   gover-js - package version management tool for JavaScript projects

USAGE:
   gover-js [global options] command [command options]

VERSION:
   0.1.0

AUTHOR:
   vanillaiice <vanillaiice1@proton.me>

COMMANDS:
   bump, b  bump version
   tag, t   tag git branch with the current version
   get, e   get the current version
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbose, -V  show verbose log (default: false)
   --help, -h     show help
   --version, -v  print the version
```

# License

GPLv3

# Author

[vanillaiice](https://github.com/vanillaiice)
