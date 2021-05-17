package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
)

func (me *T) GetBlockCount(args struct {}, ret *json.RawMessage) error {
	_,err:= me.Data.Client.QueryOne(struct {
		Collection string
		Index string
		Sort bson.M
		Filter   bson.M
		Query []string
	}{
		Collection: "Block",
		Index:  "someIndex",
		Sort: bson.M{"_id":-1},
		Filter:   bson.M{},
		Query: []string{"index"},
	},ret)
	if err!=nil {
		return err
	}
	return nil
}
