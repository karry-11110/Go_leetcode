// 1. 确定dp数组下标含义 dp[i][j] 到每一个坐标可能的路径种类
// 2. 递推公式
// if (obstacleGrid[i][j] == 0) { // 当(i, j)没有障碍的时候，再推导dp[i][j]
// dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 3. 初始化 dp[i][0]=1 dp[0][i]=1 初始化横竖就可初始化的部分，很容易忽略了障碍之后应该都是0的情况
// 4. 遍历顺序 一行一行遍历
// 5. 推导结果
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	//1.定义数组
	dp := make([][]int, len(obstacleGrid))
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	//初始化
	for i := 0; i < len(obstacleGrid); i++ {
		//如果是障碍物, 后面的就都是0, 不用循环了
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 0; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	//推导
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			// 如果obstacleGrid[i][j]这个点是障碍物, 那么我们的dp[i][j]保持为0
			if obstacleGrid[i][j] != 1 {
				// 否则我们需要计算当前点可以到达的路径数
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}