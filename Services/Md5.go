package Services

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Md5(s string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, s)
	md5Password := hex.EncodeToString(MD5.Sum(nil))
	return md5Password
}
