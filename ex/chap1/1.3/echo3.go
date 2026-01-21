package echo

import (
	"fmt"
	"strings"
)

// func main() {
// 	concat(os.Args[:])
// 	//join(os.Args[:])
// }

func Concat(args []string) string {
	var s string
	for i := 0; i < len(args); i++ {
		s += fmt.Sprintf("%v. %s\n", i, args[i])
		//sep = " "
	}

	return s
}

func Join(args []string) string {
	return strings.Join(args, " ")
}
