package utils

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/myxzjie/able-admin/utils"
)

func TestPasswordHash(t *testing.T) {

	data := utils.SecureRandom(16) //在byte切片中随机写入元素
	hex.EncodeToString(data)
	salt := "a5aee39c71ae3cb59c61af82a89618d6"
	source := base64.StdEncoding.EncodeToString([]byte("123456"))
	fmt.Println(">>salt:" + salt)
	fmt.Println(">>source:" + source)
	// salt = base64.StdEncoding.EncodeToString([]byte("dmeo" + salt))
	fmt.Println(">>salt:" + salt)
	fmt.Println(">>password:" + ToHex(hash("MD5", "123qwe", ("root"+salt), 2)))
}

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
	encode := utils.MessageDigest(algorithm, []byte(data))

	// 多次加密
	for index := 0; index < (hashIterations - 1); index++ {
		encode = utils.MessageDigest(algorithm, encode)
	}
	return encode
}

// String 将 `[]byte` 转换为 `string`
func String(b []byte) string {
	for idx, c := range b {
		if c == 0 {
			return string(b[:idx])
		}
	}
	return string(b)
}

// StringWithoutZero 将 `[]byte` 转换为 `string`
func StringWithoutZero(b []byte) string {
	s := make([]rune, len(b))
	offset := 0
	for i, c := range b {
		if c == 0 {
			offset++
		} else {
			s[i-offset] = rune(c)
		}
	}
	return string(s[:len(b)-offset-1])
}

func ToHex(b []byte) string {
	str := hex.EncodeToString(b)
	return str
}
