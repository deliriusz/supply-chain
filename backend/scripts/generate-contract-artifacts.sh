#!/bin/bash
for CONTRACT in ../frontend/src/contracts/*.sol; do
   echo CONTRACT = $CONTRACT
   CONTRACT_FULL_NAME=$(basename $CONTRACT)
   CONTRACT_NAME="${CONTRACT_FULL_NAME%.*}"
   CONTRACT_NAME_LC=$(echo $CONTRACT_NAME | tr '[:upper:]' '[:lower:]')

   mkdir -p artifacts/$CONTRACT_NAME_LC

	solc @openzeppelin=../frontend/node_modules/@openzeppelin \
		--abi ../frontend/src/contracts/$CONTRACT_FULL_NAME -o build --overwrite
	solc @openzeppelin=../frontend/node_modules/@openzeppelin \
		--bin ../frontend/src/contracts/$CONTRACT_FULL_NAME -o build --overwrite
	abigen --bin=build/$CONTRACT_NAME.bin --abi=build/$CONTRACT_NAME.abi --pkg $CONTRACT_NAME_LC --out=artifacts/$CONTRACT_NAME_LC/$CONTRACT_NAME.go
done