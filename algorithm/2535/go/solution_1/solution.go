package solution

func differenceOfSum(nums []int) int {
    var elementSum, digitSum int
    
    for _, number := range nums {
        elementSum += number
        
        for number > 0 {
            digitSum += number % 10
            number /= 10
        }
    }
    
    return elementSum - digitSum
}
