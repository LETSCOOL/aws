package common

import (
	"strings"
	"crypto/sha256"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/md5"
	"fmt"
	"encoding/base64"
	"os"
)

const (
	DEBUG_VERBOSE			= 0				// 0: off, 1~N:
)

var (
	TEST_CREDENTIALS_FILE        string
)

func init() {
	TEST_CREDENTIALS_FILE = os.Getenv("TEST_CREDENTIALS_FILE")
}

func HmacSHA256(key []byte, content string) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	return mac.Sum(nil)
}

func HmacSHA1(key []byte, content string) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(content))
	return mac.Sum(nil)
}

func HashSHA256(content []byte) string {
	h := sha256.New()
	h.Write(content)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func HashMD5(content []byte) string {
	h := md5.New()
	h.Write(content)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Concat(delim string, str ...string) string {
	return strings.Join(str, delim)
}



