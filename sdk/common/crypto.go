package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func (api *App) Token256() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (app *App) SeedDigest(input string) string {
	return hex.EncodeToString(app.SHA256([]byte(os.Getenv("SEED")), []byte(input)))
}

func (app *App) SHA1(b []byte) []byte {
	h := sha1.New()
	h.Write(b)
	return h.Sum(nil)
}

func (app *App) SHA256(b ...[]byte) []byte {
	digest := sha256.Sum256(bytes.Join(b, nil))
	return digest[:]
}

func (app *App) Encrypt(key, payload []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	return gcm.Seal(nonce, nonce, payload, nil), nil
}

func (app *App) Decrypt(key, ciphertext []byte) ([]byte, error) {

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
