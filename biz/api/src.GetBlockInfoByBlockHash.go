package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neo3fura/lib/type/h256"
	"neo3fura/var/stderr"
)

func (me *T) GetBlockInfoByBlockHash(args struct {
	BlockHash    h256.T
	Filter map[string]interface{}
}, ret *json.RawMessage) error {
	if args.BlockHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	r1, err := me.Data.Client.QueryOne(
		struct {
			Collection string
			Index      string
			Sort       bson.M
			Filter     bson.M
			Query      []string
		}{
			Collection: "Block",
			Index:      "someIndex",
			Sort:       bson.M{"index":-1},
			Filter: bson.M{"hash": args.BlockHash},
			Query: []string{},
		}, ret)
	if err != nil {
		return err
	}
	print(json.Marshal(r1))
				r3, err := me.Data.Client.QueryAggregate(struct {
					Collection string
					Index      string
					Sort       bson.M
					Filter     bson.M
					Pipeline   []bson.M
					Query      []string
				}{
					Collection: "Transaction",
					Index:      "someIndex",
					Pipeline: []bson.M{bson.M{"$lookup": bson.M{"from": "Asset", "localField": "ParentID", "foreignField": "_id", "as": "asset"}}},
					Sort:       bson.M{},
					Filter:     bson.M{},
					Query:      []string{},
				}, ret)
				if err != nil {
					return err
				}
				if len(r3) != 0{
					r1["totalNetworkFee"] = r3[0]["systemFee"]
					r1["totalSystemFee"] = r3[0]["networkFee"]
				}else {
					r1["totalNetworkFee"] = 0
					r1["totalSystemFee"] = 0
				}
					r, err := json.Marshal(r1)
	if err != nil {
		return err
	}
	*ret = json.RawMessage(r)
	return nil
}
