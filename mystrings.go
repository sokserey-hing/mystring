package mystring

// Reverse a string letf to right
// Remember to always capitalize the first letter of the function name if you want to use it outside of the package

func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return string(result)
}