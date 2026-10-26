func smallerNumbersThanCurrent(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)

    for i := 0; i < n; i++ {
        current := nums[i]
        
        // compare current with all prev and next
        if i > 0 && i < n-1 {
            for j := i-1; j >= 0; j-- {
                prev := nums[j]

                if current > prev {
                    ans[i]++
                }
            }

            for k := i + 1; k < n; k++ {
                next := nums[k]

                if current > next {
                    ans[i]++
                }
            }
        }

        // compare current with all prev
        if i == n-1 {
            for j := i-1; j >= 0; j-- {
                prev := nums[j]

                if current > prev {
                    ans[i]++
                }
            }
        } 
        
        // compare current with all next
        if i == 0 {
            for j := 1; j < n; j++ {
                next := nums[j]

                if current > next {
                    ans[i]++
                }
            }
        }

        // i know this code is sucks, at least it passes all the tests :p
    }

    return ans
}