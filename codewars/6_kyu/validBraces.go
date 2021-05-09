// https://www.codewars.com/kata/5277c8a221e209d3f6000b56/go

type Stack struct {
	data []string
}

func (s *Stack) push(elem string) {
	s.data = append(s.data, elem)
}

func (s *Stack) pop() string {
	if len(s.data) == 0 {
		return ""
	}
	removed := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return removed
}

func (s Stack) read() string {
	if len(s.data) == 0 {
		return ""
	}
	return s.data[len(s.data)-1]
}

func ValidBraces(str string) bool {
	braceTable := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}

	stack := Stack{}
	for _, charCode := range str {
		char := string(charCode)
		if char == "(" || char == "{" || char == "[" {
			stack.push(char)
		} else if char == ")" || char == "}" || char == "]" {
			popped := stack.pop()
			if braceTable[popped] != char {
				return false
			}
		}
	}

	return stack.read() == ""
}