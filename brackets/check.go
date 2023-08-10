package main

func check(s string) string {
	for i, v := range s {
		if v == 41 || v == 93 || v == 125 { //---> Isolated Closing Brackets
			return "Error"
		}
		switch v {
		case 40:
			return close(s[i+1:], 41)
		case 91, 123:
			return close(s[i+1:], v+2)
		}
	}
	return "OK" //---> No Brackets
}
