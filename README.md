# interval-merge
A simple function to merge a list of integer intervals. Two intervals are only merged when they have an overlap. Intervals without overlap should not be merged.

Example:
Input: [25,30] [2,19] [14, 23] [4,8]  Output: [2,23] [25,30]


Assumptions:
- Input is a Array with only positive Integers. If the Input could vary, extended tests would be necessary
- The Integers are below 100
- The first Integer is smaller than the second


# todo
## Datagenerator
- [x] create sample data generator
- [x] create random test data generator

## Merge Function
- [x] sort the intervals ascending for their first entry
- [x] check between two intervals if they overlap
- [x] take the min and max value and store it in a new interval

## other
- [x] Testing
- Logging