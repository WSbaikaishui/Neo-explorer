package api

import (
	"neophora/lib/type/hexs"
	"neophora/var/stderr"
)

// TransformHexToHash256 ...
// as an example:
//
// ```
// TODO
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
