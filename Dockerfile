FROM golang:1.15.6

WORKDIR $GOPATH

COPY . .

RUN  go get -u "github.com/btcsuite/btcutil/base58"
RUN  go get -u "go.mongodb.org/mongo-driver/bson"
RUN  go get -u "gopkg.in/yaml.v2"
RUN  go get -u "github.com/joeqian10/neo3-gogogo/helper"
RUN  go get -u "golang.org/x/crypto/ripemd160"

RUN  go build -o main ./src/neo3fura/app/neo3fura/src.go

EXPOSE 1926

CMD ["./main"]
