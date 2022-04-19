# Events

## What is this?

This is the `events` microservice that will handle any and all data, queries, and mutations regarding events.

## Requirements

- Golang 1.18

## Quickstart

Start by cloning this repository, then setting your GOBIN and PATH variables by doing the following:
```bash
export GOBIN=$PWD/bin
export PATH=$GOBIN:$PATH
```
Keep in mind these environmental variables are per microservice so don't reuse a terminal seesion.

TODO: Will create a script to handle further tasks

Execute the following:
```
go install github.com/99designs/gqlgen
```

## Regenerating schema

After installing gqlgen to your local bin you can do following:
```bash
gqlgen generate
```

