package leetcode

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	//fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome(" "))
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isalnum(s[l]) {
			l++
		}
		for l < r && !isalnum(s[r]) {
			r--
		}
		if l < r {
			if toLower(s[l]) != toLower(s[r]) {
				return false
			} else {
				l++
				r--
			}
		}
	}

	return true
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		ch += 32
	} else {
		return ch
	}
	return ch
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
