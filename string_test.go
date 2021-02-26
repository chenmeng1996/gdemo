package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestStrInfo(t *testing.T) {
	str := "陈濛abc"

	fmt.Printf("%x\n", str)

	// 打印每个字节的16进制
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x\t", str[i])
	}
	fmt.Println()

	fmt.Printf("%q\n", str)

	// 如果该字节表示一个字符，则打印该字符，否则输出Unicode编码
	for i := 0; i < len(str); i++ {
		fmt.Printf("%q\t", str[i])
	}
}

// 遍历字符串
func TestTraverseRune(t *testing.T) {
	str := "陈濛abc" // utf编码中，中文3字节，英文1字节，len(str)表示字节数，所以是9

	// 遍历UTF-8字符
	for _, runeValue := range str {
		fmt.Printf("%#U\t", runeValue)
	}
	fmt.Println()

	for _, runeValue := range str {
		fmt.Printf("%q", runeValue)
	}
	fmt.Println()

	for _, runeValue := range str {
		fmt.Print(string(runeValue))
	}
	fmt.Println()

	//另一种遍历UTF-8字符的方法，可以指定其他编码方式来解码byte[]
	for i, w := 0, 0; i < len(str); i += w {
		runeValue, width := utf8.DecodeRuneInString(str[i:]) // 从str的i处开始，解码出第一个合法utf-8字符，返回该字符和字符所占字节数
		fmt.Printf("%#U\t", runeValue)
		w = width
	}
}

func TestStr2Rune(t *testing.T) {
	str := "陈濛abc"
	rs := []rune(str) // str解码成utf-8字符数组，len(rs) == 5
	for i := 0; i < len(rs); i++ {
		fmt.Printf("%#U\t", rs[i])
	}
}

func TestSplit(t *testing.T) {
	str := "127.0.0.1:8080"
	strSplit := strings.Split(str, ":")
	for i := 0; i < len(strSplit); i++ {
		fmt.Printf("%s\t", strSplit[i])
	}

	str = "a b cc    dd  "
	// 按照空白分割字符串
	strSplit = strings.Fields(str)
	fmt.Println(strSplit)
}

func TestReplace(t *testing.T) {
	str := "127.0.0.1:8080"
	str = strings.Replace(str, ":", "/", -1)
	fmt.Println(str)
}

func TestContains(t *testing.T) {
	str := "127.0.0.1:8080"
	b := strings.Contains(str, "8080")
	fmt.Println(b)

	// chars是否有字符在str中（任意一个即可）
	b2 := strings.ContainsAny(str, "189")
	fmt.Println(b2)
}

func TestEqual(t *testing.T) {
	str := "1"
	str2 := "1"

	if str == str2 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

func TestStrIndex(t *testing.T) {
	str := "陈:8080"
	// 注意 index是字节数组的下标，而不是字符数组的下标
	// 第一个匹配项
	index := strings.Index(str, "8080")
	fmt.Println(index)

	// 最后匹配项
	str = "abcbc"
	index = strings.LastIndex(str, "bc")
	fmt.Println(index)

	// 匹配任一字符
	index = strings.IndexAny("abc", "bc")
	fmt.Println(index)
}

func TestStrJoin(t *testing.T) {
	s := strings.Join([]string{"a", "b", "c"}, "/")
	fmt.Println(s)
}

// 字符映射修改
func TestStrMap(t *testing.T) {
	f := func(r rune) rune {
		switch r {
		case '爸':
			return 'b'
		case '妈':
			return 'm'
		default:
			return r
		}
	}

	s := strings.Map(f, "他妈的，爸啊")
	fmt.Println(s)
}

func TestStr2Int(t *testing.T) {
	i, _ := strconv.Atoi("12")
	fmt.Println(i)
}

func TestInt2Str(t *testing.T) {
	s := strconv.Itoa(12)
	fmt.Println(s)
}

func TestIntPlusStr(t *testing.T) {
	fmt.Println(strconv.Itoa(1) + ":")
	a := 1
	fmt.Println(strconv.Itoa(a) + ":")
}

func TestRawStr(t *testing.T) {
	b := "abc"
	a := fmt.Sprintf(`"hello" 'world' %s`, b)
	fmt.Println(a)
}

func TestStrConv(t *testing.T) {
	s, _ := strconv.ParseFloat("abc", 64)
	fmt.Println(s)
}
