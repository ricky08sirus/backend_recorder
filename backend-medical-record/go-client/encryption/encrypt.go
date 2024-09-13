package encryption

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
)

func Encrypt(data string) ([]byte, error) {
    block, err := aes.NewCipher([]byte("your-secret-key"))
    if err != nil {
        return nil, err
    }

    ciphertext := make([]byte, aes.BlockSize+len(data))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(data))

    return ciphertext, nil
}

func Decrypt(data []byte) (string, error) {
    block, err := aes.NewCipher([]byte("your-secret-key"))
    if err != nil {
        return "", err
    }

    iv := data[:aes.BlockSize]
    data = data[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(data, data)

    return string(data), nil
}

