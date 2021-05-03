package main

import (
	"fmt"
	"strconv"
)

func main() {
	flasks := level1()
	for {
		printFlasks(flasks)
		if isWin(flasks) {
			fmt.Println("You won!!!")
			break
		}
		var what = getInput("What? > ", len(flasks)) - 1
		var where = getInput("Where? > ", len(flasks)) - 1
		var isMoved = moveBall(what, where, flasks)
		if !isMoved {
			fmt.Println("Can't move")
		}
		if isWin(flasks) {
			fmt.Println("You won!!!")
			break
		}
	}
}

func canPut(ball int, where int, flasks [][]int) bool {
	for i := 0; i < len(flasks[where]); i++ {
		// flask is full
		if i == 0 && flasks[where][i] != 0 {
			return false
		}
		// flask is empty
		if i == len(flasks[where])-1 && flasks[where][i] == 0 {
			return true
		}
		// can put ball
		if flasks[where][i] == 0 && flasks[where][i+1] == ball {
			return true
		}
		// can`t put ball
		if flasks[where][i] == 0 && flasks[where][i+1] != ball && flasks[where][i+1] != 0 {
			return false
		}
	}
	return false
}

func moveBall(what int, where int, flasks [][]int) bool {
	// take ball
	var wasMoved = false
	var ballToTake = getBall(what, flasks)
	if !canPut(ballToTake, where, flasks) {
		return wasMoved
	}
	if ballToTake == 0 {
		return wasMoved
	}
	var takenBall = takeBall(what, flasks)
	// put ball
	for i := 0; i < len(flasks[where]); i++ {
		// flask is full
		if i == 0 && flasks[where][i] != 0 {
			break
		}
		// flask is empty
		if i == len(flasks[where])-1 && flasks[where][i] == 0 {
			flasks[where][i] = takenBall
			wasMoved = true
			break
		}
		// can put ball
		if flasks[where][i] == 0 && flasks[where][i+1] == takenBall {
			flasks[where][i] = takenBall
			wasMoved = true
			break
		}
		// can`t put ball
		if flasks[where][i] == 0 && flasks[where][i+1] != takenBall && flasks[where][i+1] != 0 {
			break
		}
	}
	return wasMoved
}

func takeBall(what int, flasks [][]int) int {
	var ball = 0
	for i := 0; i < len(flasks[what]); i++ {
		if flasks[what][i] != 0 {
			ball = flasks[what][i]
			flasks[what][i] = 0
			break
		}
	}
	return ball
}

func getBall(what int, flasks [][]int) int {
	var ball = 0
	for i := 0; i < len(flasks[what]); i++ {
		if flasks[what][i] != 0 {
			ball = flasks[what][i]
			break
		}
	}
	return ball
}

func isWin(flasks [][]int) bool {
	var isWin = true
	for i := 0; i < len(flasks); i++ {
		var firstBall = flasks[i][0]
		for j := 1; j < len(flasks[i]); j++ {
			if flasks[i][j] != firstBall {
				isWin = false
			}
		}
	}
	return isWin
}

func getInput(question string, flaskLen int) int {
	for {
		fmt.Print(question)
		var inputString string
		_, err := fmt.Scan(&inputString)
		input, err := strconv.Atoi(inputString)
		if err != nil || !isCorrectInput(input, flaskLen) {
			fmt.Println("Wrong input. Try again.")
			continue
		} else {
			return input
		}
	}
}

func isCorrectInput(input int, flaskLen int) bool {
	if input > 0 && input < flaskLen {
		return true
	}
	return false
}



