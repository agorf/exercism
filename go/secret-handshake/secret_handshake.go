package secret

const (
	wink          = 0b1
	doubleBlink   = 0b10
	closeYourEyes = 0b100
	jump          = 0b1000
	reverse       = 0b10000
)

func reverseSlice(s []string) []string {
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		s[i], s[sLen-1-i] = s[sLen-1-i], s[i] // Swap
	}
	return s
}

func Handshake(code uint) []string {
	var words []string
	if code&wink == wink {
		words = append(words, "wink")
	}
	if code&doubleBlink == doubleBlink {
		words = append(words, "double blink")
	}
	if code&closeYourEyes == closeYourEyes {
		words = append(words, "close your eyes")
	}
	if code&jump == jump {
		words = append(words, "jump")
	}
	if code&reverse == reverse {
		words = reverseSlice(words)
	}
	return words
}
