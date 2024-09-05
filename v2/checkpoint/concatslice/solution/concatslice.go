package solution

func ConcatSlice(slice1, slice2 []int) []int {
	var result []int

	result = append(result, slice1...)
	result = append(result, slice2...)
	
	return result
}