package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"neophora/lib/type/strval"
)

func (me *T) GetNotificationByEvent(args struct {
	Event strval.T
	Limit int64
	Skip  int64
}, ret *json.RawMessage) error {
	if args.Limit == 0 {
		args.Limit = 200
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
		Collection: "Notification",
		Index:      "someIndex",
		Sort:       bson.M{},
		Filter: bson.M{
			"eventname": args.Event.Val(),
		},
		Query: []string{},
		Limit: args.Limit,
		Skip:  args.Skip,
	}, ret)
	if err != nil {
		return err
	}
	return nil
}
