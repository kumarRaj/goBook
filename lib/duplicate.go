package duplicate

import (
	"bufio"
	"fmt"
	"os"
)

//Duplicate counts the text duplicates on the standard input
func Duplicate() {
	textMap := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			break
		}
		textMap[scanner.Text()]++
	}
	for line := range textMap {
		fmt.Println(line, textMap[line])
	}
}
