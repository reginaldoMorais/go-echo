# Golang with Echo

Study of a Golang REST API with Echo framework and MongoDB.

## Summary

- [Golang with Echo](#golang-with-echo)
  - [Summary](#summary)
  - [Dependencies](#dependencies)
  - [Installation](#installation)
    - [Golang](#golang)
    - [Tools to use with Golang](#tools-to-use-with-golang)
      - [Golangci-lint - Golang linter](#golangci-lint---golang-linter)
      - [Air - Live reload](#air---live-reload)
  - [Running](#running)
  - [Thank you](#thank-you)

## Dependencies

- Go 1.19+
- MongoDB
- Docker
- Docker-compose (development mode)

## Installation

### Golang

```bash
## Faça o download e extraia os binários do Go
wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz ## Linux only
sudo tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz

## E adicione os seguintes caminhos no .bashrc ou .zshrc
export PATH="$PATH:/usr/local/go/bin"
export GOPATH="$HOME/go"
export PATH="$PATH:$GOPATH/bin"
```

### Tools to use with Golang

#### Golangci-lint - Golang linter

```bash
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.0

golangci-lint --version
```

#### Air - Live reload

```bash
# binary will be $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# or install it into ./bin/
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

air -v
```

## Running

Run Air to start application with live reload

```bash
make live-reload
```

Or to normal start application

```bash
make start-dev
```

## Thank you

Enjoy the library and Thank you for use it!

[Reginaldo Morais](mailto:reginaldo.cmorais@gmail.com)
