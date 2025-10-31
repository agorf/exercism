package isbn

func IsValidISBN(isbn string) bool {
	sum := 0
	length := 0
	for _, r := range isbn {
		if r == '-' { // Ignore dashes
			continue
		}
		if length == 10 { // ISBNs longer than 10 characters are invalid
			return false
		}
		d := int(r - '0')
		if r == 'X' {
			if length < 9 { // ISBNs may contain an X only at the end as a check digit
				return false
			}
			d = 10
		}
		sum += d * (10 - length)
		length++
	}
	return length == 10 && sum%11 == 0
}
