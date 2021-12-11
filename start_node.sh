#!/bin/bash

rm -rf $HOME/.simapp/

cd $HOME

simd init --chain-id=testing testing --home=$HOME/.simapp
simd keys add validator --keyring-backend=test --home=$HOME/.simapp
simd add-genesis-account validator 1000000000stake,1000000000valtoken --home=$HOME/.simapp --keyring-backend=test
simd gentx validator 500000000stake --keyring-backend=test --home=$HOME/.simapp --chain-id=testing
simd collect-gentxs --home=$HOME/.simapp

cat $HOME/.simapp/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="10s"' > $HOME/.simapp/config/tmp_genesis.json && mv $HOME/.simapp/config/tmp_genesis.json $HOME/.simapp/config/genesis.json

simd start --home=$HOME/.simapp
