Usage
=====

> NOTE: It's may be very CPU intensive for large repos at the moment, I'm gonna look for ways to reduce CPU usage. May also be memory intensive if you use the `--in-memory` flag. I also haven't tried it on very large repos (i.e. repos with very many many commits)

### Create a new monorepo with `git-monorepo init` 

```shell
$ git-monorepo init --sources gh:zikani03/articulated,gh:zikani03/pakadali --target toy-projects
```

Or with full git urls:

```shell
$ git-monorepo init --sources https://github.com/nndi-oss/ussdproxy,gh:nndi-oss/ussd-whois,https://github.com/nndi-oss/dialoguss --target ussd-projects
```

### Merge into existing repo with `git-monorepo merge` 

```shell
$ git-monorepo merge --sources gh:nndi-oss/ussd-whois,gh:nndi-oss/ussdproxy --target ./existing-repo
```

Or with full git urls:

```shell
$ git-monorepo merge --sources https://github.com/nndi-oss/ussdproxy,\
https://github.com/nndi-oss/ussd-whois, \
https://github.com/nndi-oss/dialoguss \
--target ./existing-repo
```
