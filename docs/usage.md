Usage
=====

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
