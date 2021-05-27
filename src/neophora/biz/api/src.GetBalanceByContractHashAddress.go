package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

func (me *T) GetBalanceByContractHashAddress(args struct {
	ContractHash h160.T
	Address      h160.T
}, ret *json.RawMessage) (map[string]interface{}, error) {
	if args.ContractHash.Valid() == false {
		return nil, stderr.ErrInvalidArgs
	}
	if args.Address.Valid() == false {
		return nil, stderr.ErrInvalidArgs
	}
	r1, err := me.Data.Client.QueryOne(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
	}{
		Collection: "TransferNotification",
		Index:      "someIndex",
		Sort:       bson.M{"_id": -1},
		Filter: bson.M{"contract": args.ContractHash.Val(), "$or": []interface{}{
			bson.M{"from": args.Address.Val()},
			bson.M{"to": args.Address.Val()},
		}},
		Query: []string{},
	}, ret)
	if err != nil {
		return nil, err
	}
	r2 := make(map[string]interface{})
	r2["latesttx"] = r1
	if r1["from"] == args.Address {
		r2["balance"] = r1["frombalance"]
	} else {
		r2["balance"] = r1["tobalance"]
	}
	r, err := json.Marshal(r2)
	if err != nil {
		return nil, err
	}
	*ret = json.RawMessage(r)
	return r2, nil
}
