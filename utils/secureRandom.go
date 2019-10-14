package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func SecureRandom(numBytes int) []byte {
	b := make([]byte, numBytes)
	_, err := rand.Read(b) //在byte切片中随机写入元素
	if err != nil {
		// errors.NewError("secure Random error ")
	}
	return b
}

func SecureRandomGenerator() string {
	bytes := SecureRandom(secureRandomGeneratorSize)
	return hex.EncodeToString(bytes)
}
