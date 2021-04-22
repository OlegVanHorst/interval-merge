package merge

import (
	"testing"

	"example.com/interval-merge/dataGenerator"
)

// Test the Merge function with 10000 Random numbers between 0 & 10000
func TestMerge(t *testing.T) {
	sortedData := dataGenerator.RandomDataGenerator(10000, 10000)
	mergedData := Merge(sortedData)

	for i := 0; i < len(mergedData)-1; i++ {
		if mergedData[i][1] > mergedData[i+1][0] {
			t.Errorf("Interval not correct, got: %v", mergedData)
		}
	}
}
