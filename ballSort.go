package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	color "github.com/fatih/color"
)

func main() {
	// err := termbox.Init()

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// tbprint(1, 1, termbox.ColorRed, termbox.ColorDefault, "Hello terminal! \n Blah")
	// termbox.Flush()

	// time.Sleep(time.Second)
	// termbox.Close()

	//termbox.Init()
	//var bTest = []int{1, 2, 3, 4, 5}
	var b1 = []int{5, 3, 2, 4}
	var b2 = []int{2, 2, 3, 1}
	var b3 = []int{1, 4, 3, 5}
	var b4 = []int{1, 5, 3, 5}
	var b5 = []int{4, 2, 1, 4}
	var b6 = []int{0, 0, 0, 0}
	var b7 = []int{0, 0, 0, 0}
	var flasks = [][]int{b1, b2, b3, b4, b5, b6, b7}
	//printFlask(b1)
	for {
		printFlasks(flasks)
		if isWin(flasks) {
			fmt.Println("You won!!!")
			break
		}
		var what = getInput("What? > ") - 1
		var where = getInput("Where? > ") - 1
		// TODO check input
		// fmt.Printf("%d %d", what, where)
		var isMoved = moveBall(what, where, flasks)
		if !isMoved {
			fmt.Println("Cant move")
		}
		if isWin(flasks) {
			fmt.Println("You won!!!")
			break
		}
	}
}

// This function is often useful:
// func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
// 	for _, c := range msg {
// 		termbox.SetCell(x, y, c, fg, bg)
// 		x += runewidth.RuneWidth(c)
// 	}
// }

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
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

// func isNoMoves() bool {

// }

func getInput(question string) int {
	for {
		fmt.Print(question)
		var inputString string
		_, err := fmt.Scan(&inputString)
		input, err := strconv.Atoi(inputString)
		if err != nil || !isCorrectInput(input) {
			fmt.Println("Wrong input. Try again.")
			continue
		} else {
			return input
		}
	}
}

func isCorrectInput(input int) bool {
	if input > 0 && input < 8 {
		return true
	}
	return false
}

func printFlasks(fs [][]int) {
	var totalHeight = len(fs[0]) + 3
	var flaskCount = len(fs)
	for i := 0; i < totalHeight; i++ {
		switch i {
		case 0:
			printFlaskNumbers(flaskCount)
		case 1:
			printFlaskTops(flaskCount)
		case (totalHeight - 1):
			printFlaskBottoms(flaskCount)
		default:
			printFlaskRow(fs, i-2)
		}
		fmt.Println("")
	}
}

func printFlaskRow(flasks [][]int, currentRow int) {
	for i := 0; i < len(flasks); i++ {
		glassColor := color.New(color.FgCyan).SprintFunc()
		glass := glassColor("¦")
		ball := getColoredBall(flasks[i][currentRow])
		fmt.Printf("%s%s%s%s%s", space(2), glass, ball, glass, space(5))
	}
}

func printFlaskBottoms(flaskCount int) {
	for i := 0; i < flaskCount; i++ {
		glassColor := color.New(color.FgCyan).SprintFunc()
		glass := glassColor("¯")
		fmt.Printf("%s%s%s", space(3), glass, space(6))
	}
}

func printFlaskNumbers(flaskCount int) {
	for i := 0; i < flaskCount; i++ {
		numberColor := color.New(color.FgWhite).SprintFunc()
		number := numberColor(i + 1)
		fmt.Printf("%s%s%s", space(3), number, space(6))
	}
}

func space(n int) string {
	return strings.Repeat(" ", n)
}

func printFlaskTops(flaskCount int) {
	for i := 0; i < flaskCount; i++ {
		glassColor := color.New(color.FgCyan).SprintFunc()
		glass := glassColor("_")
		fmt.Printf("%s%s%s%s%s", space(2), glass, space(1), glass, space(5))
	}
}

func printFlaskOld(b []int) {
	color.Cyan("_ _")
	for _, c := range b {
		glassColor := color.New(color.FgCyan).SprintFunc()
		glass := glassColor("¦")
		ball := getColoredBall(c)
		fmt.Printf("%s%s%s  %d\n", glass, ball, glass, c)
	}
	color.Cyan(" ¯")
}

func getColoredBall(num int) string {
	switch num {
	case 0:
		return " "
	case 1:
		ballColor := color.New(color.FgRed).SprintFunc()
		return ballColor("•")
	case 2:
		ballColor := color.New(color.FgYellow).SprintFunc()
		return ballColor("•")
	case 3:
		ballColor := color.New(color.FgGreen).SprintFunc()
		return ballColor("•")
	case 4:
		ballColor := color.New(color.FgBlue).SprintFunc()
		return ballColor("•")
	case 5:
		ballColor := color.New(color.FgMagenta).SprintFunc()
		return ballColor("•")
	default:
		fmt.Println("not supported color")
		return ""
	}
}
