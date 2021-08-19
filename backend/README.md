# ETB Voting Hub

## Getting started

### Build

```shell
go get ./...
go build -o hub ./cmd/hub
```

## Run

You need to copy `.env.example` to `prod.env` and set variables correctly.

#### Run with Docker

```shell
docker-compose up
```

#### Run locally

```shell
source prod.env
./hub
```
