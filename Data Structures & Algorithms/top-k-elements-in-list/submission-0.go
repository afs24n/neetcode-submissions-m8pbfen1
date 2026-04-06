func topKFrequent(nums []int, k int) []int {

    frequencyMap := make(map[int]int)


    for _, value := range nums {
        frequencyMap[value]++
    }

    bucket := make([][]int,len(nums)+1)

    for num, frequency := range frequencyMap{
        bucket[frequency] = append(bucket[frequency],num)
    }

    res := make([]int,0,k)
    for i := len(bucket)-1; i >= 0; i-- {
       if len(bucket[i]) > 0 {
            for _, num := range bucket[i]{
                res = append(res, num)

                if len(res) == k {
                    return res
                }
            }
       }

    }
    return res
}
