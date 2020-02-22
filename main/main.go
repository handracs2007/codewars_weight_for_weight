package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

type MyString struct {
	str string
}

func (obj MyString) Weight() int {
	var w = 0

	for _, ch := range obj.str {
		w += int(ch - '0')
	}

	return w
}

func OrderWeight(strng string) string {
	var myStrs = make([]MyString, 0)
	var buffer bytes.Buffer

	if len(strng) > 0 {
		for _, str := range strings.Split(strng, " ") {
			if len(str) > 0 {
				myStrs = append(myStrs, MyString{str})
			}
		}
	}

	sort.Slice(myStrs, func(i, j int) bool {
		w1 := myStrs[i].Weight()
		w2 := myStrs[j].Weight()

		if w1 == w2 {
			return myStrs[i].str < myStrs[j].str
		}

		return w1 < w2
	})

	for _, myStr := range myStrs {
		buffer.WriteString(myStr.str + " ")
	}

	return strings.TrimSpace(buffer.String())
}

func main() {
	fmt.Println(OrderWeight("103 123 4444 99 2000"))
	fmt.Println(OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123"))
	fmt.Println(OrderWeight(""))
}
