# Insertion Sort

```
# given:
arr = [3, 1, 2]

# algo:
for i := 1; i < len(arr), i++ {
    key := arr[i]
    j := i - 1
    for j >= 0 && arr[j] > key {
        arr[j+1] = arr[j]
        j -= 1
    }
    arr[j+1] = key
}

# step by step:
# i = 1
    key = arr[i]    --> 1
    j := 1 - 1      --> 0
# j equals 0, so the while loop will run once
    arr[j+1] = arr[j] --> [3, 3, 2]
# our j loop has finished running, j now equals 0
    arr[j+1] = key  --> [1, 3, 2]
# i = 2
    key := arr[i]   --> 3
    j := i - 1      --> 1
# j equals 1 and 3 (arr[j]) is > 2 (arr[j+1])
    
```

