package logging

import (
	"fmt"
)

func Log(strings ...interface{}) {
	fmt.Println(strings...)
}