# ghcr-cleanup

`ghcr-cleanup` is a cli utility that allows user to cleanup ghcr container repo while retaining min required versions.

## Usage

```
ghcr cleanup is a cli tool to cleanup old images from ghcr container registry

Usage:
  ghcr-cleanup [flags]

Flags:
      --debug                 enable debug logging
  -h, --help                  help for ghcr-cleanup
      --min-retain int        retain min versions (default 10)
      --package-name string   image name. e.g. translatethread
      --token string          github token, defaults to env variable GITHUB_TOKEN
      --username string       github username. e.g. rajatjindal
      --yes                   actually delete the packages

```bash