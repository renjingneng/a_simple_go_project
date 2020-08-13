package sensetime

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func makeNonce(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	nonce := ""
	l := len(chars)
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		r := rand.Intn(l)
		nonce = nonce + string(chars[r])
	}
	return nonce
}

func makeStringSignature(nonce, timestamp, ApiKey string) string {
	payload := []string{nonce, timestamp, ApiKey}
	sort.Strings(payload)
	joinStr := strings.Join(payload, "")
	return joinStr
}

func signString(stringSignature, ApiSecret string) string {
	key := []byte(ApiSecret)
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(stringSignature))
	resByte := mac.Sum(nil)
	res := hex.EncodeToString(resByte)
	return res
}
