package common

func Zip[T any](lists ...[]T) []T {
	numLists := len(lists)
	if numLists == 0 {
		panic("Error in Zip: must provide at least 1 list")
	}
	listLength := len(lists[0])
	for _, list := range lists {
		if len(list) != listLength {
			panic("lists must be the same length")
		}
	}
	zipped := make([]T, 0, listLength*numLists)
	for i := 0; i < listLength; i++ {
		for _, list := range lists {
			zipped = append(zipped, list[i])
		}
	}
	return zipped
}

func Unzip[T any](l []T, n int) [][]T {
	if len(l)%n != 0 {
		panic("Cannot unzip list unless elements are a multiple of n")
	}
	unzipped := make([][]T, n)
	for i, e := range l {
		dimension := i % n
		unzipped[dimension] = append(unzipped[dimension], e)
	}
	return unzipped
}
