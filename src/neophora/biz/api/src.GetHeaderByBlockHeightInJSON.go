package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetHeaderByBlockHeightInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetHeaderByBlockHeightInJSON","params":{"BlockHeight": 2500000}}'
// {"id":1,"result":{"hash":"0xeb803cff9f9103087f2bc7a949461ead34331c6d2a64ce680aa610eead71b8fb","version":0,"previousblockhash":"0xc4d8b3418ac38d12197507daba2ea6fdb302f5b279711dee2d41ccfb3a3eb0cb","merkleroot":"0xe8d9c89e75dfee2cb52df636eaeabd166c31fa2267e461e3d9d3ee974e1edeca","time":1531506160,"index":2500000,"nonce":"d17ab1dc73675001","nextconsensus":"ATobfpwv6JBXciEC4bL8GL8PjQkssDsmCR","script":{"invocation":"406970d65da82128ac77fa320e89cdb18810e6aece4b5201130d425adfde51400f809c38bf314e4d09304f28fffdd2c0b1b9ff566bb6e4265328f8d3a841ba12304074dea7ffb880d1666b668700bb2f7b8ed6b753285524cb7d858c3baec6d48fefbe41f5abfa09552aa2acc266f67d9be27e4d52aeef6ec1172dd51689a903355e40a66ce2b65854f03fcae1cbead69a76650ee69d6fbb026f83135f46cf40e044385e596a26a34af9c864c86e9f297063823af0b936400374ce2c2038227a9df04e40417f9170bee696eda8ab5c334f1350bedffecfa5edfa5bfb5c8ba962eba114314a72367dd1b622c386d06bf0e239172420876628e1de3df3c9c1153713892ae640e8f8d5f490bdbe8551214790a2214255fd7c6a0b8757276d791a7580708044fc1dfd9094dfcd5475c67f0fef3dfd575cacc423509e9ff6228ccb9c0843db6cad","verification":"5521024c7b7fb6c310fccf1ba33b082519d82964ea93868d676662d4a59ad548df0e7d21025bdf3f181f53e9696227843950deb72dcd374ded17c057159513c3d0abe20b642102aaec38470f6aad0042c6e877cfd8087d2676b0f516fddd362801b9bd3936399e2103b209fd4f53a7170ea4444e0cb0a6bb6a53c2bd016926989cf85f9b0fba17a70c2103b8d9d5771d8f513aa0869b9cc8d50986403b78c6da36890638c3d46a5adce04a2102ca0e27697b9c248f6f16e085fd0061e26f44da85b58ee835c110caa5ec3ba5542102df48f60e8f3e01c48ff40b9b7f1310d7a8b2a193188befe1c2e3df740e89509357ae"}},"error":null}
// ```
func (me *T) GetHeaderByBlockHeightInJSON(args struct {
	BlockHeight uintval.T
}, ret *json.RawMessage) error {
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdr",
		Index:  "uint.hgt",
		Keys:   []string{args.BlockHeight.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaHeader()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
