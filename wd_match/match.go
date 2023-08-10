package main

func match(wd, str string) bool {
	if len(wd) == 1 {
		return contain(str, wd)
	}
	for i, v := range wd {
		for j, w := range str {
			if v == w {
				return match(wd[i+1:], str[j+1:])
			}
		}
	}
	return false
}
