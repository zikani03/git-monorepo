git-monorepo
============

> Make Git monorepos like a boss.

`git-monorepo` is a tool that allows you to combine two or more Git repositories into one repository and preserves the commit history for each repo as well as the directory structure.

## Features

- Preserves commit history for each repo
- Preserves Directory structure of each repo
- Orders Commits chronologically

## Building

git-monorepo requires atleast Go 1.18 to build. 

```sh
$ git clone https://github.com/zikani03/git-monorepo
$ go build -o git-monorepo ./cmd/main.go
```

## Usage

> NOTE: It's currently very CPU intensive at the moment, I'm gonna look for ways to reduce CPU usage. May also be memory intensive if you use the `--in-memory` flag. I also haven't tried it on very large repos (i.e. repos with very many many commits)

### Create monorepo locally

```shell
$ git-monorepo init --sources gh:zikani03/articulated,gh:zikani03/pakadali --target toy-projects
```

Or with full git urls:

```shell
$ git-monorepo init --sources https://github.com/nndi-oss/ussdproxy,gh:nndi-oss/ussd-whois,https://github.com/nndi-oss/dialoguss --target ussd-projects
```

## IDEAS / TODO

- Support "squashing-merging" - that is merge the repos without preserving fine-grained commit history

- Make Git Submodules instead of mangling history

- Clone repositories in parallel

- Create monorepo locally and push to github (and other git hosting services)

```shell
$ git monorepo init --preserve-history \ 
  --sources gh:zikani03/articulated,gh:zikani03/pakadali \ 
  --target gh:zikani03/toy-projects \ 
  --github-token=$GH_TOKEN
```

## CREDITS

- [Go-Git](https://pkg.go.dev/github.com/go-git/go-git/v5)


## References

- [Atlassian Monorepos tutorial](https://www.atlassian.com/git/tutorials/monorepos)
- [Git SCM - Appendix B: Embedding Git in your Applications - go-git](https://git-scm.com/book/en/v2/Appendix-B%3A-Embedding-Git-in-your-Applications-go-git)
