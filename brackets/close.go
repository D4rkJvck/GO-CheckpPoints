package main

func close(s string, r rune) string {
	for i, v := range s {
		if r == 41 && (v == 91 || v == 123) {
			return "Error"
		}
		if r == 93 && v == 123 {
			return "Error"
		}
		if v == r {
			if i != len(s)-1 && check(s[i+1:]) == "Error" {
				return "Error" //---> Brackets Outside
			}
			return check(s[:i]) //---> Brackets Inside
		}
	}
	return "Error" //---> Isolated Opening Brackets
}
