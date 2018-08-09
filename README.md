# partiarelic-client

A gRPC client for [wata-gh/partiarelic](https://github.com/wata-gh/partiarelic).

## Development

We use [golang/dep](https://github.com/golang/dep) to mange libraries.

```
$ dep ensure
```

## Usage

```
$ partiarelic-client --help
Usage of partiarelic-client:
  -retry uint
        Retry count of request (default 3)
  -timeout duration
        Request timeout (default 1s)
  -url string
        URL to request
```

## Release

Use [goreleaser/goreleaser](https://github.com/goreleaser/goreleaser).

```
$ git tag x.y.z
$ GITHUB_TOKEN=xxx goreleaser
```
