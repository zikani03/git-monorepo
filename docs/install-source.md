Installation from Source
========================

Use the following steps to build the binary from the source code.<br/>
`git-monorepo` doesn't have any other dependencies besides Go itself, so building it is
as simple as cloning it and running go build. 

You must have [Go](https://golang.org) installed

```bash
$ git clone https://github.com/zikani03/git-monorepo
$ cd git-monorepo
$ go build cmd/main.go -o git-monorepo
```

This will create a `git-monorepo` binary in the directory (`git-monorepo.exe` on Windows)
