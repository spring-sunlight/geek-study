package example_hash_set_map

func robotSim(commands []int, obstacles [][]int) int {
	ans := 0
	obstaclesMap := make(map[[2]int]struct{})
	for i := 0; i < len(obstacles); i++ {
		obstaclesMap[[2]int{obstacles[i][0], obstacles[i][1]}] = struct{}{}
	}
	posX, posY := 0, 0
	dir := 0
	dirX := []int{0, 1, 0, -1}
	dirY := []int{1, 0, -1, 0}
	for i := 0; i < len(commands); i++ {
		if commands[i] == -1 {
			dir = (dir + 1) % 4
			continue
		} else if commands[i] == -2 {
			dir = (dir + 3) % 4
			continue
		} else {
			for j := 0; j < commands[i]; j++ {
				if _, ok := obstaclesMap[[2]int{posX + dirX[dir], posY + dirY[dir]}]; ok {
					break
				} else {
					posX = posX + dirX[dir]
					posY = posY + dirY[dir]
				}
			}
			ans = max(ans, posX*posX+posY*posY)
		}

	}
	return ans
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
