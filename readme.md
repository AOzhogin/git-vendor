git-vendor
==========
A git command for managing vendored dependencies.

`git-vendor` is a wrapper around `git-subtree` commands for checking out and updating vendored dependencies.

By default `git-vendor` conforms to the pattern used for vendoring golang dependencies:

* Dependencies are stored under `git-vendor/` directory in the repo.
* Dependencies are stored under the fully qualified project path.
    * e.g. `https://github.com/a.ozhogin/example` will be stored under `git-vendor/github.com/a.ozhogin/example`.

## All initial logic is taken from the project and reprogrammed in Go - with saving changes to file

  https://github.com/brettlangdon/git-vendor

## Usage

`git-vendor` provides the following commands:

* `git-vendor help` - help.
* `git-vendor completion` - generate completion for bash, fish, powershell, zsh.
* `git-vendor add <name> <repository> [<ref>]` - add a new vendored dependency.
* `git-vendor list [<name>]` - list current vendored dependencies, their source, and current vendored ref.
* `git-vendor update <name> [<ref>]` - update a vendored dependency.
* `git-vendor remove <name>` - remove vendor.
* `git-vendor upstream <name> [<ref>] [<repository>]` - share with the upstream vendored dependency.

## Installation
Manually:

```bash
git clone https://github.com/a.ozhogin/git-vendor
cd ./git-vendor
make
```

## Example

```bash
$ # Checkout git@github.com:AOzhogin/example.git@v1.0.0 under git-vendor/github.com/AOzhogin/example
$ git-vendor add example git@github.com:AOzhogin/example.git v1.0.0

git fetch git@github.com:AOzhogin/example.git v1.0.0
warning: no common commits
remote: Enumerating objects: 6, done.
remote: Counting objects: 100% (6/6), done.
remote: Compressing objects: 100% (4/4), done.
remote: Total 6 (delta 0), reused 6 (delta 0), pack-reused 0
Unpacking objects: 100% (6/6), 501 bytes | 8.00 KiB/s, done.
From github.com:AOzhogin/example
 * tag               v1.0.0     -> FETCH_HEAD
Added dir 'git-vendor/github.com/AOzhogin/example'
[master e00ef04] Add "example" from "git@github.com:AOzhogin/example.git@v1.0.0"
 Date: Fri Dec 22 00:17:13 2023 +0300

$ # List current vendored dependencies
$ git vendor list
example@v1.0.0:
        name:   example
        dir:    git-vendor/github.com/AOzhogin/example
        repo:   git@github.com:AOzhogin/example.git
        ref:    v1.0.0

$ # Update existing dependency to a newer version
$ git-vendor update example v2.0.0
From github.com:AOzhogin/example
 * tag               v2.0.0     -> FETCH_HEAD
Merge made by the 'recursive' strategy.
 git-vendor/github.com/AOzhogin/example/main.go | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)
[master e9027d9] Update "example" from "git@github.com:AOzhogin/example.git@v2.0.0"
 Date: Fri Dec 22 00:20:48 2023 +0300

$ # List current vendored dependencies
$ git vendor list
example@v2.0.0:
        name:   example
        dir:    git-vendor/github.com/AOzhogin/example
        repo:   git@github.com:AOzhogin/example.git
        ref:    v2.0.0

```