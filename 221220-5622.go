package algorithm

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func contains(e rune)int{
	alphabet := []rune{'A','D','G','J','M','P','T','W'}
	if e < 'A' {return -3}

	for i:=0;i<len(alphabet)-1;i++{
		if alphabet[i] <= e && e<alphabet[i+1]{
			return i
		}
		if 'W'<=e && e<'Z'{
			return 7
		}
	}
	return -3
}

func main() {
	var str string
	var time int

	stdin := bufio.NewReader(os.Stdin)
	_, err := fmt.Scan(&str)
	str = strings.ToUpper(str)

	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		for i := 0; i < len(str); i++ {
			time += contains(rune(str[i])) + 3
		}
		fmt.Println(time)
	}
}