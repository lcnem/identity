ノードの初期化を行うコマンドの実行例です。

```
➜  ~ docker run -it -p 26656:26656 -p 26657:26657 -v ~/.identityd:/root/.identityd -v ~/.identitycli:/root/.identitycli lcnem/identity sh init.sh
Enter key-name
alice
Enter a passphrase to encrypt your key to disk:
Repeat the passphrase:

- name: alice
  type: local
  address: cosmos1x2xzwf0znr6mr4x68ll9h2fwmxcpc67sz7qfr3
  pubkey: cosmospub1addwnpepq2lxnkkhhwj000eljax73f0f6eaj3ul3qhcngpfevzap7whgpn4nuzy9aqd
  mnemonic: ""
  threshold: 0
  pubkeys: []


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

attack joy flat thing usual burst type lemon stool pulp polar answer forward series rough trade lake quit admit lamp jungle among hold upon
Enter moniker
zeus
Enter chain-id (If left empty, randomly created)
t
{"app_message":{"auth":{"accounts":[],"params":{"max_memo_characters":"256","sig_verify_cost_ed25519":"590","sig_verify_cost_secp256k1":"1000","tx_sig_limit":"7","tx_size_cost_per_byte":"10"}},"bank":{"send_enabled":true},"distribution":{"base_proposer_reward":"0.010000000000000000","bonus_proposer_reward":"0.040000000000000000","community_tax":"0.020000000000000000","delegator_starting_infos":[],"delegator_withdraw_infos":[],"fee_pool":{"community_pool":[]},"outstanding_rewards":[],"previous_proposer":"","validator_accumulated_commissions":[],"validator_current_rewards":[],"validator_historical_rewards":[],"validator_slash_events":[],"withdraw_addr_enabled":true},"genutil":{"gentxs":null},"identity":{},"nft":{"collections":{},"owners":[]},"params":null,"slashing":{"missed_blocks":{},"params":{"downtime_jail_duration":"600000000000","max_evidence_age":"120000000000","min_signed_per_window":"0.500000000000000000","signed_blocks_window":"100","slash_fraction_double_sign":"0.050000000000000000","slash_fraction_downtime":"0.010000000000000000"},"signing_infos":{}},"staking":{"delegations":null,"exported":false,"last_total_power":"0","last_validator_powers":null,"params":{"bond_denom":"stake","max_entries":7,"max_validators":100,"unbonding_time":"1814400000000000"},"redelegations":null,"unbonding_delegations":null,"validators":null},"supply":{"supply":[]}},"chain_id":"t","gentxs_dir":"","moniker":"zeus","node_id":"e7bec7be1908417f58a361d254c2b49488025764"}
Password to sign with 'alice':
Genesis transaction written to "/root/.identityd/config/gentx/gentx-e7bec7be1908417f58a361d254c2b49488025764.json"
{"app_message":{"auth":{"accounts":[{"type":"cosmos-sdk/Account","value":{"account_number":"0","address":"cosmos1x2xzwf0znr6mr4x68ll9h2fwmxcpc67sz7qfr3","coins":[{"amount":"1000000000","denom":"stake"}],"public_key":null,"sequence":"0"}}],"params":{"max_memo_characters":"256","sig_verify_cost_ed25519":"590","sig_verify_cost_secp256k1":"1000","tx_sig_limit":"7","tx_size_cost_per_byte":"10"}},"bank":{"send_enabled":true},"distribution":{"base_proposer_reward":"0.010000000000000000","bonus_proposer_reward":"0.040000000000000000","community_tax":"0.020000000000000000","delegator_starting_infos":[],"delegator_withdraw_infos":[],"fee_pool":{"community_pool":[]},"outstanding_rewards":[],"previous_proposer":"","validator_accumulated_commissions":[],"validator_current_rewards":[],"validator_historical_rewards":[],"validator_slash_events":[],"withdraw_addr_enabled":true},"genutil":{"gentxs":[{"type":"cosmos-sdk/StdTx","value":{"fee":{"amount":[],"gas":"200000"},"memo":"e7bec7be1908417f58a361d254c2b49488025764@172.17.0.2:26656","msg":[{"type":"cosmos-sdk/MsgCreateValidator","value":{"commission":{"max_change_rate":"0.010000000000000000","max_rate":"0.200000000000000000","rate":"0.100000000000000000"},"delegator_address":"cosmos1x2xzwf0znr6mr4x68ll9h2fwmxcpc67sz7qfr3","description":{"details":"","identity":"","moniker":"zeus","security_contact":"","website":""},"min_self_delegation":"1","pubkey":"cosmosvalconspub1zcjduepqecxfqjf7ck08npn37lm79scwcs4nrsw2e3zv2fmwl7uvh78crhrswxajc4","validator_address":"cosmosvaloper1x2xzwf0znr6mr4x68ll9h2fwmxcpc67s825u0z","value":{"amount":"1000000","denom":"stake"}}}],"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"Ar5p2te7pPe/P5dN6KXp1nso8/EF8TQFOWC6HzroDOs+"},"signature":"4HgiqvzECtWxIUKqoiknFgO3PNzfwHoP2N5E1+HANIQ0eGxbBcIfrvttCdyCp+KdbTC4vC1K29kJgCPAjXyE9A=="}]}}]},"identity":{},"nft":{"collections":{},"owners":[]},"params":null,"slashing":{"missed_blocks":{},"params":{"downtime_jail_duration":"600000000000","max_evidence_age":"120000000000","min_signed_per_window":"0.500000000000000000","signed_blocks_window":"100","slash_fraction_double_sign":"0.050000000000000000","slash_fraction_downtime":"0.010000000000000000"},"signing_infos":{}},"staking":{"delegations":null,"exported":false,"last_total_power":"0","last_validator_powers":null,"params":{"bond_denom":"stake","max_entries":7,"max_validators":100,"unbonding_time":"1814400000000000"},"redelegations":null,"unbonding_delegations":null,"validators":null},"supply":{"supply":[]}},"chain_id":"t","gentxs_dir":"/root/.identityd/config/gentx","moniker":"zeus","node_id":"e7bec7be1908417f58a361d254c2b49488025764"}
➜  ~ 
```
