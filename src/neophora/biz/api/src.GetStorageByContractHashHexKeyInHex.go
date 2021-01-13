package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/hexs"
	"neophora/var/stderr"
)

// GetStorageByContractHashHexKeyInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetStorageByContractHashHexKeyInHex","params":{"ContractHash":"e345419e7377286ee5b0a39b56e30f6213ab9e4d","HexKey":"736b"}}'
// {"id":1,"result":"0007847a736853b43100","error":null}
// ```
func (me *T) GetStorageByContractHashHexKeyInHex(args struct {
	ContractHash h160.T
	HexKey       hexs.T
}, ret *string) error {
	if args.ContractHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.HexKey.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.str",
		Index:  "h160.ctr-h256.key-uint.hgt",
		Keys:   []string{args.ContractHash.Val(), args.HexKey.H256()},
	}, ret)
}
