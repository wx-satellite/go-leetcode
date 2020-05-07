package main

import "fmt"

// 二维数组的行坐标规律、列坐标规律其实很容易发现，"3*3 小正方形"的坐标规律还是比较难发现的，需要好好理解。
func isValidSudoku(board [][]byte) bool {

	// 数字 1-9 在每一行只能出现一次。
	for i := 0; i < 9; i++ {
		isExist := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			// 字符 '.' 的数值表示为 46
			if 46 != board[i][j] {
				v := isExist[board[i][j]]
				if v {
					return false
				}
				isExist[board[i][j]] = true
			}

		}
	}

	// 数字 1-9 在每一列只能出现一次。
	for i := 0; i < 9; i++ {
		isExist := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if 46 != board[j][i] {
				v := isExist[board[j][i]]
				if v {
					return false
				}
				isExist[board[j][i]] = true
			}

		}
	}


	// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
	for n := 0; n < 9; n++ {
		isExist := make(map[byte]bool)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				value := board[i + (n / 3) * 3][j + n % 3 * 3]
				if 46 != value {
					v := isExist[value]
					if v {
						return false
					}
					isExist[value] = true
				}
			}
		}
	}

	return true
}

func main() {
	broad := [][]byte{
	{'5','3','.','.','7','.','.','.','.'},
	{'6','.','.','1','9','5','.','.','.'},
	{'.','9','8','.','.','.','.','6','.'},
	{'8','.','.','.','6','.','.','.','3'},
	{'4','.','.','8','.','3','.','.','1'},
	{'7','.','.','.','2','.','.','.','6'},
	{'.','6','.','.','.','.','2','8','.'},
	{'.','.','.','4','1','9','.','.','5'},
	{'.','.','.','.','8','.','.','7','9'},
	}

	broad = [][]byte{
		{'.','.','4','.','.','.','6','3','.'},
		{'.','.','.','.','.','.','.','.','.'},
		{'5','.','.','.','.','.','.','9','.'},
		{'.','.','.','5','6','.','.','.','.'},
		{'4','.','3','.','.','.','.','.','1'},
		{'.','.','.','7','.','.','.','.','.'},
		{'.','.','.','5','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','.','.'},
	}
	fmt.Println(isValidSudoku(broad))


}
