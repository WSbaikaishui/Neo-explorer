package api

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
)

func (me *T) GetTotalVotes(args struct {
	Limit  int64
	Skip   int64
	Filter map[string]interface{}
}, ret *json.RawMessage) error {
	r1,_,err := me.Data.Client.QueryAll(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
		Limit      int64
		Skip       int64

	}{
		Collection: "Candidate",
		Index:      "someIndex",
		Sort:       bson.M{},
		Filter:     bson.M{},
		Query:      []string{"votesOfCandidate"},
		Limit:      args.Limit,
		Skip:       args.Skip,

	}, ret)
	if err != nil {
		return err
	}
	r := make(map[string]int64)
	for _,item:= range r1{
		r["total votes"] +=item["votesOfCandidate"].(int64)
	}
	r2, err := json.Marshal(r1)
	if err != nil {
		return err
	}
	*ret = json.RawMessage(r2)
	return nil
}
