package blockchain

import (
    "log"
    "math/big"
    "go-client/config"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/rpc"
)

func connectToBlockchain() (*ethclient.Client, error) {
    client, err := ethclient.Dial(config.GetEthereumNodeURL())
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }
    return client, nil
}

func AddRecord(ipfsHash string) error {
    client, err := connectToBlockchain()
    if err != nil {
        return err
    }

    auth := bind.NewKeyedTransactor(yourPrivateKey)
    tx, err := medicalContract.AddRecord(auth, ipfsHash)
    if err != nil {
        return err
    }

    log.Printf("Record added: %s", tx.Hash().Hex())
    return nil
}

