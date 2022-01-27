package gatedLaunch

import (
	"testing"
)

func BenchmarkMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Md5Hash()
	}
}

func BenchmarkSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sha1Hash()
	}
}

func BenchmarkMurmurHash32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Murmur32()
	}
}

func BenchmarkMurmurHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Murmur64()
	}
}
