package merge

// Merge function sorts the Array with Bubblesort,
// checks each Interval Tuple if it can be merged and writes the merge interval to mergedArray
// Removes all nil values
func Merge(inputData [][]int) [][]int {
	sortedData := bubbleSort(inputData)
	mergedData := sortedData

	// Takes two intervals, check if they overlap and write to mergedData:
	// First Value is nil which must be filtered.
	for k := 0; k < len(sortedData)-1; k++ {
		if checkInterval(sortedData[k], sortedData[k+1]) {
			mergedData[k], mergedData[k+1] = nil, []int{sortedData[k][0], max(sortedData[k][1], sortedData[k+1][0])}
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
