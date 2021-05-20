package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

func (me *T) GetApplicationLogByBlockHash(args struct {
	BlockHash h256.T
}, ret *json.RawMessage) error {
	if args.BlockHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	r1, err := me.Data.Client.QueryAll(struct {
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
	for _, item2 := range r1 {
		r2, err := me.Data.Client.QueryAll(struct {
			Collection string
			Index      string
			Sort       bson.M
			Filter     bson.M
			Query      []string
			Limit      int64
			Skip       int64
		}{Collection: "[Execution~Notification(Notifications)]", Index: "someIndex", Sort: bson.M{}, Filter: bson.M{"ParentID": item2["_id"]}}, ret)
		if err != nil {
			return err
		}
		notifications := make([]map[string]interface{},0)
        for _ , item3 := range r2 {
			r3, err := me.Data.Client.QueryOne(struct {
				Collection string
				Index      string
				Sort       bson.M
				Filter     bson.M
				Query      []string
			}{Collection: "Notification", Index: "someIndex", Sort: bson.M{}, Filter: bson.M{"_id": item3["ChildID"]}}, ret)
			if err != nil {
				return err
			}
			notifications = append(notifications,r3)
		}
		if len(notifications) > 0 {
			item2["notifications"] = notifications
		} else {
			item2["notifications"] = []map[string]interface{}{}
		}

	}
	r, err := json.Marshal(r1)
	if err != nil {
		return err
	}
	*ret = json.RawMessage(r)
	return nil
}
