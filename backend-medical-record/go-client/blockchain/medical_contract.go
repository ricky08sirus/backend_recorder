package blockchain

import (
    "log"
    "go-client/config"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/ethclient"
)

func connectToBlockchain() (*ethclient.Client, error) {
    client, err := ethclient.Dial(config.GetEthereumNodeURL())
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }
    return client, nil
}

func AddPatientRecord(
    patientName string, age uint, gender string, disease string,
    contactNumber string, bloodGroup string, ipfsHash string,
) error {
    client, err := connectToBlockchain()
    if err != nil {
        return err
    }

    auth := bind.NewKeyedTransactor(yourPrivateKey)
    tx, err := medicalContract.AddPatientRecord(auth, patientName, age, gender, disease, contactNumber, bloodGroup, ipfsHash)
    if err != nil {
        return err
    }

    log.Printf("Patient record added: %s", tx.Hash().Hex())
    return nil
}

