package gocode

import (
	"fmt"
	"os"
	"strings"
)

func ValidateGitPath(path string) bool {
	if strings.HasSuffix(path, "/") {
		path = path + ".git"
	} else {
		path = path + "/.git"
	}
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		fmt.Println(err)
		return false
	}
	return true
}
