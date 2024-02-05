package ralpv

import (
	"log"
)

// Ralpv (reserve alphanumeric value) to name.
// Similar to hexadecimal numbers, but with alphanumeric and left-to-right values
func NameToRalpv(name string) int64 {
	if name == "" {
		return -1
	} else {
		var ralpv int64 = 0
		s := name

		// Loop until ralpv calculation is complete
		for {
			leftMost := s[0:1]

			if v, ok := alphanumicToDecimal[leftMost]; ok {
				ralpv = ralpv + int64(v)
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
