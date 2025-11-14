package piscine

func ReverseMenuIndex(menu []string) []string {
	result := make([]string, len(menu))

	for i, item := range menu {

		reversedIndex := len(menu) - 1 - i

		result[reversedIndex] = item
	}
	return result
}
