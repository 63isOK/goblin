package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

const LocalOSSPath = "./assets/"

func main() {
	join()
	getPath()
	quoted()
}

func getPath() {
	fileName := "iot/device/a.jpg"

	savePath := LocalOSSPath + strings.TrimPrefix(fileName, "/")

	path := strings.TrimSuffix(savePath, (strings.Split(savePath, "/")[len(strings.Split(savePath, "/"))-1]))

	fmt.Println(savePath)
	fmt.Println(path)
}

type people struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
	Addr string `json:"addr"`
	Home Home   `json:"home"`
}

type Home struct {
	Addr string `json:"addr"`
	Size string `json:"size"`
}

func quoted() {
	s := "{\"abc\":\"dd\"}"
	fmt.Printf("%q", s)
	fmt.Println()
	fmt.Printf("%s", s)
	fmt.Println()

	p := people{
		Age:  2,
		Name: "aaa",
		Addr: "bbb",
		Home: Home{
			Addr: "abc",
			Size: "ac",
		},
	}

	b, _ := json.Marshal(p)
	fmt.Printf("b: %q", string(b))
	fmt.Println()
	fmt.Printf("b: %s", string(b))
	fmt.Println()

	q := fmt.Sprintf("%q", string(b))
	fmt.Println("cc: ", q)

	r, _ := strconv.Unquote(q)
	fmt.Println("r: ", r)

	a := "{\"body\":\"{  \\\"serviceCode\\\":\\\"ET040404\\\",  \\\"projectCode\\\":\\\"aaa\\\",  \\\"source\\\":\\\"17\\\",  \\\"content\\\":\\\"【${通道}】【${设备}】断网离线，请排查以免影响车辆通行！\\\",  \\\"appointmentBeginTime\\\":\\\"${当前时间}\\\",  \\\"appointmentEndTime\\\":\\\"${当前时间}\\\"}\",\"signature\":\"jindi\",\"url\":\"http://baidu.com\"}"
	fmt.Println(a)

	// r, err := strconv.Unquote(a)
	// if err != nil {
	//   panic(err)
	// }
	// fmt.Println("r: ", r)

	rawa, ok := unquoteBytes([]byte(a))
	if !ok {
		fmt.Println("json格式解双引号失败")
		return
	}

	fmt.Println(rawa)
}

func unquoteBytes(s []byte) (t []byte, ok bool) {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return
	}
	s = s[1 : len(s)-1]

	// Check for unusual characters. If there are none,
	// then no unquoting is needed, so return a slice of the
	// original bytes.
	r := 0
	for r < len(s) {
		c := s[r]
		if c == '\\' || c == '"' || c < ' ' {
			break
		}
		if c < utf8.RuneSelf {
			r++
			continue
		}
		rr, size := utf8.DecodeRune(s[r:])
		if rr == utf8.RuneError && size == 1 {
			break
		}
		r += size
	}
	if r == len(s) {
		return s, true
	}

	b := make([]byte, len(s)+2*utf8.UTFMax)
	w := copy(b, s[0:r])
	for r < len(s) {
		// Out of room? Can only happen if s is full of
		// malformed UTF-8 and we're replacing each
		// byte with RuneError.
		if w >= len(b)-2*utf8.UTFMax {
			nb := make([]byte, (len(b)+utf8.UTFMax)*2)
			copy(nb, b[0:w])
			b = nb
		}
		switch c := s[r]; {
		case c == '\\':
			r++
			if r >= len(s) {
				return
			}
			switch s[r] {
			default:
				return
			case '"', '\\', '/', '\'':
				b[w] = s[r]
				r++
				w++
			case 'b':
				b[w] = '\b'
				r++
				w++
			case 'f':
				b[w] = '\f'
				r++
				w++
			case 'n':
				b[w] = '\n'
				r++
				w++
			case 'r':
				b[w] = '\r'
				r++
				w++
			case 't':
				b[w] = '\t'
				r++
				w++
			case 'u':
				r--
				rr := getu4(s[r:])
				if rr < 0 {
					return
				}
				r += 6
				if utf16.IsSurrogate(rr) {
					rr1 := getu4(s[r:])
					if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
						// A valid pair; consume.
						r += 6
						w += utf8.EncodeRune(b[w:], dec)
						break
					}
					// Invalid surrogate; fall back to replacement rune.
					rr = unicode.ReplacementChar
				}
				w += utf8.EncodeRune(b[w:], rr)
			}

		// Quote, control characters are invalid.
		case c == '"', c < ' ':
			return

		// ASCII
		case c < utf8.RuneSelf:
			b[w] = c
			r++
			w++

		// Coerce to well-formed UTF-8.
		default:
			rr, size := utf8.DecodeRune(s[r:])
			r += size
			w += utf8.EncodeRune(b[w:], rr)
		}
	}
	return b[0:w], true
}

// getu4 decodes \uXXXX from the beginning of s, returning the hex value,
// or it returns -1.
func getu4(s []byte) rune {
	if len(s) < 6 || s[0] != '\\' || s[1] != 'u' {
		return -1
	}
	var r rune
	for _, c := range s[2:6] {
		switch {
		case '0' <= c && c <= '9':
			c = c - '0'
		case 'a' <= c && c <= 'f':
			c = c - 'a' + 10
		case 'A' <= c && c <= 'F':
			c = c - 'A' + 10
		default:
			return -1
		}
		r = r*16 + rune(c)
	}
	return r
}

func join() {
	println("============join==============")

	strs := []string{"a", "b", "c"}
	println("strings.Join(a,b,c)", strings.Join(strs, ","))

	strs = []string{"a"}
	println("strings.Join(a,b,c)", strings.Join(strs, ","))
}
