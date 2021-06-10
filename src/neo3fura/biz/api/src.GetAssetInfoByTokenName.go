package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neo3fura/lib/type/strval"
	"neo3fura/var/stderr"
)

func (me *T) GetAssetInfoByTokenName(args struct {
	TokenName strval.T
}, ret *json.RawMessage) error {
	if args.TokenName.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	_, err := me.Data.Client.QueryOne(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
	}{
		Collection: "Asset",
		Index:      "someIndex",
		Sort:       bson.M{},
		Filter:     bson.M{"tokenname": args.TokenName.Val()},
		Query:      []string{},
	}, ret)
	if err != nil {
		return err
	}
	return nil
}
