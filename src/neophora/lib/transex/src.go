package transex

import (
	"encoding/json"
	"neophora/lib/trans"
	"neophora/var/stderr"

	"github.com/neophora/neo2go/pkg/core/transaction"
	"github.com/neophora/neo2go/pkg/io"
)

// T ...
type T struct {
	trans.T
}

// BytesToJSONViaTX ...
func (me *T) BytesToJSONViaTX() error {
	switch bytes := me.V.(type) {
	case []byte:
		var tx transaction.Transaction
		reader := io.NewBinReaderFromBuf(bytes)
		tx.DecodeBinary(reader)
		ret, err := tx.MarshalJSON()
		if err != nil {
			return err
		}
		me.V = json.RawMessage(ret)
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}
