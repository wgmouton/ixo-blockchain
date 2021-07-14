#!/usr/bin/env bash

wait() {
  echo "Waiting for chain to start..."
  while :; do
    RET=$(ixod status 2>&1)
    if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
      sleep 1
    else
      echo "A few more seconds..."
      sleep 6
      break
    fi
  done
}

RET=$(ixod status 2>&1)
if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
  wait
fi

PASSWORD="12345678"
GAS_PRICES="0.025uixo"
CHAIN_ID="pandora-3"
FEE=$(yes $PASSWORD | ixod keys show fee -a)

ixod_tx() {
  # Helper function to broadcast a transaction and supply the necessary args

  # Get module ($1) and specific tx ($1), which forms the tx command
  cmd="$1 $2"
  shift
  shift

  # Broadcast the transaction
  ixod tx $cmd \
    --gas-prices="$GAS_PRICES" \
    --chain-id="$CHAIN_ID" \
    --broadcast-mode block \
    -y \
    "$@" | jq .
    # The $@ adds any extra arguments to the end
}

ixod_q() {
  ixod q "$@" --output=json | jq .
}

BOND_DID="did:ixo:U7GK8p8rVhJMKhBVRCJJ8c"
#BOND_DID_FULL='{
#  "did":"did:ixo:U7GK8p8rVhJMKhBVRCJJ8c",
#  "verifyKey":"FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW",
#  "encryptionPublicKey":"domKpTpjrHQtKUnaFLjCuDLe2oHeS4b1sKt7yU9cq7m",
#  "secret":{
#    "seed":"933e454dbcfc1437f3afc10a0cd512cf0339787b6595819849f53707c268b053",
#    "signKey":"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC",
#    "encryptionPrivateKey":"Aun1EpjR1HQu1idBsPQ4u4C4dMwtbYPe1SdSC5bUerFC"
#  }
#}'

XUSD_BOND_DID="did:ixo:PU7EJ8mGDczJJunSe5TTng"
#XUSD_BOND_DID_FULL='{
#  "did":"did:ixo:PU7EJ8mGDczJJunSe5TTng",
#  "verifyKey":"DFEo3YWTzuGtS5pQUwDxp1U9T9hoZpU1hMyicfmq9TYS",
#  "encryptionPublicKey":"EpsJw3cg5jX5mLja1ztN9pWkEJNPUvPUozfvFTBCWP6f",
#  "secret":{
#    "seed":"abef9ae93345dfd8894833d8b9f1757434b4ee23322a4e9d7fc988e3138f3f1e",
#    "signKey":"CaAdD9ZShDUVVeTYu7cGq6dT1RJk5hZBA9NGeARjCqZ7",
#    "encryptionPrivateKey":"CaAdD9ZShDUVVeTYu7cGq6dT1RJk5hZBA9NGeARjCqZ7"
#  }
#}'

MIGUEL_ADDR="ixo107pmtx9wyndup8f9lgj6d7dnfq5kuf3sapg0vx"
FRANCESCO_ADDR="ixo1cpa6w2wnqyxpxm4rryfjwjnx75kn4xt372dp3y"
SHAUN_ADDR="ixo1d5u5ta7np7vefxa7ttpuy5aurg7q5regm0t2un"
MIGUEL_DID="did:ixo:4XJLBfGtWSGKSz4BeRxdun"
MIGUEL_DID_FULL='{
  "did":"did:ixo:4XJLBfGtWSGKSz4BeRxdun",
  "verifyKey":"2vMHhssdhrBCRFiq9vj7TxGYDybW4yYdrYh9JG56RaAt",
  "encryptionPublicKey":"6GBp8qYgjE3ducksUa9Ar26ganhDFcmYfbZE9ezFx5xS",
  "secret":{
    "seed":"38734eeb53b5d69177da1fa9a093f10d218b3e0f81087226be6ce0cdce478180",
    "signKey":"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh",
    "encryptionPrivateKey":"4oMozrMR6BXRN93MDk6UYoqBVBLiPn9RnZhR3wQd6tBh"
  }
}'
FRANCESCO_DID_FULL='{
  "did":"did:ixo:UKzkhVSHc3qEFva5EY2XHt",
  "verifyKey":"Ftsqjc2pEvGLqBtgvVx69VXLe1dj2mFzoi4kqQNGo3Ej",
  "encryptionPublicKey":"8YScf3mY4eeHoxDT9MRxiuGX5Fw7edWFnwHpgWYSn1si",
  "secret":{
    "seed":"94f3c48a9b19b4881e582ba80f5767cd3f3c5d7b7103cb9a50fa018f108d89de",
    "signKey":"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM",
    "encryptionPrivateKey":"B2Svs8GoQnUJHg8W2Ch7J53Goq36AaF6C6W4PD2MCPrM"
  }
}'
FRANCESCO_DID="did:ixo:UKzkhVSHc3qEFva5EY2XHt"
SHAUN_DID_FULL='{
  "did":"did:ixo:U4tSpzzv91HHqWW1YmFkHJ",
  "verifyKey":"FkeDue5it82taeheMprdaPrctfK3DeVV9NnEPYDgwwRG",
  "encryptionPublicKey":"DtdGbZB2nSQvwhs6QoN5Cd8JTxWgfVRAGVKfxj8LA15i",
  "secret":{
    "seed":"6ef0002659d260a0bbad194d1aa28650ccea6c6862f994dfdbd48648e1a05c5e",
    "signKey":"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR",
    "encryptionPrivateKey":"8U474VrG2QiUFKfeNnS84CAsqHdmVRjEx4vQje122ycR"
  }
}'

# Ledger DIDs
echo "Ledgering DID 1/3..."
ixod_tx did add-did-doc "$MIGUEL_DID_FULL"
echo "Ledgering DID 2/3..."
ixod_tx did add-did-doc "$FRANCESCO_DID_FULL"
echo "Ledgering DID 3/3..."
ixod_tx did add-did-doc "$SHAUN_DID_FULL"

# d0 := 60000000000   // initial raise (reserve)
# p0 := 1000000       // initial price (reserve per token)
# theta := 0          // initial allocation (percentage)
# kappa := 3          // degrees of polynomial (i.e. x^2, x^4, x^6)

# R0 = 60000000000    // initial reserve (1-theta)*d0
# S0 = 60000          // initial supply
# V0 = 3600           // invariant

echo "Creating bond..."
ixod_tx bonds create-bond \
  --token=edu \
  --name="Chimple Pilot DIB" \
  --description="Chimple Pilot DIB alpha bond" \
  --function-type=augmented_function \
  --function-parameters="d0:60000000000,p0:1000000,theta:0,kappa:3" \
  --reserve-tokens=xusd \
  --tx-fee-percentage=0 \
  --exit-fee-percentage=0 \
  --fee-address="$FEE" \
  --max-supply=1000000000000edu \
  --order-quantity-limits="" \
  --sanity-rate="0" \
  --sanity-margin-percentage="0" \
  --alpha-bond \
  --allow-sells \
  --batch-blocks=1 \
  --outcome-payment="68100000000" \
  --bond-did="$BOND_DID" \
  --creator-did="$MIGUEL_DID_FULL" \
  --controller-did="$FRANCESCO_DID"
echo "Created bond..."
ixod_q bonds bond "$BOND_DID"

echo "Creating xUSD bond..."
ixod_tx bonds create-bond \
  --token=xusd \
  --name="xUSD Minting Bond" \
  --description="Bond for minting xUSD tokens" \
  --function-type=power_function \
  --function-parameters="m:1,n:0.000000000000000001,c:0" \
  --reserve-tokens=uixo \
  --tx-fee-percentage=0 \
  --exit-fee-percentage=0 \
  --fee-address="$FEE" \
  --max-supply=10000000000000xusd \
  --order-quantity-limits="" \
  --sanity-rate="0" \
  --sanity-margin-percentage="0" \
  --allow-sells \
  --batch-blocks=1 \
  --bond-did="$XUSD_BOND_DID" \
  --creator-did="$MIGUEL_DID_FULL" \
  --controller-did="$FRANCESCO_DID"
echo "Created bond..."
ixod_q bonds bond "$XUSD_BOND_DID"

echo "Minting 1000000USD tokens to Miguel using the xUSD bond..."
ixod_tx bonds buy 1000000000000xusd 1000000000001uixo "$XUSD_BOND_DID" "$MIGUEL_DID_FULL"
echo "Minting 1000000USD tokens to Francesco using the xUSD bond..."
ixod_tx bonds buy 1000000000000xusd 1000000000001uixo "$XUSD_BOND_DID" "$FRANCESCO_DID_FULL"

echo "Miguel buys 60000edu... [[founder allocation]]"
ixod_tx bonds buy 60000edu 60000000000xusd "$BOND_DID" "$MIGUEL_DID_FULL"
echo "Miguel's account..."
ixod_q bank balances "$MIGUEL_ADDR"

echo "Bond state is now open..."  # since 60000 (S0) reached
ixod_q bonds bond "$BOND_DID"

echo "Miguel buys 8682edu... [[Investment tranche 1 -- worth 30000xUSD]]"
ixod_tx bonds buy 8682edu 30000000000xusd "$BOND_DID" "$MIGUEL_DID_FULL"
echo "Miguel's account..."
ixod_q bank balances "$MIGUEL_ADDR"

echo "Changing public alpha 0.5->0.55..."
NEW_ALPHA="0.55"
ixod_tx bonds set-next-alpha "$NEW_ALPHA" "$BOND_DID" "$FRANCESCO_DID_FULL"

echo "Miguel buys 5639edu... [[Investment tranche 2 -- worth 10000xUSD]]"
ixod_tx bonds buy 5639edu 10000000000xusd "$BOND_DID" "$MIGUEL_DID_FULL"
echo "Miguel's account..."
ixod_q bank balances "$MIGUEL_ADDR"

echo "Changing public alpha 0.55->0.65..."
NEW_ALPHA="0.65"
ixod_tx bonds set-next-alpha "$NEW_ALPHA" "$BOND_DID" "$FRANCESCO_DID_FULL"

echo "Miguel buys 5631edu... [[Investment tranche 3 -- worth 10000xUSD]]"
ixod_tx bonds buy 5631edu 10000000000xusd "$BOND_DID" "$MIGUEL_DID_FULL"
echo "Miguel's account..."
ixod_q bank balances "$MIGUEL_ADDR"

echo "Changing public alpha 0.65->0.75..."
NEW_ALPHA="0.75"
ixod_tx bonds set-next-alpha "$NEW_ALPHA" "$BOND_DID" "$FRANCESCO_DID_FULL"

echo "Miguel buys 5619edu... [[Investment tranche 4 -- worth 10000xUSD]]"
ixod_tx bonds buy 5619edu 10000000000xusd "$BOND_DID" "$MIGUEL_DID_FULL"
echo "Miguel's account..."
ixod_q bank balances "$MIGUEL_ADDR"

echo "Francesco makes outcome payment of 68100xUSD..."
ixod_tx bonds make-outcome-payment "$BOND_DID" "68100000000" "$FRANCESCO_DID_FULL"
echo "Francesco's account..."
ixod_q bank balances "$FRANCESCO_ADDR"
echo "Bond outcome payment reserve is now 68100xUSD..."
ixod_q bonds bond "$BOND_DID"

echo "Francesco updates the bond state to SETTLE"
ixod_tx bonds update-bond-state "SETTLE" "$BOND_DID" "$FRANCESCO_DID_FULL"
echo "Bond outcome payment reserve is now empty (moved to main reserve)..."
ixod_q bonds bond "$BOND_DID"

echo "Miguel withdraws share..."
ixod_tx bonds withdraw-share "$BOND_DID" "$MIGUEL_DID_FULL"
echo "Miguel's account..."
ixod_q bank balances "$MIGUEL_ADDR"

echo "Bond reserve is now empty and supply is 0..."
ixod_q bonds bond "$BOND_DID"
