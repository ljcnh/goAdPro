package gatedLaunch

import (
	"crypto/md5"
	"crypto/sha1"
	"github.com/spaolacci/murmur3"
)

var str = "hello world"

func Md5Hash() [16]byte {
	return md5.Sum([]byte(str))
}
func Sha1Hash() [20]byte {
	return sha1.Sum([]byte(str))
}
func Murmur32() uint32 {
	return murmur3.Sum32([]byte(str))
}
func Murmur64() uint64 {
	return murmur3.Sum64([]byte(str))
}
