// https://www.codewars.com/kata/550f22f4d758534c1100025a/go

import (
	"fmt"
	"regexp"
	"strings"
)

func DirReduc(arr []string) []string {
	str := strings.Join(arr, " ")
	regex := regexp.MustCompile(`NORTH SOUTH|SOUTH NORTH|EAST WEST|WEST EAST`)
	for {
		length := len(str)
		str = strings.TrimSpace(regex.ReplaceAllString(str, ""))
		str = strings.ReplaceAll(str, "  ", " ")
		if length == len(str) {
			break
		}
	}
	if str == "" {
		return []string{}
	}
	return strings.Split(str, " ")
}

func main() {
	var a = []string{"NORTH", "SOUTH", "EAST", "WEST"}
	fmt.Println(DirReduc(a)) // []
	a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
	fmt.Println(DirReduc(a)) // [WEST]
	a = []string{"NORTH", "WEST", "SOUTH", "EAST"}
	fmt.Println(DirReduc(a)) // [NORTH WEST SOUTH EAST]
	a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}
	fmt.Println(DirReduc(a)) // [NORTH]
}
