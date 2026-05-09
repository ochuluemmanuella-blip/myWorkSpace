package main

import (
	"fmt"
	"os"
	"strings"
)

func readMaze(filename string) ([][]rune, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	var maze [][]rune
	for _, line := range lines {
		maze = append(maze, []rune(line))
	}
	return maze, nil
}
func findStart(maze [][]rune) (int, int, error) {
	for row := range maze {
		for col := range maze[row] {
			if maze[row][col] == 'S' {
				return row, col, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("no start position 'S' found in maze")
}
