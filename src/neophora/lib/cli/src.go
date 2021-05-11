package cli

import "C"
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// T ...
type T struct {
	C *mongo.Client
	Ctx context.Context
}

func (me *T)ListDatabaseNames() error {
	databases, err := me.C.ListDatabaseNames(me.Ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	return nil
}

func (me *T) QueryOne(args struct {
	Collection string
	Index string
	Sort bson.M
	Filter bson.M
	Query []string
}, ret *json.RawMessage ) error {
	var result map[string]interface{}
	convert := make(map[string]interface{})
	collection := me.C.Database("testdb").Collection(args.Collection)
	opts := options.FindOne().SetSort(args.Sort)
	err := collection.FindOne(me.Ctx, args.Filter ,opts).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// TODO
		return errors.New("NOT FOUND")
	} else if err != nil {
		log.Fatal(err)
	}
	if len(args.Query)== 0 {
		convert = result
	} else  {
		for _,v := range args.Query {
			convert[v] = result[v]
		}
	}
	r,err:= json.Marshal(convert)
	if err!= nil{
		return err
	}
	*ret = json.RawMessage(r)
	return nil
}


func (me *T) QueryAll(args struct {
	Collection string
	Index string
	Filter bson.M},
ret *json.RawMessage) error {
	collection :=  me.C.Database("testdb").Collection(args.Collection)
	result, err := collection.Find(me.Ctx,args.Filter)
	if err == mongo.ErrNoDocuments {
		// TODO
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	r, err := json.Marshal(result)
	*ret = json.RawMessage(r)
	return nil
}

func (me *T) Mutation(Collection string, Index string, Keys []string, reply interface{}) {

}
// Call ...
func (me *T) Call(method string, args interface{}, reply interface{}) error {
	//DBs,err := me.C.ListDatabaseNames()
	//ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	collection := me.C.Database("testing").Collection("numbers")
	res, err := collection.InsertOne(me.Ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(res)
	return nil
}
//	var client *rpc.Client
//	var err error
//	select {
//	case client = <-me.Pool:
//	default:
//		client, err = rpc.Dial("tcp", me.Address)
//	}
//	if err != nil {
//		return err
//	}
//	if err = client.Call(method, args, reply); err != nil {
//		client.Close()
//		return err
//	}
//	select {
//	case me.Pool <- client:
//	default:
//		client.Close()
//	}
//	return nil
//}
//
//// Close ...
//func (me *T) Close() {
//	close(me.Pool)
//	for client := range me.Pool {
//		client.Close()
//	}
//}
