# Boilerplate Project

just simple `echo` app `http://localhost:3000/foo`

## Source code

**Our source code must be in src/<domain\>/<user\>/<project\> directory**. 
As in this project it available in **boiler.plate/echo/foobar**

## Running

`go run src/boiler.plate/echo/foobar/main.go`

## Project Setup

This golang project use some tool to get better environment for golang project. even this small project.

As golang `$GOPATH` mechanism, this setup chosen to have more isolation for each golang project

### direnv

`direnv` is an environment switcher for the shell. It knows how to hook into bash, zsh, tcsh, fish shell and elvish to load or unload environment variables depending on the current directory. This allows project-specific environment variables without cluttering the `~/.profile` file. 

*references:*

- [http://rachbelaid.com/handling-go-workspace-with-direnv/](http://rachbelaid.com/handling-go-workspace-with-direnv/)
- [https://direnv.net/](https://direnv.net/)

**direnv used to have GOPATH changed automatically for this project only**, but leave PATH not touched.

### dep

Dependency management for Go [https://golang.github.io/dep/](https://golang.github.io/dep/)

```
dep is a tool for managing dependencies for Go projects

Usage: dep <command>

Commands:

  init     Initialize a new project with manifest and lock files
  status   Report the status of the project's dependencies
  ensure   Ensure a dependency is safely vendored in the project
  prune    Prune the vendor tree of unused packages
  version  Show the dep version information

Examples:
  dep init                               set up a new project
  dep ensure                             install the project's dependencies
  dep ensure -update                     update the locked versions of all dependencies
  dep ensure -add github.com/pkg/errors  add a dependency to the project
```

### to-do

- make sure `go`, `dep`, `direnv` installed
- go to `src/boiler.plate/echo/foobar/main.go`
- then `go ensure` to setup