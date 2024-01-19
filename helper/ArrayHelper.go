package helper

func MakeIndex(index int, count int) int {

	result := 0
	if index >= 0 {
		result = index % count
		return result
	}
	result = count + index
	return result
}
