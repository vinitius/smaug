# Tech Overview
Smaug is a small application built with a "production-ready" mindset.

Some abstractions might strike as overhead at first given the size of the app, but they were meant to help with "real world scenarios".

# Arch
**DDD inspired**, but with a flavor of [Golang Standards](https://github.com/golang-standards/project-layout).

# Packages

```
cmd
|__________main.go:         - program execution
|
internal
|__________domain:          - models
|__________listeners:       - websocket listeners
|__________publishers:      - message broker publishers
|__________aggregates:      - VWAP aggregates
|
pkg
|__________websocket:       - coinbase ws abstraction
|__________config:          - config utils
|__________broker:          - message broker clients
|
test
|__________mocks:           - generated mocks
|
docs
|_____you're here :)


```

# Assumptions
 - **No Integration Tests**: since the coinbase API is public and always available, a whole integration suite to emulate a WS server seemed unnecessary, given the business logic was already validated by the unit tests.
 - **O11y**: it would be crucial to instrument the application in a production-like scenario.
 - **Logs**: a proper log lib such as [Zap](https://github.com/uber-go/zap) could be used to structure logs and help collect them in a better way.
 - **Dependencies**: a proper Dependency Injection framework such as [Wire](https://github.com/google/wire) could be used to reduce boilerplate code upon startup.

