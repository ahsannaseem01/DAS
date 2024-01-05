package binary_search

func BinarySearch(li []int, tar int) (found bool) {
	low := 0
	high := len(li) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := li[mid]
		if guess == tar {
			found = true
			break
		} else if guess > tar {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return
}

//ihiu
