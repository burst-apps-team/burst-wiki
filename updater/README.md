# Site Updater

This script, written in Go, updates the site.

It can be compiled using `go build updater.go`

**It must be run outside the wiki repository, otherwise it will try to delete itself and fail.**

Example invocation, from Bash, **with working directory set to the root of the repository**:
```shell script
go build updader/updater.go -o ../updater
../updater
```

In order to run the updater, you need `mkdocs` and the `material` theme installed. To install them, you need to have Python installed. MkDocs claims to support Python 2 and 3.

```shell script
pip install mkdocs
pip install mkdocs-material
```

The updater will automatically build the site using MkDocs, copy it into the `gh-pages` branch, apply the redirect patch that fixes old hardlinks, and push to GitHub. So, in order to run this, you also need push permissions to the `burst-wiki` repository.
