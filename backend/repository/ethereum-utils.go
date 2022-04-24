package repository

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetDefaultTransactionProps(conn *EthereumRepoConnector) (*bind.TransactOpts, error) {
	nonce, err := conn.Client.PendingNonceAt(context.Background(), conn.Address)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	gasPrice, err := conn.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	//TODO: use new type of transactor - current is deprecated
	transactionOpts := bind.NewKeyedTransactor(conn.PrivateKey)
	transactionOpts.Nonce = big.NewInt(int64(nonce))
	transactionOpts.Value = big.NewInt(0)     // in wei
	transactionOpts.GasLimit = uint64(300000) // in units
	transactionOpts.GasPrice = gasPrice

	return transactionOpts, nil
}

func GetDefaultCallOpts() *bind.CallOpts {
	return &bind.CallOpts{Pending: true}
}
