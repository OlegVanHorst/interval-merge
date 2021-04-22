# interval-merge
A simple function to merge a list of integer intervals. Two intervals are only merged when they have an overlap. Intervals without overlap should not be merged.

Example:
Input: [25,30] [2,19] [14, 23] [4,8]  Output: [2,23] [25,30]



# todo
## Datagenerator
- create sample data generator
- create random test data generator

## Merge Function
- sort the intervals ascending for their first entry
- check between two intervals if they overlap
- take the min and max value and store it in a new interval

## Testing
- Add testing