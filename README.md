git-monorepo
============

Make Git monorepos like a boss.

## Features

- Preserve commit history
- Make Submodules instead of mangling history
- Parallel clones
- Preserve Directory structure
- Ordering Commits chronologically
- Ordering Commits by source repo 
- Lossy monorepo without full history of projects
- Lossy monorepo with a `<repo>.commit.history` file 

## Usage

### Create monorepo locally

```shell
$ git monorepo init --preserve-history \ 
  --sources gh:zikani03/articulated,gh:zikani03/pakadali \ 
  --target toy-projects
```

Or 

```shell
$ git monorepo init --preserve-history --github-username zikani03 --sources articulated,pakadali --target toy-projects
```

### Create monorepo locally and push to github

```shell
$ git monorepo init --preserve-history \ 
  --sources gh:zikani03/articulated,gh:zikani03/pakadali \ 
  --target gh:zikani03/toy-projects \ 
  --github-token=$GH_TOKEN
```

Or 

```shell
$ git monorepo init --preserve-history --github-username zikani03 --sources articulated,pakadali --target gh:toy-projects
```

## CREDITS

- [Go-Git](https://pkg.go.dev/github.com/go-git/go-git/v5)


## References

- [Atlassian Monorepos tutorial](https://www.atlassian.com/git/tutorials/monorepos)
- [Git SCM - Appendix B: Embedding Git in your Applications - go-git](https://git-scm.com/book/en/v2/Appendix-B%3A-Embedding-Git-in-your-Applications-go-git)
