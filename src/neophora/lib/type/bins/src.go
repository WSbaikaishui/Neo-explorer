package bins

import (
	"encoding/json"

	"github.com/neophora/neo2go/pkg/core/block"
	"github.com/neophora/neo2go/pkg/io"
)

// T ...
type T []byte

// Valid ...
func (me T) Valid() bool {
	return me != nil
}

// Val ...
func (me T) Val() []byte {
	return []byte(me)
}

// JSONViaBlock ...
func (me T) JSONViaBlock() (json.RawMessage, error) {
	var blk block.Block
	reader := io.NewBinReaderFromBuf(me.Val())
	blk.DecodeBinary(reader)
	ret, err := json.Marshal(blk)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(ret), nil
}
