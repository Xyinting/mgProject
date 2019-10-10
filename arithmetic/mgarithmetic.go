package main

import (
	"fmt"
	"os"
)

func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic("file not exist")
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze


}

type point struct {
	i, j int
}

func walk(maze [][]int, start, end point){




}

func main() {

	maze := readMaze("arithmetic/maze.in")

	for _, row := range maze {
		for _, val := range row{
			fmt.Printf("%d", val)
		}
		fmt.Println()
	}

	walk(maze, point{0, 0}, point{len(maze)-1 , len(maze[0]) - 1})
}
