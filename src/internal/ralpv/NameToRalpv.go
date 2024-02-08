package ralpv

import (
	"log"
)

// Ralpv (reserve alphanumeric (and dash) value) to name.
// Similar to hexadecimal numbers, but with alphanumeric and left-to-right values
func NameToRalpv(name string) int64 {
	if name == "" {
		return -1
	} else {
		var ralpv int64 = 0
		var rank int64 = 1
		s := name

		// Loop until ralpv calculation is complete
		for {
			leftMost := s[0:1]

			if v, ok := alpnumdashToDecimal[leftMost]; ok {
				ralpv = ralpv + rank*int64(v)
				rank *= int64(len(Alpnumdash))
			} else {
				log.Fatalln("Not found in letters:", leftMost, name)
			}

			rest := s[1:]
			if rest == "" {
				break
			} else {
				// Continue with rest of word
				s = rest
			}
		}

		return ralpv
	}
}
