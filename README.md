# neophora

NEOPHORA

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

- `block://height/{NUM}`
- `block://hash/{HASH}`
- `header://height/{NUM}`
- `header://hash/{HASH}`
- `hash://height/{NUM}`
- `tx://hash/{HASH}`
- `adhocstateroot://height/{NUM}`
- `adhocstateroot://hash/{HASH}`
- `adhoclog://tx/{HASH}`
- `adhocblockinfo://height/{NUM}`
- `adhocblockinfo://hash/{HASH}`
- `adhocheaderinfo://height/{NUM}`
- `adhocheaderinfo://hash/{HASH}`
- `adhocsysfee://height/{NUM}`
- `adhoctxinfo://hash/{HASH}`

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
