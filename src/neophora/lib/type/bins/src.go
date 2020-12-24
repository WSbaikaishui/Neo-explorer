package bins

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"neophora/var/stderr"

	"github.com/neophora/neo2go/pkg/core/block"
	"github.com/neophora/neo2go/pkg/core/state"
	"github.com/neophora/neo2go/pkg/core/transaction"
	"github.com/neophora/neo2go/pkg/encoding/address"
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

// JSONViaUTXO ...
func (me T) JSONViaUTXO() (json.RawMessage, error) {
	var op transaction.Output
	reader := io.NewBinReaderFromBuf(me.Val())
	op.DecodeBinary(reader)
	ret, err := json.Marshal(map[string]interface{}{
		"asset":   op.AssetID,
		"value":   op.Amount,
		"address": address.Uint160ToString(op.ScriptHash),
	})
	if err != nil {
		return nil, err
	}
	return json.RawMessage(ret), nil
}

// JSONViaAsset ...
func (me T) JSONViaAsset() (json.RawMessage, error) {
	if len(me) < 55 {
		return nil, stderr.ErrInvalidArgs
	}
	bytes := append(me[1:len(me)-46-8], me[len(me)-46:]...)
	var as state.Asset
	obj := make(map[string]interface{})
	reader := io.NewBinReaderFromBuf(bytes)
	as.DecodeBinary(reader)
	obj["id"] = as.ID
	switch as.AssetType {
	case 0x00:
		obj["type"] = "GoverningToken"
	case 0x01:
		obj["type"] = "UtilityToken"
	case 0x08:
		obj["type"] = "Currency"
	case 0x40:
		obj["type"] = "CreditFlag"
	case 0x80:
		obj["type"] = "DutyFlag"
	case 0x80 | 0x10:
		obj["type"] = "Share"
	case 0x80 | 0x18:
		obj["type"] = "Invoice"
	case 0x80 | 0x20:
		obj["type"] = "Token"
	}
	obj["name"] = as.Name
	obj["amount"] = as.Amount
	obj["available"] = as.Available
	obj["precision"] = as.Precision
	obj["owner"] = as.Owner
	obj["admin"] = as.Admin
	obj["issuer"] = as.Issuer
	obj["admin"] = as.Admin
	obj["expiration"] = as.Expiration
	obj["frozen"] = as.IsFrozen
	ret, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(ret), nil
}

// JSONViaCoinState ...
func (me T) JSONViaCoinState() (json.RawMessage, error) {
	if len(me) == 8 {
		return json.RawMessage(`"Untracted State: ` + hex.EncodeToString(me) + `"`), nil
	}
	return json.RawMessage(`"Invalid State: ` + hex.EncodeToString(me) + `"`), nil
}

// JSONViaSpentNEO ...
func (me T) JSONViaSpentNEO() (json.RawMessage, error) {
	if len(me) != 16 {
		return nil, stderr.ErrInvalidArgs
	}
	mint := binary.BigEndian.Uint64(me[0:8])
	burn := binary.BigEndian.Uint64(me[8:16])
	js, err := json.Marshal(map[string]interface{}{
		"mint": mint,
		"burn": burn,
	})
	if err != nil {
		return nil, stderr.ErrInvalidArgs
	}
	return js, nil
}
