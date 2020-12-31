package api

import (
	"neophora/lib/type/addr"
	"neophora/var/stderr"
)

// GetAccountByAddressInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetAccountByAddressInHex","params":{"Address":"AZCcft1uYtmZXxzHPr5tY7L6M85zG7Dsrv"}}'
// {"id":1,"result":"00bf28093023643bb0858a119f0e065b450b14564c000002e72d286979ee6cb1b7e65dfddfb2e384100b8d148e7758de42e4168b71792c60d6350c5bee0100009b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc50092b4478e010000","error":null}
// ```
func (me *T) GetAccountByAddressInHex(args struct {
	Address addr.T
}, ret *string) error {
	if args.Address.Valid() == false {
		return stderr.ErrInvalidArgs
	}

	return me.Data.GetLatestUint64ValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.act",
		Index:  "h160.act-uint.hgt",
		Keys:   []string{args.Address.H160()},
	}, ret)
}
