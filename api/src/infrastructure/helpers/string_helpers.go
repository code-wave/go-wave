package helpers

import (
	"errors"
	"fmt"
	"strings"
)

// CheckStringMinChar 빈 글자인지 & 최소 글자 수를 만족하는지
func CheckStringMinChar(s string, minCharNum int) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return errors.New("empty string")
	}

	if len(s) < minCharNum {
		return errors.New(fmt.Sprintf("must be at least %d characters long", minCharNum))
	}

	return nil
}

// ConvertStringArray string 배열을 스페이스 삭제 & 소문자로 변환
func ConvertStringArray(s []string) error {
	if len(s) == 0 {
		return errors.New("empty array")
	}

	for idx, val := range s {
		val = strings.TrimSpace(val)
		val = strings.ToLower(val)
		s[idx] = val
	}

	return nil
}
