package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//电脑用光标移动,会数出六个字符 6*3 = 18
	const s = "สวัสดี"
	//打印的长长度是计算的byte值
	fmt.Println("Len:", len(s))
	//打印[]byte值 共18个
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	//e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	//分别打印UTF-8 code point值和其偏移量
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
