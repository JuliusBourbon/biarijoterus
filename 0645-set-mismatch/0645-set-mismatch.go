func findErrorNums(nums []int) []int {
    n := len(nums)
    totalDuplicate := make([]int, n+1) // n+1 so n number will be save on n index

    // blank identifier (_) to store the index of loop, but it will never used
    // num is for the value
    for _, num := range nums {
        totalDuplicate[num]++
        // same as :
        // totalDuplicate[num] = totalDuplicate[num] + 1
    }

    duplicate := 0
    missing := 0

    for i := 1; i <= n; i++ {
        if totalDuplicate[i] == 2 {
            duplicate = i
        } else if totalDuplicate[i] == 0 {
            missing = i
        }
    }

    return []int{duplicate, missing}
}