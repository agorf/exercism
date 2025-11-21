package binarysearch

func SearchInts(list []int, key int) int {
	if len(list) == 0 {
		return -1
	}
	mid := len(list) / 2
	val := list[mid]
	switch {
	case key == val:
		return mid
	case key < val:
		return SearchInts(list[0:mid], key)
	case key > val:
		index := SearchInts(list[mid+1:], key)
		if index == -1 {
			return -1
		}
		return mid + 1 + index
	}
	panic("should never happen")
}
