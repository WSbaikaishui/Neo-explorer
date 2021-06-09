## INSTALL

```
GOPATH=$(pwd) go get "github.com/btcsuite/btcutil/base58" 
GOPATH=$(pwd) go get "go.mongodb.org/mongo-driver/bson"
GOPATH=$(pwd) go get "gopkg.in/yaml.v2"
GOPATH=$(pwd) go get "github.com/joeqian10/neo3-gogogo/helper"
GOPATH=$(pwd) go get "golang.org/x/crypto/ripemd160"
```

## USAGE

GOPATH=$(pwd) NEOPHORA_PORT="YOUR PORT HERE" go run src/neophora/app/neophora/src.go

