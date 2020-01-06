# Identity

## Docker

https://hub.docker.com/r/lcnem/identity/

## ノードの準備

testnetへ接続するための初期化を行います。

```Shell
$ docker run -it -p 26656:26656 -p 26657:26657 -v ~/.identityd:/root/.identityd -v ~/.identitycli:/root/.identitycli lcnem/identity sh init.sh
```

mainnetに接続するための準備を行います。testnetへの接続が完了していることが前提です。開発中は実行しないでください。

```Shell
$ docker run -it -p 26656:26656 -p 26657:26657 -v ~/.identityd:/root/.identityd -v ~/.identitycli:/root/.identitycli lcnem/identity cp genesis.json ~/.identityd/config/genesis.json
```

## ノードの稼働開始

```Shell
$ docker run -it --name identity_node -p 26656:26656 -p 26657:26657 -v ~/.identityd:/root/.identityd -v ~/.identitycli:/root/.identitycli lcnem/identity identityd start
```

以下はノードが動作しているコンテナへ接続して実行されます。

CLIコマンドを直接操作したい場合
>$ docker exec -it identity_node sh

CLIコマンドの詳細については `identitycli --help` をご覧ください。

RESTサーバー起動
>$ docker exec -d -it identity_node identitycli rest-server --chain-id t --trust-node=true --laddr tcp://0.0.0.0:1318

## RESTサーバーへのアクセス

通常は `identity-client-ts` クライアントを利用してRESTへアクセスしますが、ダイレクトにアクセスすることも可能です。

例）
>http://localhost:1317/node_info