package lib

import "fmt"

func Panic(str string, args ...interface{}) {
	s := fmt.Sprintf(str, args...)
	panic(s)
}
