package main

import (
	"fmt"
	"strings"
)

func StringIncreases(i string, origin string) string {
	if i == "" {
		return ""
	}

	var carryNum []byte
	if origin == "A" {
		carryNum = []byte(origin)
	} else {
		carryNum = []byte("0")
	}

	target := strings.ToUpper(i)
	targetby := []byte(target)

	l := len(targetby)
	newby := make([]byte, l)
	carry := false

	targetby[(l-1)]++
	for i := (l - 1); i >= 0; i-- {
		if carry {
			targetby[i]++
		}

		if targetby[i] > 90 {
			newby[i] = carryNum[0]
			carry = true
		} else {
			newby[i] = targetby[i]
			carry = false
		}
	}

	var ret string
	ret = string(newby)
	if carry {
		if origin == "A" {
			ret = "A" + ret
		} else {
			ret = "1" + ret
		}
	}

	return ret
}

func main() {
	//字符串转[]byte(string) ascii码
	by := []byte("AZaz")
	fmt.Println("byte=", by) //每个字符的ascii码

	target := "AZZZ10"
	target = strings.ToUpper(target)
	targetby := []byte(target)
	fmt.Println("byte=", target, targetby)      //每个字符的ascii码
	fmt.Println("len(targetby)", len(targetby)) //每个字符的ascii码
	l := len(targetby)
	newby := make([]byte, l)
	carry := false

	targetby[(l-1)]++
	for i := (l - 1); i >= 0; i-- {
		fmt.Println(targetby[i], carry)
		if carry {
			targetby[i]++
		}

		if targetby[i] > 90 {
			newby[i] = 65
			carry = true
		} else {
			newby[i] = targetby[i]
			carry = false
		}
	}

	fmt.Println("newby:", string(newby))

	fmt.Println(StringIncreases("AZZZ10AZ", "A"))
	fmt.Println(StringIncreases("ZZZZ", "A"))
	fmt.Println(StringIncreases("ZZZZ", "0"))

}
