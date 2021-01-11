package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetBlockByBlockHashLEInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetBlockByBlockHashLEInJSON","params":{"BlockHashLE":"6b156c0805a229af1efab17b8249b979c9e321217a6fbdee5f500417cc0d5b40"}}'
// {"id":1,"result":{"hash":"0x6b156c0805a229af1efab17b8249b979c9e321217a6fbdee5f500417cc0d5b40","version":0,"previousblockhash":"0x429679f35ceb746961c3e707b79a969fa33c775e94f72c728bdc4b85554b5954","merkleroot":"0xe480d24815f4098cac75b430425cc5b91155662740d7e9c45dec9eb974123083","time":1554824743,"index":3600000,"nonce":"fc4d834b19e27872","nextconsensus":"ANuupE2wgsHYi8VTqSUSoMsyxbJ8P3szu7","script":{"invocation":"400c1c56e15d3ba99dd0e1c3004467e9c661557965607bd790810904fbf562b90f9d91e6997238a66fc43a42e49545d3dcf8264dbfaa537d2cce43a0d158646be2408062b56cbd36fb23e18dc118678bae428a3c904a51d4d0fbdc38c8a536dcfc602d8b129316857d87984ccbde57f3fc7f3dbd3c4ec0b584a74c49e2602fde3fca401d6b7e2e685b65de3d5f8487e1e1ce4d5509ef0c1f0dcdaae7e1a39f13fb49a62d9815166b055a304b3ab3913f5d3eeb6e3173e46d7e57040dbad3c74fabaa6340f86fdde9caab35eb44a082b877aef7b5ee8e85f793478a46e85f6b15c958d2e326515be6b63903bbff68ad2811d1ac5b2e051519d0b0ce9eca20c7f782481e3c409f2cc46bce4483ca621ec004f41ee16ecc55d73d9b2cb7a4c0c80418841edad06f1973402ec91d8f867cbaa3b2a91a9527bd5aefcd7ce8bd02cd592e26a1a74b","verification":"5521024c7b7fb6c310fccf1ba33b082519d82964ea93868d676662d4a59ad548df0e7d21025bdf3f181f53e9696227843950deb72dcd374ded17c057159513c3d0abe20b6421035e819642a8915a2572f972ddbdbe3042ae6437349295edce9bdc3b8884bbf9a32103b209fd4f53a7170ea4444e0cb0a6bb6a53c2bd016926989cf85f9b0fba17a70c2103b8d9d5771d8f513aa0869b9cc8d50986403b78c6da36890638c3d46a5adce04a2102ca0e27697b9c248f6f16e085fd0061e26f44da85b58ee835c110caa5ec3ba5542102df48f60e8f3e01c48ff40b9b7f1310d7a8b2a193188befe1c2e3df740e89509357ae"},"tx":[{"txid":"0xc1ad75d3521772d40c3f3b7891a570df4ebb94d42e01039a62c42a0da77ed5b0","size":10,"type":"MinerTransaction","version":0,"attributes":[],"vin":[],"vout":[],"scripts":[],"nonce":434272370},{"txid":"0xb9afb1ff5911c691bd66966ac3a5aa82feeb5c1698f685cfe3d44ec208e0bf9a","size":240,"type":"InvocationTransaction","version":1,"attributes":[{"usage":"Script","data":"4efa64ae7276a61a3295ad10230a82634cafd6d7"},{"usage":"Remark","data":"00be32fc6a470100000000000000000000000000000000000000000000000000"}],"vin":[],"vout":[],"scripts":[{"invocation":"406a9e846c1ca26abba8e185c2ab173b9aa9638b3170f460da297701d3a6a057951f2fdbeaf48f45168bc96dfc50b900284df0d43502a47bb7aa2184403900bb78","verification":"21025cec967dd8d4246d2116a57bc3e332670942be14d9b7a6c5099a5a9865610c11ac"}],"script":"20999b73ed98557acb63ce57d27af17e7a20821c28a4f0ed4f80539d4c2cbeb80c51c10b63616e63656c4f66666572670ffd83cc1e61bd749ecbb9f998d754860e228894"}]},"error":null}
// ```
func (me *T) GetBlockByBlockHashLEInJSON(args struct {
	BlockHashLE h256.T
}, ret *json.RawMessage) error {
	if args.BlockHashLE.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.blk",
		Index:  "h256.blk",
		Keys:   []string{args.BlockHashLE.RevVal()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaBlock()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
