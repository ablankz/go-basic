package standard

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func Crypt() {
	// md5
	h := md5.New()
	io.WriteString(h, "ABCDE")

	// ハッシュ値のバイト配列を返す
	fmt.Println(h.Sum(nil))

	// xで16進数表記
	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)

	// sha-
	s1 := sha1.New()
	io.WriteString(s1, "ABCDE")
	fmt.Printf("%x\n", s1.Sum(nil))

	s256 := sha256.New()
	io.WriteString(s256, "ABCDE")
	fmt.Printf("%x\n", s256.Sum(nil))

	s512 := sha512.New()
	io.WriteString(s512, "ABCDE")
	fmt.Printf("%x\n", s512.Sum(nil))
}
