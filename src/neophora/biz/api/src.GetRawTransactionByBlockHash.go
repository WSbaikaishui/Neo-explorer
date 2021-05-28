package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

func (me *T) GetRawTransactionByBlockHash(args struct {
	BlockHash h256.T
	Limit     int64
	Skip      int64
	Query     []string
}, ret *json.RawMessage) error {
	if args.Limit == 0 {
		args.Limit = 200
	}
	if args.BlockHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	_, err := me.Data.Client.QueryAll(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
		Limit      int64
		Skip       int64
	}{
		Collection: "Transaction",
		Index:      "someIndex",
		Sort:       bson.M{},
		Filter:     bson.M{"blockhash": args.BlockHash.Val()},
		Query:      []string{},
		Limit:      args.Limit,
		Skip:       args.Skip,
	}, ret)
	if err != nil {
		return err
	}
	return nil
}
