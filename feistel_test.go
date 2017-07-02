package feistel

import (
	"fmt"
	"testing"
)

func TestCipher(*testing.T) {
	for i := 0; i <= 1000; i++ {
		fmt.Printf("%d -> %d\n", i, Cipher(uint64(i)))
	}

}
func TestCipherStr(*testing.T) {
	for i := 0; i <= 1000; i++ {
		fmt.Printf("%d -> %s\n", i, CipherStr(uint64(i)))
	}

}
func ExampleCipher() {
	fmt.Printf("Hello\n")
}
