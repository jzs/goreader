package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/taylorskalyo/goreader/epub"
)

func LoadState(book *epub.Rootfile) (int, int) {
	key := fmt.Sprintf("%v-%v-%v",
		book.Publisher,
		book.Creator,
		book.Title,
	)

	skey := fmt.Sprintf("%x.bm", sha256.Sum256([]byte(key)))

	data, err := ioutil.ReadFile(skey)
	if err != nil {
		return 0, 0
	}

	elems := strings.Split(string(data), ",")

	chap, _ := strconv.Atoi(elems[0])
	line, _ := strconv.Atoi(elems[1])

	return chap, line
}

func SaveState(book *epub.Rootfile, chapter, line int) {
	key := fmt.Sprintf("%v-%v-%v",
		book.Publisher,
		book.Creator,
		book.Title,
	)

	skey := fmt.Sprintf("%x.bm", sha256.Sum256([]byte(key)))

	ioutil.WriteFile(skey, []byte(fmt.Sprintf("%v,%v", chapter, line)), 0600)
}
