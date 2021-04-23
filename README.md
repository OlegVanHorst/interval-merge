# interval-merge
A simple function to merge a list of integer intervals. Two intervals are only merged when they have an overlap. Intervals without overlap should not be merged.

## Example:
Input: [25,30] [2,19] [14, 23] [4,8]  Output: [2,23] [25,30]


## Assumptions:
- Input is a Array with only positive Integers. If the Input could vary, extended tests would be necessary
- The first Integer is always smaller than the second

# Build 
The Application is written in go. Install golang via https://golang.org/doc/install#install on your system if necessary.
## Clone the repository
```
git clone https://github.com/OlegVanHorst/interval-merge.git
```
## Make
The Application can be build with the make command. A simple test of the merge function with random numbers is executed and the binary will be created in the build folder
```
make
```

## Execute
Either via binary or with go run
```
./build/merge-interval
go run main.go
```

# Runtime of the Merge package
The runtime of the Application is mostly determined by the BubbleSort Algorithm. 
- Best Case: O(n)
- Worst Case: O(n²)

The other loops add each a runtime of O(n)
- removeNils
- Loop around checkInterval

# Robustness
The data input should be checked for the correct format. 
Great inputs could either be declined, or split into multiple inputs. All results must be merged in a second step. (Only works when intervals overlap)


# Memory Usage
The stats were determined by the merge_test:
```
cd merge
go test -memprofile mem.prof -bench .
```

The total memory of the Merge function is at around 1.7 mb with a Array of length 10000
```
      flat  flat%   sum%        cum   cum%
 1743.85kB 50.69% 50.69%  1743.85kB 50.69%  example.com/interval-merge/dataGenerator.RandomDataGenerator
 1184.27kB 34.42% 85.11%  1184.27kB 34.42%  runtime/pprof.StartCPUProfile
  512.44kB 14.89%   100%   512.44kB 14.89%  flag.(*FlagSet).Var
         0     0%   100%  1743.85kB 50.69%  example.com/interval-merge/merge.TestRandomMerge
         0     0%   100%   512.44kB 14.89%  flag.(*FlagSet).String
         0     0%   100%   512.44kB 14.89%  flag.(*FlagSet).StringVar (inline)
         0     0%   100%   512.44kB 14.89%  flag.String (inline)
         0     0%   100%  1696.71kB 49.31%  main.main
         0     0%   100%  1696.71kB 49.31%  runtime.main
         0     0%   100%  1184.27kB 34.42%  testing.(*M).Run
```
