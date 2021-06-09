package helpers

import (
	"errors"
	"fmt"
	"strings"
)

// CheckStringMinChar 최소 글자 수를 만족하는지
func CheckStringMinChar(s string, minCharNum int) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return errors.New("empty string") // TODO: 나중에 errors 디렉토리로 옮겨서 처리
	}

	if len(s) < minCharNum {
		return errors.New(fmt.Sprintf("must be at least %d characters long", minCharNum)) // TODO: 얘도 나중에 errors로
	}

	return nil
}
