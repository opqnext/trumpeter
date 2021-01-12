package hash

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

func FileMd5(filename string) (string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return "", errors.New(fmt.Sprintf("[trumpeter] open file error: %v", err))
	}

	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", errors.New(fmt.Sprintf("[trumpeter] io copy error: %v", err))
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func StringM5d(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
