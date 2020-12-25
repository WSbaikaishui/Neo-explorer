package api

import (
	"neophora/lib/type/hexs"
	"neophora/var/stderr"
)

// TransformHexToReversedHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) TransformHexToReversedHex(args struct {
	Hex hexs.T
}, ret *string) error {
	if args.Hex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	*ret = args.Hex.RevVal()
	return nil
}
