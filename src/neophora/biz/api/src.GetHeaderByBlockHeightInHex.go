package api

import (
	"neophora/lib/type/uintval"
)

// GetHeaderByBlockHeightInHex ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetHeaderByBlockHeightInHex","params":{"BlockHeight": 2500000}}'
// {"id":1,"result":"00000000cbb03e3afbcc412dee1d7179b2f502b3fda62ebada077519128dc38a41b3d8c4cade1e4e97eed3d9e361e46722fa316c16bdeaea36f62db52ceedf759ec8d9e8f0ed485ba025260001506773dcb17ad183f5011bdaaccee8c4d2555c829fa51e31551ef201fd4501406970d65da82128ac77fa320e89cdb18810e6aece4b5201130d425adfde51400f809c38bf314e4d09304f28fffdd2c0b1b9ff566bb6e4265328f8d3a841ba12304074dea7ffb880d1666b668700bb2f7b8ed6b753285524cb7d858c3baec6d48fefbe41f5abfa09552aa2acc266f67d9be27e4d52aeef6ec1172dd51689a903355e40a66ce2b65854f03fcae1cbead69a76650ee69d6fbb026f83135f46cf40e044385e596a26a34af9c864c86e9f297063823af0b936400374ce2c2038227a9df04e40417f9170bee696eda8ab5c334f1350bedffecfa5edfa5bfb5c8ba962eba114314a72367dd1b622c386d06bf0e239172420876628e1de3df3c9c1153713892ae640e8f8d5f490bdbe8551214790a2214255fd7c6a0b8757276d791a7580708044fc1dfd9094dfcd5475c67f0fef3dfd575cacc423509e9ff6228ccb9c0843db6cadf15521024c7b7fb6c310fccf1ba33b082519d82964ea93868d676662d4a59ad548df0e7d21025bdf3f181f53e9696227843950deb72dcd374ded17c057159513c3d0abe20b642102aaec38470f6aad0042c6e877cfd8087d2676b0f516fddd362801b9bd3936399e2103b209fd4f53a7170ea4444e0cb0a6bb6a53c2bd016926989cf85f9b0fba17a70c2103b8d9d5771d8f513aa0869b9cc8d50986403b78c6da36890638c3d46a5adce04a2102ca0e27697b9c248f6f16e085fd0061e26f44da85b58ee835c110caa5ec3ba5542102df48f60e8f3e01c48ff40b9b7f1310d7a8b2a193188befe1c2e3df740e89509357ae00","error":null}
// ```
func (me *T) GetHeaderByBlockHeightInHex(args struct {
	BlockHeight uintval.T
}, ret *string) error {
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdr",
		Index:  "uint.hgt",
		Keys:   []string{args.BlockHeight.Hex()},
	}, ret)
}
