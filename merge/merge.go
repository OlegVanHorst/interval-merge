package merge

// Simple sort function. The Interval is sorted ascending for the first value using BubbleSort.
func BubbleSort(rawData [][]int) [][]int {
	for i := 0; i < len(rawData); i++ {
		for j := 0; j < len(rawData)-i-1; j++ {
			if rawData[j][0] > rawData[j+1][0] {
				rawData[j], rawData[j+1] = rawData[j+1], rawData[j]
			}
		}
	}
	return rawData
}

func Merge(sortedData [][]int) [][]int {
	mergedData := sortedData

	// take two intervals, check if they overlap and write to mergedData. First Value is nul which must be filtered. Could be made better...
	for k := 0; k < len(sortedData)-1; k++ {
		if checkInterval(sortedData[k], sortedData[k+1]) {
			mergedData[k], mergedData[k+1] = nil, []int{sortedData[k][0], max(sortedData[k][1], sortedData[k+1][0])}
		}
	}
	//Remove nil values
	mergedData = removeNils(mergedData)

	return mergedData
}

// Checks if
func checkInterval(valA []int, valB []int) bool {
	if valA[1] >= valB[0] {
		return true
	} else {
		return false
	}
}

// Outputs the max value of two Integers
func max(valA int, valB int) int {
	if valA > valB {
		return valA
	} else {
		return valB
	}
}

// Remove nil values from a [][]int Array
func removeNils(input [][]int) [][]int {
	var output [][]int
	var val []int

	for _, val = range input {
		switch val != nil {
		case true:
			output = append(output, val)
		}
	}
	return output
}
