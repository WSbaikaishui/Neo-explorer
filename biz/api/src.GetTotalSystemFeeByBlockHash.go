package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neo3fura/lib/type/h256"
	"neo3fura/var/stderr"
)

func (me *T) GetTotalSystemFeeByBlockHash(args struct {
	BlockHash h256.T
	Limit     int64
	Skip      int64
	Filter    map[string]interface{}
}, ret *json.RawMessage) error {
	if args.Limit == 0 {
		args.Limit = 200
	}
	if args.BlockHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	r1,err := me.Data.Client.QuerySum(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string

	}{
		Collection: "Transaction",
		Index:      "someIndex",
		Sort:       bson.M{},
		Filter:     bson.M{"blockhash": args.BlockHash.Val()},
		Query:      []string{"sysfee"},
	}, ret)
	if err != nil {
		return err
	}
	r, err := json.Marshal(r1)
	if err != nil {
		return err
	}
	*ret = json.RawMessage(r)
	return nil
}
