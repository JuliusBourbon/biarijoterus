func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
    max := nums[0]
    
    // max used for creating a new array with maximum value in nums (in case nums length != maximum value)
    for i := 0; i < n-1; i++ {
        if nums[i+1] > max {
            max = nums[i+1]
        }

        // in case max value is lower than nums length
        if max < n {
            max = n
        }
    }

    // find duplicates while also found missing number
    duplicated := make([]int, max)
    missing := make([]int, 0, max)
    for _, j := range nums {
        duplicated[j-1]++
    }

    for k := 0; k < max; k++{
        if duplicated[k] == 0 {
            missing = append(missing, k+1)
        }
    }

    return missing
}