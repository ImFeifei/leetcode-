package main
import "fmt"

func maxCoins(nums []int) int {
    if len(nums) == 0 { return 0 }
    if len(nums) == 1 { return nums[0] }

    // dp初始化
    // dp[i][j]表示戳爆(i,j)内所有气球可以得到达最大值 注意：不戳i,j
    // dp实际上只用到了右上三角，因为 j > i
    // 首位各加一个值为1的虚拟气球，这样最终的输出就是 dp[0][length-1]
    length := len(nums) + 2
    nums = append([]int{1}, nums...)
    nums = append(nums, 1)
    dp := make([][]int, length)
    for i := 0; i < length; i++ {
        dp[i] = make([]int, length)
    }

    for i := length - 2; i >= 0; i-- {
        for j := i + 1; j < length; j++ {
            if j == i + 1 { 
                // i,j之间没有其他气球，没有戳破的选择，故得到0个硬币
                dp[i][j] = 0 
            } else if j == i + 2 {
                // i,j之间仅有一个气球，是唯一的选择，戳破中间气球得到三数乘积个硬币,报错是因为dp和num下标不一样
                //因此 在上面重新修改了num数组增加虚拟气球使得和dp中一样
                dp[i][j] = nums[i]*nums[i+1]*nums[j]
            } else {
                // i,j之间有多个气球，需要依赖子问题选取最大值
                for k := i + 1; k < j; k++{
                    dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[j]*nums[k])
                }
            }
        }
    }
    return dp[0][length-1]
}
func max(x, y int)int {
    if x > y { return x }
    return y
}
func main() {
    nums := []int{3, 1, 5, 8}
    fmt.Println("input: ", nums)
    fmt.Println("output: ", maxCoins(nums))
}