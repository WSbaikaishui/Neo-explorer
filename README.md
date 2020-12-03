# neophora

NEOPHORA

# requirements

```sh
GOPATH=$(pwd) go get github.com/tecbot/gorocksdb
```

# services

## cli

neocli service

```sh
GOPATH=$(pwd) DBS_ADDRESS="<address_of_dbs>" CLI_PORT="<port>" go run neophora/app/cli/main
```

## rpc

jsonrpc service

TODO ...

## dbs

database service

```sh
GOPATH=$(pwd) DBS_DBPATH="<path_to_database>" DBS_PORT="<port>"  go run neophora/app/dbs/main
```

# database

key is in url form.

value is in binary form.

## key

- `block://height/{UINT64}`
- `block://hash/{HASH}`
- `header://height/{UINT64}`
- `header://hash/{HASH}`
- `hash://height/{UINT64}`
- `tx://hash/{HASH}`
- `height://tx/{HASH}`
- `storage://dbkey-height/{HEX}/{UINT64}`

- `adhocstateroot://height/{UINT64}`
- `adhocstateroot://hash/{HASH}`
- `adhoclog://tx/{HASH}`
- `adhocblockinfo://height/{UINT64}`
- `adhocblockinfo://hash/{HASH}`
- `adhocheaderinfo://height/{UINT64}`
- `adhocheaderinfo://hash/{HASH}`
- `adhocsysfee://height/{UINT64}`
- `adhoctxinfo://hash/{HASH}`
- `adhocaccountstate://account-height/{HASH}/{UINT64}`
- `adhocassetstate://hash-height/{HASH}/{UINT64}`
- `adhoccontractstate://hash-height/{HASH}/{UINT64}`
- `adhocclaimable://account-height/{HASH}/{UINT64}`
- `adhocunspents://account-height/{HASH}/{UINT64}`

- `adhocnep5balance://account-height/{HASH}/{UINT64}`
- `adhocstateheight://height/{UINT64}`
- `adhoctxout://tx-n/{HASH}/{UINT64}`
- `validators://height/{UINT64}`

## todo

- [x] gettransactionheight
- [ ] gettxout
- [ ] getproof
- [x] getunclaimed ; top
- [x] getunspents ; top
- [ ] getvalidators ; top
- [ ] invokefunction
- [ ] invokescript
- [ ] sendrawtransaction
- [ ] submitblock
- [ ] validateaddress
- [ ] verifyproof
