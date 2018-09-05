package brackets

func Bracket(s string) (bool, error) {

	output, err := BraceCounting(s)

	return output, err
}

func BraceCounting(s string) (bool, error) {
	var str []rune
	for _, symbol := range s {
		switch symbol {
		case '{', '[', '(':
			str = append(str, symbol)

		case '}', ']', ')':
			if len(str) > 0 {
				if symbol == '}' && str[len(str)-1] == '{' {
					str = str[:len(str)-1]
				} else if symbol == ']' && str[len(str)-1] == '[' {
					str = str[:len(str)-1]
				} else if symbol == ')' && str[len(str)-1] == '(' {
					str = str[:len(str)-1]
				} else {
					return false, nil
				}

			}
		}
	}
	if len(str) == 0 {
		return true, nil
	}
	return false, nil
}
