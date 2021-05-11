package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

func (me *T) GetBlock(args struct {
	BlockHeight uintval.T
}, ret *json.RawMessage) error {
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.Client.QueryOne(struct {
		Collection string
		Index string
		Sort bson.M
		Filter   bson.M
		Query []string
	}{
		Collection: "Block",
		Index:  "someIndex",
		Sort: bson.M{},
		Filter:   bson.M{"index": args.BlockHeight},
		Query: []string{},
	},ret)
}