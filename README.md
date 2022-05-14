git-monorepo
============

Make Git monorepos like a boss.

## Features

- Preserve commit history for each repo
- Preserve Directory structure of each repo
- Ordering Commits chronologically

## Usage

### Create monorepo locally

```shell
$ git-monorepo init --sources gh:zikani03/articulated,gh:zikani03/pakadali --target toy-projects
```

Or with full git urls:

```shell
$ git-monorepo init --sources https://github.com/nndi-oss/ussdproxy,gh:nndi-oss/ussd-whois,https://github.com/nndi-oss/dialoguss --target ussd-projects
```

## IDEAS / TODO

- Make Submodules instead of mangling history
- Parallel clones
- Create monorepo locally and push to github

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
