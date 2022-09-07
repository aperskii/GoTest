package main

import "fmt"

func main() {

	xoBoard := [3][3]string{}
	player := "X"

	for {

		fmt.Println("Player : ", player)

		var row int
		fmt.Println("please enter row number, 0,1 or 2")

		fmt.Scanln(&row)
		if row > 2 {
			fmt.Println("Invalid input")
			continue
		}

		var column int
		fmt.Println("please enter column number, 0,1 or 2")
		fmt.Scanln(&column)
		if column > 2 {
			fmt.Println("Invalid input")
			continue
		}
		fmt.Println("row:", row, "column:", column)

		if xoBoard[row][column] == "" {
			xoBoard[row][column] = player
		} else {
			fmt.Println("invalid input")
			continue
		}
		fmt.Println(xoBoard[0])
		fmt.Println(xoBoard[1])
		fmt.Println(xoBoard[2])

		for i := 0; i <= 2; i++ {
			if (xoBoard[i][0] == player && xoBoard[i][0] == xoBoard[i][1] && xoBoard[i][0] == xoBoard[i][2]) ||
				(xoBoard[0][i] == player && xoBoard[0][i] == xoBoard[1][i] && xoBoard[0][i] == xoBoard[2][i]) {
				fmt.Println("Winner")
				return
			}
		}
		if (xoBoard[0][0] == player && xoBoard[0][0] == xoBoard[1][1] && xoBoard[0][0] == xoBoard[2][2]) || (xoBoard[0][2] == player && xoBoard[0][2] == xoBoard[1][1] && xoBoard[0][2] == xoBoard[2][0]) {
			fmt.Println("Winner")
			return
		}

		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}
}
