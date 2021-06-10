package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neo3fura/lib/type/h256"
	"neo3fura/var/stderr"
)

func (me *T) GetExecutionByBlockHash(args struct {
	BlockHash h256.T
}, ret *json.RawMessage) error {
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
		Collection: "Execution",
		Index:      "someIndex",
		Sort:       bson.M{},
		Filter:     bson.M{"blockhash": args.BlockHash.Val()},
		Query:      []string{},
	}, ret)
	if err != nil {
		return err
	}
	return nil
}
