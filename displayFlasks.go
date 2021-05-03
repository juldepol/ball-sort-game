package main

import (
	"fmt"
	"strings"

	color "github.com/fatih/color"
)

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

func printFlaskNumbers(flaskCount int) {
	for i := 0; i < flaskCount; i++ {
		numberColor := color.New(color.FgWhite).SprintFunc()
		number := numberColor(i + 1)
		fmt.Printf("%s%s%s", space(3), number, space(6))
	}
}

func printFlaskTops(flaskCount int) {
	for i := 0; i < flaskCount; i++ {
		glassColor := color.New(color.FgCyan).SprintFunc()
		glass := glassColor("_")
		fmt.Printf("%s%s%s%s%s", space(2), glass, space(1), glass, space(5))
	}
}

func printFlaskBottoms(flaskCount int) {
	for i := 0; i < flaskCount; i++ {
		glassColor := color.New(color.FgCyan).SprintFunc()
		glass := glassColor("¯")
		fmt.Printf("%s%s%s", space(3), glass, space(6))
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

func space(n int) string {
	return strings.Repeat(" ", n)
}