#!/bin/sh

echo "Enter key-name"
read key_name
identitycli keys add $key_name

echo "Enter moniker"
read moniker

echo "Enter chain-id (If left empty, randomly created)"
read chain_id
identityd init $moniker --chain-id $chain_id

identityd add-genesis-account $key_name 1000000000stake
valconspub=$(identityd tendermint show-validator)
identityd gentx --amount 1000000stake --name $key_name --pubkey $valconspub
identityd collect-gentxs
