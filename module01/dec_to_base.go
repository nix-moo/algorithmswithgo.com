package module01

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"
func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEF"
	var res string
	for dec != 0 {
		i := dec % base
		// MULTIPLE WAYS TO PRINT
		//  %X will format to base 16
		res = fmt.Sprintf("%X%s", i, res)
		// OR
		// You cna create your own charset
		res = string(charset[i]) + res

		dec = dec / base
	}

	return res
}
