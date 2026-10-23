func getConcatenation(nums []int) []int {
    // create new array with 0 length and 2*len(nums) capacity
    ans := make([]int, 0, 2*len(nums))

    // append(destination, element)
    ans = append(ans, nums...,)
    ans = append(ans, nums...,)
    return ans
}