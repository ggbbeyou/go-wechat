package go_wechat

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func GenRequestString(params map[string]string) (string, error) {
	v := url.Values{}
	for k := range params {
		value := params[k]
		if value != "" {
			v.Set(k, value)
		}
	}
	escapedString := v.Encode()
	return url.QueryUnescape(escapedString)
}

func HashMd5(signStr string) string {
	hasher := md5.New()
	hasher.Write([]byte(signStr))
	sign := strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))
	return sign
}
