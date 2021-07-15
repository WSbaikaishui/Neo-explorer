package api

import (
	"encoding/json"
	"fmt"
	"neo3fura/lib/type/h160"
	"neo3fura/var/stderr"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func (me *T) GetAssetHoldersByContractHash(args struct {
	ContractHash h160.T
	Limit        int64
	Skip         int64
	Filter       map[string]interface{}
}, ret *json.RawMessage) error {
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var r1 map[string]interface{}
	r1, err := me.Data.Client.QueryOne(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
	}{
		Collection: "Asset",
		Index:      "someIndex",
		Sort:       bson.M{},
		Filter:     bson.M{"hash": args.ContractHash.Val()},
		Query:      []string{"_id", "totalsupply"},
	}, ret)
	if err != nil {
		return err
	}
	supply, err := strconv.Atoi(r1["totalsupply"].(string))
	r2, count, err := me.Data.Client.QueryAll(
		struct {
			Collection string
			Index      string
			Sort       bson.M
			Filter     bson.M
			Query      []string
			Limit      int64
			Skip       int64
		}{
			Collection: "[Asset~Address(Addresses)]",
			Index:      "someIndex",
			Sort:       bson.M{}, 
			Filter: bson.M{"ParentID": r1["_id"]},
			Query: []string{"ChildID"},
			Limit: args.Limit,
			Skip:  args.Skip},
		ret)
	if err != nil {
		return err
	}
	r3 := make([]map[string]interface{}, 0)
	for _, item := range r2 {
		r, err := me.Data.Client.QueryOne(struct {
			Collection string
			Index      string
			Sort       bson.M
			Filter     bson.M
			Query      []string
		}{Collection: "Address", Index: "someIndex", Sort: bson.M{}, Filter: bson.M{"_id": item["ChildID"]}}, ret)
		if err != nil {
			return err
		}
		var raw map[string]interface{}
		var filter map[string]interface{}
		if args.Filter["balanceinfo"] == nil {
			filter = nil
		} else {
			filter = args.Filter["balanceinfo"].(map[string]interface{})
		}
		err = me.GetBalanceByContractHashAddress(struct {
			ContractHash h160.T
			Address      h160.T
			Filter       map[string]interface{}
			Raw          *map[string]interface{}
		}{
			ContractHash: args.ContractHash,
			Address:      h160.T(fmt.Sprint(r["address"])),
			Filter:       filter,
			Raw:          &raw,
		}, ret)
		if err != nil {
			return err
		}
		r["balance"] = raw["balance"]
		balance, err := strconv.Atoi(raw["balance"].(string))
		if supply != 0 {
			r["percentage"] = float64(balance) / float64(supply)
		} else {
			r["percentage"] = -1
		}
		r3 = append(r3, r)
	}
	r4, err := me.FilterArrayAndAppendCount(r3, count, args.Filter)
	if err != nil {
		return err
	}
	r, err := json.Marshal(r4)
	if err != nil {
		return err
	}
	*ret = json.RawMessage(r)
	return nil
}
