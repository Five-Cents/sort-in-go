Some sorting algorithms implemented in golang. Currently, []int is the only accepted type.

Run tests with `go test`.

Run benchmarks with `go test -bench=.`

### Algorithms:

| Algorithm | Best | Worst | Space |
| ------------- |:-------------:| -----:| -----:|
| Bubble Sort | O(n) | O(n^2) | O(1)  |
| Insertion Sort | O(n) | O(n^2) | O(1)  |
| Selection Sort | O(n^2) | O(n^2) | O(1)  |
| Merge Sort | O(n log n) | O(n log n) | O(n log n)  |
| Quick Sort | O(n log n) | O(n^2) | O(n) / O(log n) avg. |