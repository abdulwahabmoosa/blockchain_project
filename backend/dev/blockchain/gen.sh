#!/usr/bin/env bash

ABI_DIR="../smart_contract/deployments/sepolia"

for abi_file in "$ABI_DIR"/*.abi.json; do
    if [ ! -f "$abi_file" ]; then
        continue
    fi

    contract_name=$(basename "$abi_file" .abi.json)
    snake_case_name=$(echo "$contract_name" | sed 's/\([A-Z]\)/_\1/g' | tr '[:upper:]' '[:lower:]' | sed 's/^_//')

    output_file="blockchain/${snake_case_name}/bindings.go"

    mkdir -p blockchain/${snake_case_name}

    abigen \
        --abi "$abi_file" \
        --pkg "${snake_case_name}" \
        --out "$output_file"

    echo "Generated bindings: $abi_file -> $output_file"

done
