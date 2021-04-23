package merge

// Merge function sorts the Array with Bubblesort,
// checks each Interval Tuple if it can be merged and writes the merge interval to mergedArray
// Removes all nil values
func Merge(inputData [][]int) [][]int {
	// Sort the Data ascending
	sortedData := bubbleSort(inputData)
	mergedData := sortedData

	var valA []int
	var valB []int
	dataLen := len(sortedData)

	// Takes two intervals, check if they overlap and write to mergedData:
	// First Value is nil which must be filtered.
	for k := 0; k < dataLen-1; k++ {
		valA = sortedData[k]
		valB = sortedData[k+1]

		if checkInterval(valA, valB) {
			// when the intervals overlap: nil the first value and set the second as the full interval
			mergedData[k], mergedData[k+1] = nil, []int{valA[0], max(valA[1], valB[1])}
		}
	}

	//Remove nil values
	mergedData = removeNils(mergedData)

	return mergedData
}

// Simple sort function. The Interval is sorted ascending for the first value using BubbleSort. https://en.wikipedia.org/wiki/Bubble_sort
func bubbleSort(rawData [][]int) [][]int {
	for i := 0; i < len(rawData); i++ {
		for j := 0; j < len(rawData)-i-1; j++ {
			if rawData[j][0] > rawData[j+1][0] {
				rawData[j], rawData[j+1] = rawData[j+1], rawData[j]
			}
		}
	}
	return rawData
}

// Takes two intervals [[x1,x2],[y1,y2]] as input and checks if y2>=y1
func checkInterval(valA []int, valB []int) bool {
	if valA[1] >= valB[0] {
		return true
	} else {
		return false
	}
}

// Outputs the max value of two Integers
func max(valA int, valB int) int {
	if valA >= valB {
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
