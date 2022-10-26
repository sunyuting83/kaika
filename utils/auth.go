package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	LevelDB "kaika/Leveldb"

	"github.com/gin-gonic/gin"
)

// VerifyMiddleware Verify middleware
func VerifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) > 67 {
			secret_key, _ := c.Get("secret_key")
			SECRET_KEY := secret_key.(string)
			token = token[7:]
			if CheckToken(SECRET_KEY, token) {
				c.Next()
			} else {
				c.AbortWithStatus(403)
			}
		} else {
			c.AbortWithStatus(403)
		}
	}
}

// CheckToken is a check token function
func CheckToken(s, a string) bool {
	AEStoken, err := DecryptByAes(a, []byte(s))
	if err != nil {
		return false
	}
	token, err := LevelDB.Get(string(AEStoken))
	if err != nil {
		return false
	}
	if string(token) == string(AEStoken) {
		return true
	}
	return false
}

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

// AesEncrypt 加密
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

// AesDecrypt 解密
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

// EncryptByAes Aes加密 后 base64 再加
func EncryptByAes(data, PwdKey []byte) (string, error) {
	res, err := AesEncrypt(data, PwdKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

// DecryptByAes Aes 解密
func DecryptByAes(data string, PwdKey []byte) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return AesDecrypt(dataByte, PwdKey)
}
