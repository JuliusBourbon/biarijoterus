/**
 * @param {number[]} nums
 * @param {number} n
 * @return {number[]}
 */
var shuffle = function(nums, n) {
  x = nums.slice(0, n)
  y = nums.slice(n)
  ans = []
  
  for (let i = 0; i < n; i++){
    ans.push(x[i])
    ans.push(y[i])
  }
  
  return ans
};

const test1 = [2,5,1,3,4,7];
const test2 = [1,2,3,4,4,3,2,1];
const test3 = [1,1,2,2];
shuffle(test1, 3)
shuffle(test2, 4)
shuffle(test3, 2)
