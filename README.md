# neophora

NEOPHORA

# requirements

```sh
GOPATH=$(pwd) go get golang.org/x/crypto/ripemd160
GOPATH=$(pwd) go get github.com/btcsuite/btcutil/base58
GOPATH=$(pwd) go get github.com/tecbot/gorocksdb
GOPATH=$(pwd) go get github.com/neophora/neo2go/pkg/core/transaction
```

# build


```sh
GOPATH=$(pwd) go install neophora/app/collector
GOPATH=$(pwd) go install neophora/app/neophora
GOPATH=$(pwd) go install neophora/app/server
# GOPATH=$(pwd) go install neophora/app/neorpc
GOPATH=$(pwd) go install neophora/app/neodb
GOPATH=$(pwd) go install neophora/app/txsender
GOPATH=$(pwd) go install neophora/app/importer
GOPATH=$(pwd) go install neophora/app/mrdoc
GOPATH=$(pwd) go install neophora/app/neop
GOPATH=$(pwd) go install neophora/app/dbcleaner
```