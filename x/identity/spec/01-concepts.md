# identity Concepts

アドレスにKey Value Storeを対応させているだけが、IDブロックチェーンになる。
例えばKVSにて`twitter`の値を`lcnem_ja`としたとして、
twitterにて`lcnem_ja`のアカウントから、アドレスだけで文が構成されるツイートを投稿する。
この投稿のURLを`X`とすると、
KVSにて`twitter-proof-url`の値を`X`にする。

## Set

KeyにValueを入れる。

## Import

他のアドレスからデータをインポートする。移行に使う。
他人のを丸コピしたところで、`proof-url`とかが機能してなければどうせ偽物なので、問題ない。

## Delete

消す。