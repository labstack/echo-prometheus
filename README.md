[![Sourcegraph](https://sourcegraph.com/github.com/labstack/echo-prometheus/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/labstack/echo-prometheus?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/labstack/echo-prometheus)
[![Go Report Card](https://goreportcard.com/badge/github.com/labstack/echo-prometheus?style=flat-square)](https://goreportcard.com/report/github.com/labstack/echo-prometheus)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/labstack/echo-prometheus/main/LICENSE)

# Echo Prometheus metrics and monitoring middleware

[Prometheus](https://prometheus.io/) middleware for [Echo](https://github.com/labstack/echo) framework.

## Versioning

* version `v0.x.y` tracks the latest Echo version (`v5`).
* `main` branch is compatible with the latest Echo version (`v5`).

## Usage

Add Prometheus middleware dependency

```bash
go get github.com/labstack/echo-prometheus
```

Use as an import statement

```go
import echoprometheus "github.com/labstack/echo-prometheus"
```

Add middleware to your Echo instance
```go
e.Use(echoprometheus.NewMiddleware("myapp"))
```
or with config
```go
e.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
    Subsystem: "myapp",
}))
```

Add handler to your Echo instance to expose metrics
```go
e.GET("/metrics", echoprometheus.NewHandler())
```

## Full example

See [example](example/main.go)
