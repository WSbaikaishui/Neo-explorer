package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/hexs"
	"neophora/var/stderr"
)

// GetStorageByContractHashLEHexKeyInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetStorageByContractHashLEHexKeyInHex","params":{"ContractHashLE":"4d9eab13620fe3569ba3b0e56e2877739e4145e3","HexKey":"736b"}}'
// {"id":1,"result":"0007f5ff4e435db43100","error":null}
// ```
func (me *T) GetStorageByContractHashLEHexKeyInHex(args struct {
	ContractHashLE h160.T
	HexKey         hexs.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
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
		Keys:   []string{args.ContractHashLE.RevVal(), args.HexKey.H256()},
	}, ret)
}
