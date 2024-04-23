package fn

import (
	"fmt"
	"strings"
)

func DoSmth(str string) {
	fmt.Println(strings.ToUpper(str))
}

func WhoAmI(name string, age int) string {
	me := fmt.Sprintf("I am %s , My Age is %d", name, age)
	fmt.Println(me)
	return me
}
