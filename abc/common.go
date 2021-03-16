package abc

import (
	"fmt"
	"os"
)

func Pwd() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
}
