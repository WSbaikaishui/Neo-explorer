package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neo3fura/lib/type/h256"
	"neo3fura/var/stderr"
)

func (me *T) GetBlockHeaderByBlockHash(args struct {
	BlockHash h256.T
}, ret *json.RawMessage) error {
	if args.BlockHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	_, err := me.Data.Client.QueryOne(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
	}{
		Collection: "Header",
		Index:      "someIndex",
		Sort:       bson.M{},
		Filter:     bson.M{"hash": args.BlockHash},
		Query:      []string{},
	}, ret)
	if err != nil {
		return err
	}
	return nil
}
