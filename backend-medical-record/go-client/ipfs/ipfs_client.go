package ipfs

import (
    "github.com/ipfs/go-ipfs-api"
    "go-client/config"
    "log"
)

func UploadToIPFS(data []byte) (string, error) {
    sh := shell.NewShell(config.GetIPFSNodeURL())
    cid, err := sh.Add(bytes.NewReader(data))
    if err != nil {
        log.Fatalf("Failed to upload to IPFS: %v", err)
        return "", err
    }
    return cid, nil
}

