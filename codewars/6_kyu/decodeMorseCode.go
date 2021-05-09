// https://www.codewars.com/kata/54b724efac3d5402db00065e/go

package kata

import "strings"

// func DecodeMorse(morseCode string) string {
//   words := strings.Split(strings.Trim(morseCode, " "), "   ")
//   for idx, word := range words {
//     var decodedWord string
//     for _, char := range strings.Split(word, " ") {
//       decodedWord += MORSE_CODE[char]
//     }
//     words[idx] = decodedWord
//   }
//   return strings.Join(words, " ")
// }

var MORSE_CODE = map[string]string{} // placeholder: remove for intended execution

func DecodeMorse(morseCode string) string {
	words := strings.Split(strings.Trim(morseCode, " "), "   ")
	for idx, word := range words {
		var decodedWord strings.Builder
		for _, char := range strings.Split(word, " ") {
			decodedWord.WriteString(MORSE_CODE[char])
		}
		words[idx] = decodedWord.String()
	}
	return strings.Join(words, " ")
}
