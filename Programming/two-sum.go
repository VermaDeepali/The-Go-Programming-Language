// Two Sum - Pair with Given Sum
// Difficulty: EasyAccuracy: 30.61%Submissions: 463K+Points: 2Average Time: 20m
// Given an array arr[] of integers and another integer target. Determine if there exist two distinct indices such that the sum of their elements is equal to the target.

// Examples:

// Input: arr[] = [0, -1, 2, -3, 1], target = -2
// Output: true
// Explanation: arr[3] + arr[4] = -3 + 1 = -2
// Input: arr[] = [1, -2, 1, 0, 5], target = 0
// Output: false
// Explanation: None of the pair makes a sum of 0
// Input: arr[] = [11], target = 11
// Output: false
// Explanation: No pair is possible as only one element is present in arr[]
// Constraints:
// 1 ≤ arr.size ≤ 105
// -105 ≤ arr[i] ≤ 105
// -2*105 ≤ target ≤ 2*105



func twoSum(nums []int, target int) []int {
    // sum := 0
    // for i:=0; i<len(nums); i++ {
    //     for j:=i+1; j<len(nums); j++ {
    //         sum = nums[i]+ nums[j]
    //         if(sum == target){
    //             return []int{i,j}
    //         }
    //     }
    // }
    // return []int{}
    m := make(map[int]int)

    for i, num := range nums {

        diff := target - num

        if idx, found := m[diff]; found {
            return []int{idx, i}
        }
        m[num]=i
    }
    return []int{}
}
