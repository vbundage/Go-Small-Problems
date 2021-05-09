// https://www.codewars.com/kata/54dc6f5a224c26032800005c/go

import (
	"fmt"
	"strconv"
	"strings"
)

func makeCategories(categories []string) map[string]int {
	categoryMap := map[string]int{}
	for _, category := range categories {
		categoryMap[category] = 0
	}

	return categoryMap
}

func mapQuantities(categoryMap map[string]int, stocklist []string) map[string]int {
	for _, book := range stocklist {
		quantity, _ := strconv.Atoi(strings.Fields(book)[1])
		_, exists := categoryMap[string(book[0])]
		if exists {
			categoryMap[string(book[0])] += quantity
		}
	}

	return categoryMap
}

func formatCategories(categories []string, categoryMap map[string]int) string {
	strSlice := []string{}
	for _, cat := range categories {
		strSlice = append(strSlice, fmt.Sprintf("(%s : %d)", cat, categoryMap[cat]))
	}

	return strings.Join(strSlice, " - ")
}

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	categoryMap := makeCategories(listCat)
	categoryMap = mapQuantities(categoryMap, listArt)
	return formatCategories(listCat, categoryMap)
}

func main() {
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}
	fmt.Println(StockList(b, c) == "(A : 0) - (B : 1290) - (C : 515) - (D : 600)")

	b = []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"}
	c = []string{"A", "B"}
	fmt.Println(StockList(b, c) == "(A : 200) - (B : 1140)")
}