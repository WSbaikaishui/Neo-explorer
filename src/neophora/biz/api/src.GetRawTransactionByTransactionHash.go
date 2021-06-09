package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

func (me *T) GetRawTransactionByTransactionHash(args struct {
	TransactionHash h256.T
}, ret *json.RawMessage) error {
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	_, err := me.Data.Client.QueryOne(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
	}{
		Collection: "Transaction",
		Index:      "someIndex",
		Sort:       bson.M{},
		Filter:     bson.M{"hash": args.TransactionHash.Val()},
		Query:      []string{},
	}, ret)
	if err != nil {
		return err
	}
	return nil
}