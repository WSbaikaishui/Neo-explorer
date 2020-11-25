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
- `adhocstateroot://height/{UINT64}`
- `adhocstateroot://hash/{HASH}`
- `adhoclog://tx/{HASH}`
- `adhocblockinfo://height/{UINT64}`
- `adhocblockinfo://hash/{HASH}`
- `adhocheaderinfo://height/{UINT64}`
- `adhocheaderinfo://hash/{HASH}`
- `adhocsysfee://height/{UINT64}`
- `adhoctxinfo://hash/{HASH}`

- `nep5://asset-account-height/{HASH}/{HASH}/{UINT64}`

## todo

- [ ] getaccountstate ; top
- [ ] getapplicationlog
- [ ] getassetstate ; const
- [ ] getbestblockhash ; top
- [ ] getblockcount ; top
- [ ] getclaimable ; top
- [ ] getcontractstate ; top
- [ ] getnep5balances ; top
- [ ] getrawtransaction
- [ ] getstateheight ; top
- [ ] getstateroot
- [ ] getstorage ; top
- [ ] gettransactionheight
- [ ] gettxout
- [ ] getproof
- [ ] getunclaimed ; top
- [ ] getunspents ; top
- [ ] getutxotransfers ; top
- [ ] getvalidators ; top
- [ ] invokefunction
- [ ] invokescript
- [ ] sendrawtransaction
- [ ] submitblock
- [ ] validateaddress
- [ ] verifyproof
