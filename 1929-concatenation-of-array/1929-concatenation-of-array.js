/**
 * @param {number[]} nums
 * @return {number[]}
 */
var getConcatenation = function(nums) {
    n = nums.length;
    ans = [...nums, ...nums];

    return ans;
};

getConcatenation([1, 2, 3])
getConcatenation([1, 3, 2, 1])