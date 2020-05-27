package main

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"net/http"
)

func parseString(s string) ([]byte, error) {
	var b []byte
	b = []byte(s)
	hash := sha3.NewLegacyKeccak256()
	_, err := hash.Write(b)
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func IndexHandler(rw http.ResponseWriter, req *http.Request) {
	str := req.URL.Path
	content := str[1:]
	a, err := parseString(content)
	if err != nil {
		fmt.Println(err)
	}
	s1 := a[0:4]
	_, err = fmt.Fprint(rw, hex.EncodeToString(s1))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	_ = http.ListenAndServe("127.0.0.1:8000", nil)
}
