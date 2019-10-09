# Trust

## Docker

https://hub.docker.com/r/lcnem/identity/

## Testnet

```Shell
$ docker run -it -p 26656:26656 -p 26657:26657 -v ~/.identityd:/root/.identityd -v ~/.identitycli:/root/.identitycli lcnem/identity bash init.sh
```

## Mainnet

```Shell
$ docker run -it -p 26656:26656 -p 26657:26657 -v ~/.identityd:/root/.identityd -v ~/.identitycli:/root/.identitycli lcnem/identity cp genesis.json ~/.identityd/config/genesis.json
```

## Start

```Shell
$ docker run -it -p 26656:26656 -p 26657:26657 -v ~/.identityd:/root/.identityd -v ~/.identitycli:/root/.identitycli lcnem/identity identityd start
```