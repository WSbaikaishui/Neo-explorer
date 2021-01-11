package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetBlockByBlockHeightInJSON ...
// as an example:
//
// ```
// $ curl https://example.neophora.io -d '{"jsonrpc": "2.0","id": 1,"method": "GetBlockByBlockHeightInJSON","params":{"BlockHeight":2100000}}'
// {"id":1,"result":{"hash":"0xa163e39321ad75cd3043bd25f3b1618c134c7f5aa6ac503b3a3a783ab86e4fc0","version":0,"previousblockhash":"0x4ddbf266a1cfd711b0e19c4dc6d38a02bafa398d1da35bdb6fb83175381b5b27","merkleroot":"0x0b999a39f2a24a662af215c8288a1f2bfd4344cc6e5f2b6daeda572407029601","time":1522735664,"index":2100000,"nonce":"ff51a98a79b0f548","nextconsensus":"APyEx5f4Zm4oCHwFWiSTaph1fPBxZacYVR","script":{"invocation":"4049424bcdb329469b9a74ebb77d19baf52565b6d11b69f33ec7adce404f6b05b76245b0ad70ac0a3c6dcbd89b357511cb5850954832ab21b7956173db54f3c01440b27ad81c5b29a00bbb816dcb42ef5224b1c8e656daf4476a4d93b3012b438af6aafb199e017e1d15ab7d64624f4a2cc2bfa0bdf6afedc2ea99ecdc32dcd6c24f4043978575374f2bb266a44488481e59d529e2e554acd56ef32f2ce9a3eb41b7f4b134d75e4938e9dc33393d930c64c4a2613a0a0324ed26d16712f7ddba13395e40e33e198830b82315686f9876f3505578d2c35ccf58feec32d02648e969562e476bbe5819bbce92264f28d09a735850b9174e5b70d3c7bb39181a8ae55c9ba9e940932658355e58304eb1ab1f777425d7d1d35027acb3c6eca595cc354a1c8465587135b2b4b99262bda3d12a6f712d8c9a9653ca94dd4852f600439230d619be1d","verification":"552102486fd15702c4490a26703112a5cc1d0923fd697a33406bd5a1c00e0013b09a7021024c7b7fb6c310fccf1ba33b082519d82964ea93868d676662d4a59ad548df0e7d2102aaec38470f6aad0042c6e877cfd8087d2676b0f516fddd362801b9bd3936399e2103b209fd4f53a7170ea4444e0cb0a6bb6a53c2bd016926989cf85f9b0fba17a70c2103b8d9d5771d8f513aa0869b9cc8d50986403b78c6da36890638c3d46a5adce04a2102ca0e27697b9c248f6f16e085fd0061e26f44da85b58ee835c110caa5ec3ba5542102df48f60e8f3e01c48ff40b9b7f1310d7a8b2a193188befe1c2e3df740e89509357ae"},"tx":[{"txid":"0xec5f3381edd9727a87b5c9dd3ee591617041b710ea51e3a0fee2022c7df06b82","size":10,"type":"MinerTransaction","version":0,"attributes":[],"vin":[],"vout":[],"scripts":[],"nonce":2041640264},{"txid":"0x26081981915a1cbae8bd7ccb80a67b38ebf1484a58af86550d7547a9c4b8820f","size":237,"type":"ClaimTransaction","version":0,"attributes":[],"vin":[],"vout":[{"address":"ALUcr5KEQVw8VTW9udjudwgFz87BRpz3vk","asset":"0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7","n":0,"value":"0.0033992"}],"scripts":[{"invocation":"401a9cdefe67d8e11f485c096f4cc79c0892f930a7fde62d78984d83e72658e4ab0dfba62935de9666fc67640242dcf3c9d5e5164838eb4b1e8d7c73f26c2d965c","verification":"2103543a95130f08d4cd1adb99655b2b90928b42c9845c7916136267937ccc521289ac"}],"claims":[{"txid":"0x96c39ed91e6ce72f44bec2685d4bd8bdf454fb428076e08ad5845fe840b4e013","vout":0},{"txid":"0xb49e065fd6323f969f31434e61aaef61ed9fbcc443c680b439d1a684de811734","vout":0}]},{"txid":"0x9ec9ee641eaa1132ff0aeaec8c658d12e0b540a9706ca9f39ad6a51f8e8a1464","size":230,"type":"InvocationTransaction","version":1,"attributes":[{"usage":"Script","data":"bf28093023643bb0858a119f0e065b450b14564c"}],"vin":[],"vout":[],"scripts":[{"invocation":"40fc9e7b966c9be1c32cb4aa2dd45309b17b7306d923d4a4a5cd6d62433a4476a7f00250b4ed69abfdcaf55caa0a7090b29a04f76af61ea782482494ea86c3aa3f","verification":"21035bdb8f130dfbb25037f93a26afdb0ea04973d893711351ed9afca51e3f4a17daac"}],"script":"076027b6e5dff60014aaa45a7cdf6aa49189a3d038f63d97b16a87f1ac14bf28093023643bb0858a119f0e065b450b14564c53c1087472616e7366657267187fc13bec8ff0906c079e7f4cc8276709472913f1660a9fb290d3beece8"},{"txid":"0x56acb6cc7314baffc03a5de8ee9ca127e15267468f1d93f4d4491df33a926513","size":237,"type":"ClaimTransaction","version":0,"attributes":[],"vin":[],"vout":[{"address":"AehHHKP7nWPqueMoNB62tmxErBD3pVFiwk","asset":"0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7","n":0,"value":"5.59725"}],"scripts":[{"invocation":"4014836376e25f9b6f741794d0308ce4f13a518c456da85cf9f14a5b834564271d19b9c8e1a5ba2c8ed304b8432bcc530566e35737263b083e6c609a79a5e07a89","verification":"2103fe4ec8baff71288ee5fd4c283755afe9141a044549c27c24f608991409383451ac"}],"claims":[{"txid":"0x9b2372460a6874eb827ab1dba630a484731f692f017dcdf3d2896a51299e52d3","vout":0},{"txid":"0xb420c9f686a1663d915f8f5fd88c2d54ecfd44c991f6be15273e2561be3fabce","vout":0}]},{"txid":"0xc02035d225cc5098f9c180ed0c9aadfc13fc41a85ab46bfa811636cc45c707c8","size":238,"type":"InvocationTransaction","version":1,"attributes":[{"usage":"Script","data":"b95c6c28f165637d5346108098a051292faaeb09"},{"usage":"Remark1","data":"4f5443474f"}],"vin":[],"vout":[],"scripts":[{"invocation":"40534ae4c46c8b75b3fc0551b0b2cf2fd8112f87c84ec89bc7852bbeab238451e69d14e1cbe83ae5ebb83f56608f410e7b0397334d639c368365da6c0b4c432643","verification":"2102fe007f17e0c492044006434f82b9481a3a80c2a71321583f151fdb4437e7a50aac"}],"script":"080078e4ac14000000144b31e72ad2c57077fa63c5f1c6600d0f04850d5614b95c6c28f165637d5346108098a051292faaeb0953c1087472616e7366657267972e166ea1f8d3c3b14fd8766e7a0dad4084f9e8f166b834f4891afc16f5"},{"txid":"0xe03f7738a630c3e21c512deac3edc1354de7404e28e16c2c6f0a29eb65701887","size":203,"type":"ClaimTransaction","version":0,"attributes":[],"vin":[],"vout":[{"address":"AZHc8LanUdCNotE1eWYg5f2XKNkSAbGcfX","asset":"0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7","n":0,"value":"0.79751448"}],"scripts":[{"invocation":"40dce7317990fdbf0f2ab77435627156dab58a9ea3bfb43f2a08482eab62540575383d84b2d634adc8a2387140f9f0e10b3bc29dd5b8514e7d4a258b48660c1b33","verification":"2102d21b4aae9f05c9bcef25e676c6ffd6acebda80cc73613a3b4d2b2999587e5d03ac"}],"claims":[{"txid":"0x3e8144abe2edda263593f24defd4557d403efa1b4fa839409ac217d6f8a87d3a","vout":0}]},{"txid":"0x3ed654531ba3511c339001778b791005085103d7184f714e53cbb74ed64c9e1a","size":271,"type":"ContractTransaction","version":0,"attributes":[{"usage":"Remark15","data":"6e656f2d6f6e65"}],"vin":[{"txid":"0xccf723d03255b9edad92ee444b4c5d85e0b2de56de76001c04762f9c23ff5576","vout":1}],"vout":[{"address":"AbY6cbQa6yQ8gosHYJAXTE6qupAG93i8dZ","asset":"0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b","n":0,"value":"50"},{"address":"AGYP1uqV9b78sBZwJKbJbgs3mvFbLjzYfB","asset":"0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b","n":1,"value":"20350"}],"scripts":[{"invocation":"40036c78461ba6b4adc2b0b7bbe294f62da92544f685be415965a0009dbcb189141d0fab4b5b413f91f7d1ae2a86da0f157b7ffa291648236ecbe3b392de559969","verification":"2102af2308d7a4ced1fb23b7e95750eb257324a1d848a2cef81c9141c16bf2949b77ac"}]},{"txid":"0xf0f166002fecb9fcaa1f44387ebb3af3c4cca9e1568bf54859edb7f994bc1ea0","size":230,"type":"InvocationTransaction","version":1,"attributes":[{"usage":"Script","data":"bf28093023643bb0858a119f0e065b450b14564c"}],"vin":[],"vout":[],"scripts":[{"invocation":"40a91f9d41564e333703ef9bff9d2e1a116a7abb82cac935f2a52229a08039189e7e19d2abd2434f435b6a53f80d9a70f81cd0c189692b7847debb2edad2c3a62f","verification":"21035bdb8f130dfbb25037f93a26afdb0ea04973d893711351ed9afca51e3f4a17daac"}],"script":"0700b396364e2b01146b9228a7f60cdcd6aba983871726c81f68dc703814bf28093023643bb0858a119f0e065b450b14564c53c1087472616e7366657267187fc13bec8ff0906c079e7f4cc8276709472913f166cfbd71c3c844c93a"},{"txid":"0xb6de6318f21fc2849579bdc113f588f02327baf1592baee9ece81b24be0f1105","size":228,"type":"InvocationTransaction","version":1,"attributes":[{"usage":"Script","data":"2c0e2d921fee4169d714f034e7cc52b56a3e31b4"}],"vin":[],"vout":[],"scripts":[{"invocation":"4007d5f1d0c29bd91f789538fe34138e9525e81b237ee3e93ebe9dc06918628497f44ebb1b09f1f728bfd9d1746dccc1d320986003689b458b3a752fe256f03b86","verification":"2102aa5b52c31458f25b78014504ac11cee9fe6eafee8e228983fc5dff4b0b5ff24dac"}],"script":"05401256656014568d26d1cbe6abc67402aedc9d6e9b5ea68e157d142c0e2d921fee4169d714f034e7cc52b56a3e31b453c1087472616e7366657267f91d6b7085db7c5aaf09f19eeec1ca3c0db2c6ecf166599d93f5ec02c55b"}]},"error":null}
// ```
func (me *T) GetBlockByBlockHeightInJSON(args struct {
	BlockHeight uintval.T
}, ret *json.RawMessage) error {
	if args.BlockHeight.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetArgsInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.blk",
		Index:  "uint.hgt",
		Keys:   []string{args.BlockHeight.Hex()},
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
