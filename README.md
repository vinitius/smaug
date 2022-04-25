# Smaug
![](https://img.shields.io/badge/coverage-88.5%25-brightgreen) ![](https://img.shields.io/github/go-mod/go-version/vinitius/smaug)

Smaug is a real-time VWAP (Volume-Weighted Average Price) calculation engine.

It uses [Coinbase's WS API](https://docs.cloud.coinbase.com/exchange/docs/websocket-overview) as data feed.

![](https://c.tenor.com/YPOJQhDow3kAAAAC/smaug-treasure.gif)

# Dependencies

- `Go` >= 1.17

- `Make` (optional for a better build experience)

- `Docker` (optional for a better deploy experience)

# Config
You can customize config properties located in `.env` according to the environment:

```
# Application
LOG_LEVEL=debug
SLIDING_WINDOW_SIZE=200

# Coinbase
COINBASE_SERVICE_ADDRESS=ws-feed.exchange.coinbase.com
COINBASE_PRODUCT_IDS=BTC-USD|ETH-USD|ETH-BTC
COINBASE_CHANNELS=matches
```

To create your own `.env` file:
```
make create-env
```

# Run

You can either:

```
make run
```

Or even (if you wish to run it as a standalone container):

```
make docker-run
```

You can also:

*Run in your favorite IDE or straight up: `go run cmd/main.go`*

# Test

Generate mocks:

```
make mock-generate
```

Run tests and generate coverage report:

```
make test
```

# Lint

Run lint:

```
make lint
```

# Tools

Download build tools:

```
make tools
```
Organize and format code:

```
make fmt
```

# Docs
You can find a more detailed overview regarding assumptions and decisions over [here](docs/).

# Profiling
`macOS Catalina i7 16GB`
```
Type: inuse_space
Time: Apr 25, 2022 at 5:34pm (-04)
Nodes accounting for 1744.35kB, 96.72% of 1803.53kB total Dropped 51 nodes (cum <= 9.02kB)
```
More details in [here](docs/smaug_mem_profile.pdf).