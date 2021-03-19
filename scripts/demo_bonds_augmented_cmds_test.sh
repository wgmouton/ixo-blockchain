PASSWORD="12345678"
GAS_PRICES="0.025uixo"

SHAUN_DID_FULL='{"did":"did:sov:CYCc2xaJKrp8Yt947Nc6jd","verifyKey":"7HjjYKd4SoBv26MqXp1SzmvDiouQxarBZ2ryscZLK22x","encryptionPublicKey":"FaE44kz98vbKdKh3YWzhe7PTPZ8YsbpDFpdwveGjDgv6","secret":{"seed":"29a58bc799e8ce6a0ee87cc1e42107fc93e9d904f345501fcd92c20172b2603a","signKey":"3oa8GeqqCYpmdXa1TW8Q8CtU1M1PELhkTnNYbhcTamBX","encryptionPrivateKey":"3oa8GeqqCYpmdXa1TW8Q8CtU1M1PELhkTnNYbhcTamBX"}}'
SHAUN_DID="did:sov:CYCc2xaJKrp8Yt947Nc6jd"

# ------------------------------------------------------------------------------

BOND_1_DID_FULL='{"did":"did:ixo:75SZoisuNpsDezmYURaBb7","verifyKey":"4K6MBrvXwCwn6WBx7MUBjsxH8UnA3Yv1MTSfNsiD8VaL","encryptionPublicKey":"4rJt6UJvZAsNUknteeUKj2nPigTgFKf2vZLiV3wXmSHf","secret":{"seed":"08ac9795490535074eb530a516a72a468d4bd89e56131d1282e8d915fa63da42","signKey":"aruNTNChBW8nLbaP3YyAQSCSKjdk9zRrpyVPkfWP1Vj","encryptionPrivateKey":"aruNTNChBW8nLbaP3YyAQSCSKjdk9zRrpyVPkfWP1Vj"}}'
BOND_1_DID="did:ixo:75SZoisuNpsDezmYURaBb7"
#{
#  "did": "did:ixo:75SZoisuNpsDezmYURaBb7",
#  "verifyKey": "4K6MBrvXwCwn6WBx7MUBjsxH8UnA3Yv1MTSfNsiD8VaL",
#  "encryptionPublicKey": "4rJt6UJvZAsNUknteeUKj2nPigTgFKf2vZLiV3wXmSHf",
#  "secret": {
#    "seed": "08ac9795490535074eb530a516a72a468d4bd89e56131d1282e8d915fa63da42",
#    "signKey": "aruNTNChBW8nLbaP3YyAQSCSKjdk9zRrpyVPkfWP1Vj",
#    "encryptionPrivateKey": "aruNTNChBW8nLbaP3YyAQSCSKjdk9zRrpyVPkfWP1Vj"
#  }
#}

#{
#  "name": "fee1",
#  "type": "local",
#  "address": "ixo19ugeqzwz4rqrz4zp4q4vgvfchgmqma9akm2k9c",
#  "pubkey": "ixopub1addwnpepq2f3l9ndxzhx9dmp6qs6rj6jtsmhlwuggsmmkkg7dv5t37hyr3amv8qaarl",
#  "mnemonic": "holiday soul coin patrol naive blood blossom unique grid among work alpha mountain garlic mesh animal fire scissors secret kind flee satisfy summer roof"
#}
ixocli keys add fee1 --recover
FEE_1=$(yes $PASSWORD | ixocli keys show fee1 -a)

ixocli tx bonds create-bond \
  --token=edutest \
  --name="Chimple Pilot DIB" \
  --description="Chimple Pilot DIB alpha bond" \
  --function-type=augmented_function \
  --function-parameters="d0:60000000000,p0:1000000,theta:0,kappa:3" \
  --reserve-tokens=xusd \
  --tx-fee-percentage=0 \
  --exit-fee-percentage=0 \
  --fee-address="$FEE_1" \
  --max-supply=1000000000000edutest \
  --order-quantity-limits="" \
  --sanity-rate="0" \
  --sanity-margin-percentage="0" \
  --alpha-bond \
  --allow-sells \
  --batch-blocks=1 \
  --outcome-payment="68100000000" \
  --bond-did="$BOND_1_DID" \
  --creator-did="$SHAUN_DID_FULL" \
  --controller-did="$SHAUN_DID" \
  --broadcast-mode block --gas-prices="$GAS_PRICES" -y

# ------------------------------------------------------------------------------

BOND_2_DID_FULL='{"did":"did:ixo:49BSStn5nAwrfZwvGT6HFa","verifyKey":"2iJ4xcXkgMaEUszFB9efBA5NSSzytTfgudoDdrEuQN9e","encryptionPublicKey":"8mmKbhHjdKtJTRg4oqT8PizWdquSGabFkyExj4XAn6LZ","secret":{"seed":"5c52c2afde4f3ed6e3e5a9d472eda67d39a1d206063b194723a9df625a90b0e8","signKey":"7DPhiWu3ov1vLUHFPd4eDPgFWw31PRYoMp6dHhwxMqyh","encryptionPrivateKey":"7DPhiWu3ov1vLUHFPd4eDPgFWw31PRYoMp6dHhwxMqyh"}}'
BOND_2_DID="did:ixo:49BSStn5nAwrfZwvGT6HFa"
#{
#  "did": "did:ixo:49BSStn5nAwrfZwvGT6HFa",
#  "verifyKey": "2iJ4xcXkgMaEUszFB9efBA5NSSzytTfgudoDdrEuQN9e",
#  "encryptionPublicKey": "8mmKbhHjdKtJTRg4oqT8PizWdquSGabFkyExj4XAn6LZ",
#  "secret": {
#    "seed": "5c52c2afde4f3ed6e3e5a9d472eda67d39a1d206063b194723a9df625a90b0e8",
#    "signKey": "7DPhiWu3ov1vLUHFPd4eDPgFWw31PRYoMp6dHhwxMqyh",
#    "encryptionPrivateKey": "7DPhiWu3ov1vLUHFPd4eDPgFWw31PRYoMp6dHhwxMqyh"
#  }
#}

#{
#  "name": "fee2",
#  "type": "local",
#  "address": "ixo120tg9tmzppx2w3xzhpgnzza02cqytax3hgrrzy",
#  "pubkey": "ixopub1addwnpepqgaygsz8henzhc6anc7ln44wz86a0ywu8yhu96m7jvqg830rn3rn5net2zw",
#  "mnemonic": "quote soul lamp cram regret lawsuit reduce black demise excess tuition master swim stomach decorate holiday leaf flag divide pull cloth moral swallow toast"
#}
ixocli keys add fee2 --recover
FEE_2=$(yes $PASSWORD | ixocli keys show fee2 -a)

ixocli tx bonds create-bond \
  --token=abctest \
  --name="Relayer ABC Campaign" \
  --description="Chimple Pilot DIB augmented bonding curve" \
  --function-type=augmented_function \
  --function-parameters="d0:1000000,p0:1000000,theta:0,kappa:1.1358" \
  --reserve-tokens=uixo \
  --tx-fee-percentage=0 \
  --exit-fee-percentage=0.2 \
  --fee-address="$FEE_2" \
  --max-supply=300000abctest \
  --order-quantity-limits="" \
  --sanity-rate="0" \
  --sanity-margin-percentage="0" \
  --allow-sells \
  --batch-blocks=1 \
  --outcome-payment="60000000000" \
  --bond-did="$BOND_2_DID" \
  --creator-did="$SHAUN_DID_FULL" \
  --controller-did="$SHAUN_DID" \
  --broadcast-mode block --gas-prices="$GAS_PRICES" -y

# ------------------------------------------------------------------------------
