Usage
=====

> NOTE: It's may be very CPU intensive for large repos at the moment, I'm gonna look for ways to reduce CPU usage. May also be memory intensive if you use the `--in-memory` flag. I also haven't tried it on very large repos (i.e. repos with very many many commits)

### Create monorepo locally

```shell
$ git-monorepo init --sources gh:zikani03/articulated,gh:zikani03/pakadali --target toy-projects
```

Or with full git urls:

```shell
$ git-monorepo init --sources https://github.com/nndi-oss/ussdproxy,gh:nndi-oss/ussd-whois,https://github.com/nndi-oss/dialoguss --target ussd-projects
```


# Command-line options

```text
Usage: git-monopore init

Initialize a monorepo from source repos

Flags:
  -h, --help                      Show context-sensitive help.
      --config="monopore.toml"    Location of configuration file
      --debug                     Enable debug mode
      --version                   Show version and quit

      --daemonize                 Daemonize or run in foreground
      --mangle                    Combine files from repos in one directory (not
                                  recommended!)
      --preserve-history          Preserve history from the repos
      --make-submodules           Add child repositories as submodules (not
                                  ideal!)
      --target                    The target directory to create repo in. Must
                                  not exist
      --sources=SOURCES,...       Source repositories with support for
                                  'git-down' shortcuts
```
