package util

import (
	"fmt"
)

func Pid2Pkg(id int) string {
	return fmt.Sprintf("p%03d", id)
}

func DefaultCode(id int, lang string) (string, error) {
	return "", nil
}

func DefaultTestCode(id int, code string, lang string) (string, error) {
	return "", nil
}
