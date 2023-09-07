<p align="center">
    <img alt="Goenv Logo" src="/static/logo.png?v=1.19.0" width="200" />
    <h3 align="center">Goenv</h3>
    <p align="center">Manage Your Applications Go Environment</p>
    <p align="center">
        <a href="https://github.com/clivern/Goenv/actions/workflows/build.yml">
            <img src="https://github.com/clivern/Goenv/actions/workflows/build.yml/badge.svg">
        </a>
        <a href="https://github.com/clivern/Goenv/releases">
            <img src="https://img.shields.io/badge/Version-v1.19.0-red.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/clivern/Goenv">
            <img src="https://goreportcard.com/badge/github.com/clivern/Goenv?v=1.19.0">
        </a>
        <a href="https://godoc.org/github.com/clivern/goenv">
            <img src="https://godoc.org/github.com/clivern/goenv?status.svg">
        </a>
        <a href="https://github.com/clivern/Goenv/blob/main/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg">
        </a>
    </p>
</p>
<br/>
<p align="center">
    <a href="https://www.youtube.com/watch?v=nC344VXoFs4" target="_blank">
        <img src="/static/screenshot-1.png" width="95%" />
    </a>
</p>

Goenv helps you to work with multiple `golang` versions at the same time whether on mac or linux operating system. It supports both global and per-application version configuration.


## Usage

Download [the latest `goenv` binary](https://github.com/clivern/Goenv/releases). Make it executable from everywhere.

```zsh
$ export GOENV_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/clivern/Goenv/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

# For Linux
$ curl -sL https://github.com/clivern/Goenv/releases/download/v{$GOENV_LATEST_VERSION}/goenv_Linux_x86_64.tar.gz | tar xz

# For Mac (Arm M chips)
$ curl -sL https://github.com/clivern/Goenv/releases/download/v{$GOENV_LATEST_VERSION}/goenv_Darwin_arm64.tar.gz | tar xz


# For Mac (Intel)
$ curl -sL https://github.com/clivern/Goenv/releases/download/v{$GOENV_LATEST_VERSION}/goenv_Darwin_x86_64.tar.gz | tar xz
```

Configure the goenv using the following command

```zsh
$ goenv config
```

Add `goenv` shims to `PATH` using the following command. also append it to `~/.profile` file to make it permanent.

```zsh
$ export PATH="$HOME/.goenv/shims:"$PATH

# OR

$ eval "$(goenv init)"
```

Install a new `go` version `1.18` and set as a global

```zsh
$ goenv install 1.18
$ goenv global 1.18
```

To configure a local version different from the global

```zsh
$ goenv local 1.18
```

To Uninstall a version

```zsh
$ goenv uninstall 1.18
```

Show the used version either from current directory or parent directories or the global version.

```zsh
$ goenv version
```

To list all installed versions

```zsh
$ goenv versions
```

for a list of all available commands

```zsh
$ goenv --help

🐺 Manage Your Applications Go Environment.

If you have any suggestions, bug reports, or annoyances please report
them to our issue tracker at <https://github.com/clivern/goenv/issues>

Usage:
  goenv [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Configure the goenv application.
  exec        Show the current go version.
  global      Set or show the global go version.
  help        Help about any command
  info        Print the goenv version
  init        Init the import path for goenv shims.
  install     Install a go version.
  license     Print the license
  local       Set or show the local application-specific go version.
  rehash      Refresh binaries under goenv shim directory.
  satisfy     Satisfy the current directry go version.
  uninstall   Uninstall a specific go version.
  version     Show the current go version.
  versions    List installed go versions.

Flags:
  -h, --help   help for goenv

Use "goenv [command] --help" for more information about a command.
```


## Under The Hood

Goenv is inspired by and works like `rbenv`. At a high level, `goenv` intercepts `Go` commands using `shim` executables injected into your `PATH`, determines which Go version has been specified by your application or globally, and passes your commands to the correct `Go` installation `bin` folder.

**Understanding PATH**

When you run a command like `go` or `gofmt`, your operating system searches through a list of directories to find an executable file with that name. This list of directories lives in an environment variable called `PATH`, with each directory in the list separated by a colon:

```
/usr/local/bin:/usr/bin:/bin
```

Directories in `PATH` are searched from left to right, so a matching executable in a directory at the beginning of the list takes precedence over another one at the end. In this example, the `/usr/local/bin` directory will be searched first, then `/usr/bin`, then `/bin`.

**Understanding Shims**

`goenv` works by inserting a directory of shims at the front of your `PATH`:

```zsh
~/.goenv/shims:/usr/local/bin:/usr/bin:/bin
```

Through a process called rehashing, `goenv` maintains shims in that directory to match every `Go` command across every installed version of `go` like `gofmt` and so on.

`shims` are lightweight executables that simply pass your command to the right binary under the current go version, your operating system will do the following:

1. Search your `PATH` for an executable file named `gofmt`.
2. Find the goenv shim named `gofmt` at the beginning of your `PATH`
3. Run the shim named `gofmt`, which in turn fetch the target go version and use the `gofmt` inside `go/bin` directory.

**Choosing the Go Version**

When you execute a shim, `goenv` determines which Go version to use by reading it from the following sources, in this order:

1. The first `.go-version `file found by searching the current working directory and each of its parent directories until reaching the root of your filesystem. You can modify the `.go-version` file in the current working directory with the `goenv local x.x.x` command.
2. The global `$HOME/.goenv/.go-version` file. You can modify this file using the `goenv global x.x.x` command.


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Goenv is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/goenv/releases) for changelogs for each release version of Goenv. It contains summaries of the most noteworthy changes made in each release. Also see the [Milestones section](https://github.com/clivern/goenv/milestones) for the future roadmap.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/goenv/issues


## Security Issues

If you discover a security vulnerability within Goenv, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2022, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Goenv** is authored and maintained by [@clivern](http://github.com/clivern).
