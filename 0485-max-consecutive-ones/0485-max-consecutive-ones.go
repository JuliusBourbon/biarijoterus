func findMaxConsecutiveOnes(nums []int) int {
    count := 0
    max := 0
    n := len(nums)

    for i := 0; i < n; i++{
        pointer := nums[i]
        if pointer == 1 {
            count += 1
        } else {
            count = 0
        }

        if max < count {
            max += 1
        }
    }

    return max
}