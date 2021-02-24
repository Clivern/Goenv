<p align="center">
    <img alt="Goenv Logo" src="/static/logo.png?v=1.0.0" width="200" />
    <h3 align="center">Goenv</h3>
    <p align="center">Manage Your Applications Go Environment</p>
    <p align="center">
        <a href="https://github.com/Clivern/Goenv/actions/workflows/build.yml">
            <img src="https://github.com/Clivern/Goenv/actions/workflows/build.yml/badge.svg">
        </a>
        <a href="https://github.com/Clivern/Goenv/releases">
            <img src="https://img.shields.io/badge/Version-v1.0.0-red.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/Clivern/Goenv">
            <img src="https://goreportcard.com/badge/github.com/clivern/Goenv?v=1.0.0">
        </a>
        <a href="https://godoc.org/github.com/clivern/goenv">
            <img src="https://godoc.org/github.com/clivern/goenv?status.svg">
        </a>
        <a href="https://github.com/Clivern/Goenv/blob/main/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg">
        </a>
    </p>
</p>
<br/>

## Usage

Download [the latest goenv binary](https://github.com/Clivern/Goenv/releases). Make it executable from everywhere.

```zsh
$ export GOENV_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Clivern/Goenv/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

$ curl -sL https://github.com/Clivern/Goenv/releases/download/v{$GOENV_LATEST_VERSION}/goenv_{$GOENV_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz
```


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
