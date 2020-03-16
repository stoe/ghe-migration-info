# ghe-migration-info

> Get info about your GitHub Enterprise Server,
> e.g. when planning a migration to another GitHub Enterprise Server or to GitHub Enterprise Cloud

[![build](https://github.com/stoe/ghe-migration-info/workflows/build/badge.svg)](https://github.com/stoe/ghe-migration-info/actions?query=workflow%3Abuild) [![release](https://github.com/stoe/ghe-migration-info/workflows/release/badge.svg)](https://github.com/stoe/ghe-migration-info/actions?query=workflow%3Arelease)

## Install

```sh
$ go get github.com/stoe/ghe-migration-info
```

Or download the the latest release binary for your platform: [github.com/stoe/ghe-migration-info](https://github.com/stoe/ghe-migration-info/releases)

## Usage

```sh
USAGE:
  ghe-migration-info [OPTIONS]

OPTIONS:
  -c, --config string     path to the YAML config file (defaults to $HOME/)
      --help              print this help
  -h, --hostname string   hostname
  -t, --token string      personal access token

EXAMPLE:
  $ ghe-migration-info -h github.example.com -t AA123...
```

The scripts requires a personal access token with a `site_admin` scope.

Create a Personal Access Token (PAT) for GitHub Enterprise Server, `https://HOSTNAME/settings/tokens/new?description=ghe-migration-info&scopes=site_admin`.

## License

MIT © [Stefan Stölzle](https://github.com/stoe)
