git-monorepo
============

> Make Git monorepos like a boss.

`git-monorepo` is a tool that allows you to combine two or more Git repositories into one repository and preserves the commit history for each repo as well as the directory structure.

## Features

- Preserves commit history for each repo
- Preserves Directory structure of each repo
- Orders Commits chronologically

## Example

Combine two repos into a new `toy-projects` repo using `gh:` short-cut for GitHub.

```shell
$ git-monorepo init --target toy-projects \
   --sources gh:zikani03/articulated,gh:zikani03/pakadali 
```

Or with full git urls:

```shell
$ git-monorepo init  --target ussd-projects \
   --sources https://bitbucket.org/nndi/phada,git@github.com:nndi-oss/ussdproxy,gh:nndi-oss/ussd-whois7
```

## CREDITS

- [Go-Git](https://pkg.go.dev/github.com/go-git/go-git/v5)

