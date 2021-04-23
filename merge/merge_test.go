package merge

import (
	"testing"

	"example.com/interval-merge/dataGenerator"
)

// Interval: [[x1,x2],[y1,y2]]
// When the data is sorted and merged, x2 must be smaller than y1

// Test the Merge function with the example data
func TestMerge(t *testing.T) {
	sortedData := dataGenerator.StaticDataGenerator()
	mergedData := Merge(sortedData)

	for i := 0; i < len(mergedData)-1; i++ {
		if mergedData[i][1] > mergedData[i+1][0] {
			t.Errorf("Interval not correct, got: %v", mergedData)
		}
	}
}

// Test the Merge function with 10000 Random numbers between 0 & 10000
func TestRandomMerge(t *testing.T) {
	sortedData := dataGenerator.RandomDataGenerator(10000, 10000)
	mergedData := Merge(sortedData)

	for i := 0; i < len(mergedData)-1; i++ {
		if mergedData[i][1] > mergedData[i+1][0] {
			t.Errorf("Interval not correct, got: %v", mergedData)
		}
	}
}
