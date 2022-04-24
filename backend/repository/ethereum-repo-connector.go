package repository

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"rafal-kalinowski.pl/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumRepoConnector struct {
	Client     *ethclient.Client
	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

// GetConnector implements RepoConnector
func (c *EthereumRepoConnector) GetConnector() *EthereumRepoConnector {
	return c
}

// InitConnection implements RepoConnector
func (c *EthereumRepoConnector) InitConnection(name, connectionString string) error {
	client, err := ethclient.Dial(config.ETH_PROVIDER_URI)
	if err != nil {
		log.Fatal(err)
		return err
	}

	privateKey, err := crypto.HexToECDSA(config.ETH_PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
		return fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	c.Client = client
	c.PrivateKey = privateKey
	c.PublicKey = publicKeyECDSA
	c.Address = fromAddress

	return nil
}
