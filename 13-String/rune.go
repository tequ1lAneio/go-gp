package main

import (
	"fmt"
	"unicode/utf8"
)

func encodeRune() {
	var r rune = 0x4E2D
	fmt.Printf("%c\n", r)
	buf := make([]byte, 3)
	utf8.EncodeRune(buf, r)
	fmt.Printf("utf-8 representation is 0x%X\n", buf)
}

func decodeRune() {
	var buf = []byte{0xE4, 0xB8, 0xAD}
	r, _ := utf8.DecodeRune(buf)
	fmt.Printf("Rhe unicode charactor after decoding [0x#4, 0xB8, 0xAD] is %s\n", string(r))
}

func main() {
	encodeRune()
	decodeRune()
}