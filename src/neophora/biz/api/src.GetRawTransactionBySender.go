package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neophora/lib/type/addr"
	"neophora/var/stderr"
)

// this function may be not supported any more, we only support address in the formart of script hash
func (me *T) GetRawTransactionBySender(args struct {
	Sender addr.T
	Limit  int64
	Skip   int64
}, ret *json.RawMessage) error {
	if args.Sender.Valid() == false {
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
		Filter:     bson.M{"sender": args.Sender.Val()},
		Query:      []string{},
		Limit:      args.Limit,
		Skip:       args.Skip,
	}, ret)
	if err != nil {
		return err
	}
	return nil
}
