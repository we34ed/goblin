package documentation

// код функции
func codew(code map[string][]int) {
	for key, value := range code {
		counter := 0.0
		for _, element := range value {
			counter += float64(element)
		}
		fmt.Println(key, counter/float64(len(value)))
	}
}

// мэйн
func main() {
	mapp := map[string][]int{
		"stas":  {2, 5, 5, 5, 4},
		"vanya": {3, 4, 4, 3, 5},
		"danya": {4, 4, 4, 5, 5},
		"sanya": {2, 5, 5, 5, 4},
	}
	codew(mapp)
}
