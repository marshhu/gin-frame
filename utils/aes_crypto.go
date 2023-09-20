package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/marshhu/gin-frame/log"
	"strings"
)

// EncType 字节数组编码类型
type EncType int

const (
	Hex    EncType = iota + 1 // 16进制编码
	Base64                    // base64编码
)

// defaultEncType 默认字节数组编码类型,统一编码类型
var defaultEncType = Base64

// SetDefaultEncType 设置默认字节数组编码类型
func SetDefaultEncType(encType EncType) {
	defaultEncType = encType
}

// EncBytes 编码字节数组
func EncBytes(bts []byte) string {
	switch defaultEncType {
	case Hex:
		return hex.EncodeToString(bts)
	case Base64:
		return base64.StdEncoding.EncodeToString(bts)
	default:
		return ""
	}
}

// DecStr 编码字符串
func DecStr(str string) ([]byte, error) {
	switch defaultEncType {
	case Hex:
		return hex.DecodeString(str)
	case Base64:
		return base64.StdEncoding.DecodeString(str)
	default:
		return nil, errors.New(fmt.Sprintf("未知编码类型: %v", defaultEncType))
	}
}

// AesCBCEncryptString aes cbc模式加密
func AesCBCEncryptString(plainText, key []byte, ivAes []byte) (cipherText string, err error) {
	if len(strings.TrimSpace(string(plainText))) == 0 {
		return "", nil
	}
	defer func() {
		if e := recover(); e != nil {
			log.Infof("[AesCBCEncryptString]panic plainText:%s,key:%s,iv:%s,err:%v", plainText, string(key),
				string(ivAes), err, e)
			cipherText = string(plainText) //加密失败返回原值
		}
	}()
	err = check(key, ivAes)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	paddingText := pKCS7Padding(plainText, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, ivAes)
	cipherBytes := make([]byte, len(paddingText))
	blockMode.CryptBlocks(cipherBytes, paddingText)
	return EncBytes(cipherBytes), nil
}

// AesCBCDecryptString aes cbc模式解密
func AesCBCDecryptString(cipherText string, key []byte, ivAes []byte) (plainText string, err error) {
	if len(strings.TrimSpace(cipherText)) == 0 {
		return "", nil
	}
	defer func() {
		if e := recover(); e != nil {
			log.Infof("[AesCBCDecryptString]panic cipherText:%s,key:%s,iv:%s，err:%v", cipherText, string(key),
				string(ivAes), e)
			plainText = cipherText //解密失败返回原值
		}
	}()
	bts, err := DecStr(cipherText)
	if err != nil {
		return "", fmt.Errorf("密文编码格式错误，err:%v", err)
	}
	err = check(key, ivAes)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, ivAes)
	paddingText := make([]byte, len(bts))
	blockMode.CryptBlocks(paddingText, bts)

	plainBytes, err := pKCS7UnPadding(paddingText)
	if err != nil {
		return "", err
	}
	return string(plainBytes), nil
}

// pKCS7Padding PKCS7打包
func pKCS7Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - (len(plainText) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}

// pKCS7UnPadding PKCS7解包
func pKCS7UnPadding(plainText []byte) ([]byte, error) {
	length := len(plainText)
	number := int(plainText[length-1])
	if number > length {
		return nil, errors.New("填充长度有误，请检查密钥和向量")
	}
	return plainText[:length-number], nil
}

func check(key []byte, ivAes []byte) (err error) {
	//检查密钥长度
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return errors.New("密钥长度必须为16，24，32字节,请检查密钥长度")
	}
	if len(ivAes) != 16 {
		return errors.New("向量长度必须为16字节，请检查向量长度")
	}
	return nil
}
