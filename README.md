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
- `adhocblockinfo://height/{NUM}`
- `adhocblockinfo://hash/{HASH}`
- `adhocheaderinfo://height/{NUM}`
- `adhocheaderinfo://hash/{HASH}`
- `adhocsysfee://height/{NUM}`

## todo

- [ ] getaccountstate ; top
- [ ] getapplicationlog
- [ ] getassetstate ; const
- [ ] getbestblockhash ; top
- [ ] getblockcount ; top
- [ ] getblocksysfee
- [ ] getclaimable
- [ ] getconnectioncount
- [ ] getcontractstate
- [ ] getmetricblocktimestamp
- [ ] getnep5balances
- [ ] getnep5transfers
- [ ] getnewaddress
- [ ] getrawmempool
- [ ] getrawtransaction
- [ ] getstateheight
- [ ] getstateroot
- [ ] getstorage
- [ ] gettransactionheight
- [ ] gettxout
- [ ] getpeers
- [ ] getproof
- [ ] getunclaimedgas
- [ ] getunclaimed
- [ ] getunspents
- [ ] getutxotransfers
- [ ] getvalidators
- [ ] getversion
- [ ] getwalletheight
- [ ] importprivkey
- [ ] invokefunction
- [ ] invokescript
- [ ] listplugins
- [ ] listaddress
- [ ] sendfrom
- [ ] sendrawtransaction
- [ ] sendtoaddress
- [ ] sendmany
- [ ] submitblock
- [ ] validateaddress
- [ ] verifyproof
