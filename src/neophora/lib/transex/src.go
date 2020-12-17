package transex

import (
	"encoding/hex"
	"encoding/json"
	"neophora/lib/trans"
	"neophora/var/stderr"

	"github.com/neophora/neo2go/pkg/core/state"
	"github.com/neophora/neo2go/pkg/crypto/keys"
	"github.com/neophora/neo2go/pkg/util"

	"github.com/neophora/neo2go/pkg/core/block"
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
		ret, err := json.Marshal(tx)
		if err != nil {
			return err
		}
		me.V = json.RawMessage(ret)
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}

// BytesToJSONViaBlock ...
func (me *T) BytesToJSONViaBlock() error {
	switch bytes := me.V.(type) {
	case []byte:
		var blk block.Block
		reader := io.NewBinReaderFromBuf(bytes)
		blk.DecodeBinary(reader)
		ret, err := json.Marshal(blk)
		if err != nil {
			return err
		}
		me.V = json.RawMessage(ret)
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}

// BytesToJSONViaHeader ...
func (me *T) BytesToJSONViaHeader() error {
	switch bytes := me.V.(type) {
	case []byte:
		var hd block.Header
		reader := io.NewBinReaderFromBuf(bytes)
		hd.DecodeBinary(reader)
		ret, err := json.Marshal(hd)
		if err != nil {
			return err
		}
		me.V = json.RawMessage(ret)
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}

// BytesToJSONViaAsset ...
func (me *T) BytesToJSONViaAsset() error {
	switch bytes := me.V.(type) {
	case []byte:
		if len(bytes) < 55 {
			return stderr.ErrInvalidArgs
		}
		bytes = append(bytes[1:len(bytes)-46-8], bytes[len(bytes)-46:]...)
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
			return err
		}
		me.V = json.RawMessage(ret)
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}

// BytesToJSONViaAccount ...
func (me *T) BytesToJSONViaAccount() error {
	switch bytes := me.V.(type) {
	case []byte:
		obj := make(map[string]interface{})
		var version byte
		var sh util.Uint160
		var frozen bool
		var votes []*keys.PublicKey
		var balances []map[string]interface{}

		reader := io.NewBinReaderFromBuf(bytes)
		version = reader.ReadB()
		reader.ReadBytes(sh[:])
		frozen = reader.ReadBool()
		reader.ReadArray(&votes)
		n := int(reader.ReadVarUint())
		balances = make([]map[string]interface{}, 0, n)
		for i := 0; i < n; i++ {
			var asset util.Uint256
			var value util.Fixed8
			balance := make(map[string]interface{})
			reader.ReadBytes(asset[:])
			value.DecodeBinary(reader)
			balance["asset"] = asset.StringBE()
			balance["value"] = value.String()
			balances = append(balances, balance)
		}
		obj["version"] = version
		obj["script_hash"] = "0x" + sh.StringLE()
		obj["frozen"] = frozen
		obj["votes"] = votes
		obj["balances"] = balances

		ret, err := json.Marshal(obj)
		if err != nil {
			return err
		}
		me.V = json.RawMessage(ret)
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}

// BytesToJSONViaContract ...
func (me *T) BytesToJSONViaContract() error {
	switch bytes := me.V.(type) {
	case []byte:
		if len(bytes) < 1 {
			return stderr.ErrInvalidArgs
		}
		var cs state.Contract
		obj := make(map[string]interface{})
		reader := io.NewBinReaderFromBuf(bytes[1:])
		cs.DecodeBinary(reader)
		obj["author"] = cs.Author
		obj["properties"] = cs.Properties
		obj["email"] = cs.Email
		obj["parameters"] = cs.ParamList
		obj["hash"] = cs.ScriptHash().StringBE()
		obj["script"] = hex.EncodeToString(cs.Script)
		obj["returntype"] = cs.ReturnType
		obj["name"] = cs.Name
		obj["code_version"] = cs.CodeVersion
		obj["description"] = cs.Description
		ret, err := json.Marshal(obj)
		if err != nil {
			return err
		}
		me.V = json.RawMessage(ret)
		return nil
	default:
		return stderr.ErrInvalidArgs
	}
}
