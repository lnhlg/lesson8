package lesson8

//numIslands: 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//此外，你可以假设该网格的四条边均被水包围。
/*parameter
grid: 取值为0和1的byte型二维数组
return: 岛屿的数量
*/
func NumIslands(grid [][]byte) int {

	rows, cols := len(grid), len(grid[0])

	getFather := func() ([]int, int) {

		count := 0
		father := make([]int, cols*rows)

		for i := 0; i < rows; i++ {

			for j := 0; j < cols; j++ {

				if grid[i][j] == '1' {

					count++
					father[i*cols+j] = i*cols + j
				}
			}
		}

		return father, count
	}

	father, count := getFather()

	var find func(int) int
	find = func(item int) int {

		if father[item] != item {

			father[item] = find(father[item])
		}

		return father[item]
	}

	union := func(i, j, a, b int) {

		c1 := i*cols + j
		c2 := a*cols + b

		fc1 := find(c1)
		fc2 := find(c2)

		if fc1 != fc2 {

			father[fc1] = fc2
			count--
		}
	}

	for i := 0; i < rows; i++ {

		for j := 0; j < cols; j++ {

			if grid[i][j] == '1' {

				d := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
				for n := range d {

					x := i + d[n][0]
					y := j + d[n][1]

					if x >= 0 && x < rows && y >= 0 && y < cols && grid[x][y] == '1' {

						union(i, j, x, y)
					}
				}
			}
		}
	}

	return count
}
