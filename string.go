package toolbox

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func TrimString(s string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(s, ""), nil
}

func HasStrings(str string, ss ...string) (bool, int) {
	matches := 0
	isCompleteMatch := true

	for _, sub := range ss {
		if strings.Contains(str, sub) {
			matches += 1
		} else {
			isCompleteMatch = false
		}
	}
	return isCompleteMatch, matches
}

func HasString(ss []string, str string) bool {
	for _, v := range ss {
		if v == str {
			return true
		}
	}
	return false
}

func LastString(ss []string) string {
	return ss[len(ss)-1]
}

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func AllSameStrings(a []string) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
