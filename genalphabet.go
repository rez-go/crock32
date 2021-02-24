// Copyright (c) 2015 The btcsuite developers
// Copyright (c) 2018 Saeed Rasooli <saeed.gnu@gmail.com>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

//+build ignore

package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
)

var (
	start = []byte(`// Copyright (c) 2015 The btcsuite developers
// Copyright (c) 2018 Saeed Rasooli <saeed.gnu@gmail.com>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// AUTOGENERATED by genalphabet.go; do not edit.

package crock32

const (
	// alphabet is the modified base32 alphabet used by Bitcoin.
	alphabet = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
	// alphabetLower is the lower-cased version of alphabet. It's
	// cheaper to use lower-cased alphabet than lower-casing the results.
	alphabetLower = "0123456789abcdefghjkmnpqrstvwxyz"

	alphabetIdx0 = '0'
)

var b32 = [256]byte{`)

	end = []byte(`}`)

	alphabet = []byte("0123456789ABCDEFGHJKMNPQRSTVWXYZ")
	tab      = []byte("\t")
	invalid  = []byte("255")
	comma    = []byte(",")
	space    = []byte(" ")
	nl       = []byte("\n")
)

func write(w io.Writer, b []byte) {
	_, err := w.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fi, err := os.Create("alphabet.go")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	write(fi, start)
	write(fi, nl)
	for i := byte(0); i < 32; i++ {
		write(fi, tab)
		for j := byte(0); j < 8; j++ {
			c := i*8 + j
			if c >= 'a' && c <= 'z' {
				c -= 32
			}
			switch c {
			case 'O':
				c = '0'
			case 'I', 'L':
				c = '1'
			}
			idx := bytes.IndexByte(alphabet, c)
			if idx == -1 {
				write(fi, invalid)
			} else {
				write(fi, strconv.AppendInt(nil, int64(idx), 10))
			}
			write(fi, comma)
			if j != 7 {
				write(fi, space)
			}
		}
		write(fi, nl)
	}
	write(fi, end)
	write(fi, nl)
}
