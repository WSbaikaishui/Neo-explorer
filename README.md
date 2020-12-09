# neophora

NEOPHORA

# requirements

```sh
GOPATH=$(pwd) go get github.com/tecbot/gorocksdb
```

# services

## cli

neocli service

```sh
GOPATH=$(pwd) DBS_ADDRESS="<address_of_dbs>" CLI_PORT="<port>" go run neophora/app/cli/main
```

## rpc

jsonrpc service

TODO ...

## dbs

database service

```sh
GOPATH=$(pwd) DBS_DBPATH="<path_to_database>" DBS_PORT="<port>"  go run neophora/app/dbs/main
```

# database

key is in url form.

value is in binary form.
