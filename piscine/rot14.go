package piscine

func Rot14(s string) string {
	word := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			if char > 'L' {
				gap := 'Z' - char
				remainder := 14 - gap
				word += string(64 + remainder)
				continue
			}
			word += string(char + 14)
		} else if char >= 'a' && char <= 'z' {
			if char > 'l' {
				gap := 'z' - char
				remainder := 14 - gap
				word += string(96 + remainder)
				continue
			}
			word += string(char + 14)
		} else {
			word += string(char)
		}
	}
	return word
}
