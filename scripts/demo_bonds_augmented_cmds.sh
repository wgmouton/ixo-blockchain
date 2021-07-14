PASSWORD="12345678"
GAS_PRICES="0.025uixo"

SHAUN_DID_FULL='{"did":"did:sov:CYCc2xaJKrp8Yt947Nc6jd","verifyKey":"7HjjYKd4SoBv26MqXp1SzmvDiouQxarBZ2ryscZLK22x","encryptionPublicKey":"FaE44kz98vbKdKh3YWzhe7PTPZ8YsbpDFpdwveGjDgv6","secret":{"seed":"29a58bc799e8ce6a0ee87cc1e42107fc93e9d904f345501fcd92c20172b2603a","signKey":"3oa8GeqqCYpmdXa1TW8Q8CtU1M1PELhkTnNYbhcTamBX","encryptionPrivateKey":"3oa8GeqqCYpmdXa1TW8Q8CtU1M1PELhkTnNYbhcTamBX"}}'
SHAUN_DID="did:sov:CYCc2xaJKrp8Yt947Nc6jd"

# ------------------------------------------------------------------------------

BOND_1_DID_FULL='{"did":"did:ixo:WnRkS9irrycaQNV3txyWUy","verifyKey":"HEWN5zPVewRrQoCjUnHrtwVZXgzhEd9upioub91ZPupc","encryptionPublicKey":"H2wqKmVMJdroe3EMGJdZA8inDqJcKozLinTMBY2HpnPT","secret":{"seed":"1b852020faa6b543e27422e6c6b0d6110c066547811f08581a148ac0ae7424f9","signKey":"2rRkCCkhQEsmGsjyQraoGcZTSLdXowAZNx3GuPft2GFS","encryptionPrivateKey":"2rRkCCkhQEsmGsjyQraoGcZTSLdXowAZNx3GuPft2GFS"}}'
BOND_1_DID="did:ixo:WnRkS9irrycaQNV3txyWUy"
#{
#  "did": "did:ixo:WnRkS9irrycaQNV3txyWUy",
#  "verifyKey": "HEWN5zPVewRrQoCjUnHrtwVZXgzhEd9upioub91ZPupc",
#  "encryptionPublicKey": "H2wqKmVMJdroe3EMGJdZA8inDqJcKozLinTMBY2HpnPT",
#  "secret": {
#    "seed": "1b852020faa6b543e27422e6c6b0d6110c066547811f08581a148ac0ae7424f9",
#    "signKey": "2rRkCCkhQEsmGsjyQraoGcZTSLdXowAZNx3GuPft2GFS",
#    "encryptionPrivateKey": "2rRkCCkhQEsmGsjyQraoGcZTSLdXowAZNx3GuPft2GFS"
#  }
#}

#{
#  "name": "fee1",
#  "type": "local",
#  "address": "ixo1ywa07p3syqf0y5a66kwg6fq3kc5m09vdq8hg99",
#  "pubkey": "ixopub1addwnpepqtvv8rjgatzl7r2fmef53dvamdn3xj4dkrj2qgh5h27prm788l0f5lcfcwv",
#  "mnemonic": "away burden flip leader scout suit course yellow hill donor sell wait guard baby frog assault tray disagree snack horse bean hedgehog leg purity"
#}
ixod keys add fee1 --recover
FEE_1=$(yes $PASSWORD | ixod keys show fee1 -a)

ixod tx bonds create-bond \
  --token=edu \
  --name="Chimple Pilot DIB" \
  --description="Chimple Pilot DIB alpha bond" \
  --function-type=augmented_function \
  --function-parameters="d0:60000000000,p0:1000000,theta:0,kappa:3" \
  --reserve-tokens=xusd \
  --tx-fee-percentage=0 \
  --exit-fee-percentage=0 \
  --fee-address="$FEE_1" \
  --max-supply=1000000000000edu \
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

BOND_2_DID_FULL='{"did":"did:ixo:PU7EJ8mGDczJJunSe5TTng","verifyKey":"DFEo3YWTzuGtS5pQUwDxp1U9T9hoZpU1hMyicfmq9TYS","encryptionPublicKey":"EpsJw3cg5jX5mLja1ztN9pWkEJNPUvPUozfvFTBCWP6f","secret":{"seed":"abef9ae93345dfd8894833d8b9f1757434b4ee23322a4e9d7fc988e3138f3f1e","signKey":"CaAdD9ZShDUVVeTYu7cGq6dT1RJk5hZBA9NGeARjCqZ7","encryptionPrivateKey":"CaAdD9ZShDUVVeTYu7cGq6dT1RJk5hZBA9NGeARjCqZ7"}}'
BOND_2_DID="did:ixo:PU7EJ8mGDczJJunSe5TTng"
#{
#  "did": "did:ixo:PU7EJ8mGDczJJunSe5TTng",
#  "verifyKey": "DFEo3YWTzuGtS5pQUwDxp1U9T9hoZpU1hMyicfmq9TYS",
#  "encryptionPublicKey": "EpsJw3cg5jX5mLja1ztN9pWkEJNPUvPUozfvFTBCWP6f",
#  "secret": {
#    "seed": "abef9ae93345dfd8894833d8b9f1757434b4ee23322a4e9d7fc988e3138f3f1e",
#    "signKey": "CaAdD9ZShDUVVeTYu7cGq6dT1RJk5hZBA9NGeARjCqZ7",
#    "encryptionPrivateKey": "CaAdD9ZShDUVVeTYu7cGq6dT1RJk5hZBA9NGeARjCqZ7"
#  }
#}

#{
#  "name": "fee2",
#  "type": "local",
#  "address": "ixo12xzmp7vy5y3367jva4wp83nqg4mdppjfx8k7ue",
#  "pubkey": "ixopub1addwnpepqg252zaxsmv9gj372u99ezmthqls6xa7ve5uyart5p7tml3pgks6czg8kn3",
#  "mnemonic": "stone poem affair brisk gather circle elevator cabin add discover wish cycle other census wreck scene unique tiger moon title cable balcony human equal"
#}
ixod keys add fee2 --recover
FEE_2=$(yes $PASSWORD | ixod keys show fee2 -a)

ixod tx bonds create-bond \
  --token=abcc \
  --name="Relayer ABC Campaign" \
  --description="Chimple Pilot DIB augmented bonding curve" \
  --function-type=augmented_function \
  --function-parameters="d0:1000000,p0:1000000,theta:0,kappa:1.1358" \
  --reserve-tokens=uixo \
  --tx-fee-percentage=0 \
  --exit-fee-percentage=0.2 \
  --fee-address="$FEE_2" \
  --max-supply=300000abcc \
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
