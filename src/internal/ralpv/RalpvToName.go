package ralpv

import "strings"

// Name to ralpv (reserve alphanumeric value).
// Similar to hexadecimal numbers, but with alphanumeric and left-to-right values
func RalpvToName(ralpv int64) string {
	if ralpv < 0 {
		return ""
	} else {
		// Base of Ralphv
		base := len(alphanumicToDecimal)
		var nameBuilder strings.Builder
		dividend := ralpv

		// Loop until name generation is complete
		for {
			remainder := dividend % int64(base)

			for k, v := range alphanumicToDecimal {
				if int64(v) == remainder {
					nameBuilder.WriteString(k)
					break
				}
			}

			if dividend < int64(base) {
				break
			} else {
				// Continue with next letter
				dividend = dividend / int64(base)
			}
		}

		return nameBuilder.String()
	}
}
