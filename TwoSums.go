func twoSum(nums []int, target int) []int {
    var res []int
    for i:=0; i<len(nums);i++ {
        for j:= i+1; j< len(nums);j++ {
            if nums[i] + nums[j] == target {
                res= append(res,nums[i])
                res= append(res,nums[j])
            }
        }
    }
    return res
}
