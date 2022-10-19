package totp

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"time"
)

func GenerateTotp(keyString string) (int32, error) {
	key, err := base32.StdEncoding.DecodeString(keyString)
	if err != nil {
		return 0, err
	}
	now := time.Now().Unix() / 30
	mac := hmac.New(sha1.New, key)
	err = binary.Write(mac, binary.BigEndian, now)
	if err != nil {
		return 0, err
	}
	msg := mac.Sum(nil)
	offs := msg[len(msg)-1] & 0xF
	var hash int32
	rdr := bytes.NewReader(msg[offs : offs+4])
	binary.Read(rdr, binary.BigEndian, &hash)
	code := (hash & 0x7FFFFFFF) % 1000000
	return code, nil
}

func CheckTotp(key string, code int32) (bool, error) {
	expected, err := GenerateTotp(key)
	if err != nil {
		return false, err
	}
	return expected == code, nil
}
