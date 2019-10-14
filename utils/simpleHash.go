package utils

import (
	"encoding/hex"
)

func SimpleHash(algorithm string, source string, salt string, hashIterations int) string {
	iterations := 1
	if iterations < hashIterations {
		iterations = hashIterations
	}
	bytes := hash(algorithm, source, salt, iterations)
	return hex.EncodeToString(bytes)
}

func hash(algorithm string, source string, salt string, hashIterations int) []byte {
	data := salt + source
	encode := MessageDigest(algorithm, []byte(data))

	// 多次加密
	for index := 0; index < (hashIterations - 1); index++ {
		encode = MessageDigest(algorithm, encode)
	}
	return encode
}
