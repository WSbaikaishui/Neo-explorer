package api

import (
	"neophora/lib/type/hexs"
	"neophora/var/stderr"
)

// TransformHexToHash256 ...
// as an example:
//
// ```
// $ curl http://127.0.0.1:8888 -d '{"jsonrpc": "2.0","id": 1,"method": "TransformHexToHash256","params":{"Hex":"abcd"}}'
// {"id":1,"result":"179980f6862aedb22205ac97c8af29c77e25d02e189b52926bb1d93796bb3c94","error":null}
// ```
func (me *T) TransformHexToHash256(args struct {
	Hex hexs.T
}, ret *string) error {
	if args.Hex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	*ret = args.Hex.H256()
	return nil
}
