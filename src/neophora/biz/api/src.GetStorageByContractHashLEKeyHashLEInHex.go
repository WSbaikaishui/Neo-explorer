package api

import (
	"neophora/lib/type/h160"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetStorageByContractHashLEKeyHashLEInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetStorageByContractHashLEKeyHashLEInHex","params":{"ContractHashLE":"4d9eab13620fe3569ba3b0e56e2877739e4145e3","KeyHashLE":"976923b8377e04bee6bf9b98f6b310bef4b093d078d37735acfee094ee03b121"}}'
// {"id":1,"result":"0007f5ff4e435db43100","error":null}
// ```
func (me *T) GetStorageByContractHashLEKeyHashLEInHex(args struct {
	ContractHashLE h160.T
	KeyHashLE      h256.T
}, ret *string) error {
	if args.ContractHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.KeyHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.str",
		Index:  "h160.ctr-h256.key-uint.hgt",
		Keys:   []string{args.ContractHashLE.RevVal(), args.KeyHashLE.RevVal()},
	}, ret)
}
