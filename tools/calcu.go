package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

var (
	cliKey     = flag.String("key", "nil", "Key of hamc-md5, in the format of uuid v4.")
	cliversion = flag.Bool("v", false, "Print verion info.")
)

func main() {
	flag.Parse()
	if *cliversion {
		fmt.Println("v1.0.0")
		return
	}
	if isValidUUID4(*cliKey) == false {
		fmt.Println("Invalid key.")
		os.Exit(1)
	}
	t := timeSlice()
	hmacToken, err := hmacMd5(t, *cliKey)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(hmacToken)
}

func timeSlice() uint64 {
	t := time.Now().Unix()
	t = t - (t % 5)
	return uint64(t)
}

func hmacMd5(msg uint64, key string) (string, error) {
	buffer := make([]byte, 8)

	binary.BigEndian.PutUint64(buffer, msg)

	h := hmac.New(md5.New, []byte(key))
	_, err := h.Write(buffer)

	if err != nil {
		return "", errors.Wrap(err, "Hmac-md5 fails")
	}

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func isValidUUID4(v string) bool {
	sb := bytes.Runes([]byte(v))

	if len(sb) != 36 {
		return false
	}

	for i := 0; i < 36; i++ {
		ch := sb[i]
		switch i {
		case 8, 13, 18, 23:
			if ch != '-' {
				return false
			}
		case 14:
			if ch != '4' {
				return false
			}
		case 19:
			switch ch {
			case '8', '9', 'A', 'B', 'a', 'b':
				break
			default:
				return false
			}
		default:
			switch {
			case ch >= '0' && ch <= '9':
				break
			case ch >= 'A' && ch <= 'F':
				break
			case ch >= 'a' && ch <= 'f':
				break
			default:
				return false
			}
		}
	}
	return true
}
