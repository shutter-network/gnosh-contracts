#!/bin/bash

CONTRACTS=(
    "Sequencer"
    "ValidatorRegistry"
)
OUTPUT_DIR="gnoshcontracts"
PACKAGE_NAME="gnoshcontracts"

mkdir -p "$OUTPUT_DIR"

forge build

for contract in "${CONTRACTS[@]}"; do
    abigen --abi <(jq '.["abi"]' "out/${contract}.sol/${contract}.json") --pkg "${PACKAGE_NAME}" --out "$OUTPUT_DIR/${contract}.go"
done
