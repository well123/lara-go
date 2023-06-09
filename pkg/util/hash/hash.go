package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func MD5(b []byte) string {
	h := md5.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func MD5String(s string) string {
	return MD5([]byte(s))
}

func SHA1(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func SHA1String(s string) string {
	return SHA1([]byte(s))
}
