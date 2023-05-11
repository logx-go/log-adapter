# LogX - Log Adapter

Adapter to wrap loggers from Golang log package (https://pkg.go.dev/log)

## Install

```shell
go get -u github.com/logx-go/log-adapter
```


## Usage

```golang
package main

import (
    "log"
    "os"

    "github.com/logx-go/contract/pkg/logx"
    "github.com/logx-go/log-adapter/pkg/logadapter"
)

func main() {
    // we are using the standard adapter from https://pkg.go.dev/log
    logger := logadapter.New(log.New(os.Stdout, "", 0))

    logSomething(logger)
}

func logSomething(logger logx.Logger) {
    logger.Info("This is log message")
}
```

## Development

### Requirement
- Golang >=1.20
- golangci-lint (https://golangci-lint.run/)

### Tests

```shell
go test ./... -race
```

### Lint

```shell
golangci-lint run
```

## License

MIT License (see [LICENSE](LICENSE) file)
