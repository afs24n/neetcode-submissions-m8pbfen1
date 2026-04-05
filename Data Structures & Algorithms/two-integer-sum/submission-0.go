func twoSum(nums []int, target int) []int {
    
    burnt := make(map[int]int)
    
    for i := 0; i < len(nums); i++ {
    
        difference := target - nums[i]

        if pos, exists := burnt[nums[i]]; exists {
            return []int{pos,i}
        }

        burnt[difference] = i
        
    }
    return nil
}
