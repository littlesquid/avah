package security

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	rand "math/rand"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

type AesEncryption struct {
	AesGcm cipher.AEAD
	Key    []byte
}

func (aesEncryption *AesEncryption) Init() {
	block, err := aes.NewCipher(aesEncryption.Key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	aesEncryption.AesGcm = aesGCM
}

func (aesEncryption *AesEncryption) GenerateKey() {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, 32)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	aesKey := string(s)

	fmt.Println(aesKey)

	aesEncryption.SaveKey(aesKey)

	aesEncryption.Key = keyToBytes(aesKey)
}

func (aesEncryption *AesEncryption) Encrypt(rawData string) string {
	plaintext := []byte(rawData)
	nouce := make([]byte, aesEncryption.AesGcm.NonceSize())
	if _, err := io.ReadFull(crand.Reader, nouce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesEncryption.AesGcm.Seal(nouce, nouce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func (aesEncryption *AesEncryption) Decrypt(encryptedData string) string {
	enc, _ := hex.DecodeString(encryptedData)
	nouceSize := aesEncryption.AesGcm.NonceSize()
	fmt.Println("nouce size: ", nouceSize)
	nouce, ciphertext := enc[:nouceSize], enc[nouceSize:]

	plaintext, err := aesEncryption.AesGcm.Open(nil, nouce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

func (aesEncryption *AesEncryption) SaveKey(key string) {
	client := redis.NewClient(
		&redis.Options{
			Addr:     viper.GetString("db.redis.host"),
			Password: viper.GetString("db.redis.password"),
			DB:       0,
		})

	_, err := client.Ping().Result()

	err = client.Set("aes_key", key, 0).Err()

	if err != nil {
		fmt.Println(err)
	}
}

func (aesEncryption *AesEncryption) LoadKey() {
	client := redis.NewClient(
		&redis.Options{
			Addr:     viper.GetString("db.redis.host"),
			Password: viper.GetString("db.redis.password"),
			DB:       0,
		})

	_, err := client.Ping().Result()

	aesKey, err := client.Get("aes_key").Result()

	if err != nil {
		fmt.Println(err)
	}

	aesEncryption.Key = keyToBytes(aesKey)
}

func keyToBytes(key string) []byte {
	bytes := []byte(key)

	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	return bytes
}
