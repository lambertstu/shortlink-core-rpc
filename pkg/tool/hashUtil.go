package hashUtil

import (
	"github.com/spaolacci/murmur3"
)

// Base62 字符集
var base62Chars = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var base62Size = uint64(len(base62Chars)) // 运行时计算字符集大小

// 将十进制数转换为 Base62
func convertDecToBase62(num uint64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	sb := make([]byte, 0)
	for num > 0 {
		remainder := num % base62Size
		sb = append([]byte{base62Chars[remainder]}, sb...) // 逆序
		num /= base62Size
	}

	return string(sb)
}

// HashToBase62 计算字符串的 MurmurHash 并转换为 Base62
func HashToBase62(str string) string {
	hash := murmur3.Sum32([]byte(str)) // 计算 32 位 MurmurHash
	num := uint64(hash)
	if hash < 0 { // 处理负数情况
		num = uint64(^hash) // 取反
	}
	return convertDecToBase62(num)
}
