// https://www.codewars.com/kata/554e4a2f232cdd87d9000038/go

package kata

func DNAStrand(dna string) (result string) {
	var dnaMap = map[string]string{
		"A": "T",
		"T": "A",
		"C": "G",
		"G": "C",
	}

	for _, key := range dna {
		result += dnaMap[string(key)]
	}
	return
}

// import "strings"

// func DNAStrand(dna string) string {
// 	replacer := strings.NewReplacer("A", "T", "T", "A", "C", "G", "G", "C")
// 	return replacer.Replace(dna)
// }
