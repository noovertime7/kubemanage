package utils

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"os"
	"unsafe"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func GormExist(err error) bool {
	return !errors.Is(gorm.ErrRecordNotFound, err)
}

func IsStrEmpty(str string) bool {
	return str == ""
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// ZlibCompress 进行zlib压缩
func ZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// FileExist 判断文件或文件夹是否存在
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func AesDecryptCBC2Hex(encrypted string) string {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ase解密失败: %v", err)
		}
	}()

	key := []byte("NxD3S0yuCc9udD6D")
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	srcData, _ := hex.DecodeString(encrypted)
	decrypted := make([]byte, len(srcData))   // 创建数组
	blockMode.CryptBlocks(decrypted, srcData) // 解密
	decrypted = pkcs5UnPadding(decrypted)     // 去除补全码
	return string(decrypted)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
