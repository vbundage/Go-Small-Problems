// https://www.codewars.com/kata/59df2f8f08c6cec835000012/go

import (
	"fmt"
	"sort"
	"strings"
)

func splitNames(names string) (firstName string, lastName string) {
	firstName, lastName = strings.Split(names, ":")[0], strings.Split(names, ":")[1]
	return
}

func Meeting(s string) string {
	names := strings.Split(strings.ToUpper(s), ";")
	for idx, name := range names {
		firstName, lastName := splitNames(name)
		names[idx] = fmt.Sprintf("(%s, %s)", lastName, firstName)
	}
	sort.Strings(names)
	return strings.Join(names, "")
}

func main() {
	testStr1 := "Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn"
	testStr2 := "John:Gates;Michael:Wahl;Megan:Bell;Paul:Dorries;James:Dorny;Lewis:Steve;Alex:Meta;Elizabeth:Russel;Anna:Korn;Ann:Kern;Amber:Cornwell"
	fmt.Println(Meeting(testStr1) == "(ARNO, ANN)(BELL, JOHN)(CORNWELL, ALEX)(DORNY, ABBA)(KERN, LEWIS)(KORN, ALEX)(META, GRACE)(SCHWARZ, VICTORIA)(STAN, MADISON)(STAN, MEGAN)(WAHL, ALEXIS)")
	fmt.Println(Meeting(testStr2) == "(BELL, MEGAN)(CORNWELL, AMBER)(DORNY, JAMES)(DORRIES, PAUL)(GATES, JOHN)(KERN, ANN)(KORN, ANNA)(META, ALEX)(RUSSEL, ELIZABETH)(STEVE, LEWIS)(WAHL, MICHAEL)")
}
