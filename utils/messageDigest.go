package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"strings"
)

func MessageDigest(stype string, source []byte) []byte {
	stype = strings.ToLower(stype)
	switch stype {
	case "md5":
		return MD5(source)
	case "sha256":
		return SHA256(source)
	case "sha512":
		return SHA512(source)
	default:
		return nil
	}
}

func SHA512(source []byte) []byte {
	hash := sha512.New()
	//输入数据
	hash.Write(source)
	//计算哈希值
	bytes := hash.Sum(nil)

	return bytes
}
func SHA256(source []byte) []byte {
	hash := sha256.New()
	//输入数据
	hash.Write(source)
	//计算哈希值
	bytes := hash.Sum(nil)

	return bytes
}
func MD5(source []byte) []byte {
	digest := md5.New()
	digest.Write(source)
	data := digest.Sum(nil)
	return data
}
