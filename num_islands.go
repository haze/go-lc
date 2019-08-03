package main

import "fmt"

func printIsland(islands [][]byte) {
	for x := 0; x < len(islands); x += 1 {
		for y := 0; y < len(islands[x]); y += 1 {
			if y == len(islands[x])-1 {
				fmt.Printf("%b", islands[x][y])
			} else {
				fmt.Printf("%b, ", islands[x][y])
			}
		}
		fmt.Println()
	}
}

func numIslands(islands [][]byte) int {
	totalIslands := 0
	for x := 0; x < len(islands); x += 1 {
		for y := 0; y < len(islands[x]); y += 1 {
			node := islands[x][y]
			fmt.Printf("note: %b\n", node)
			printIsland(islands)
			if node == 1 {
				dfs(&islands, x, y)
				totalIslands += 1
			}
		}
	}
	return totalIslands
}

func dfs(src *[][]byte, nx, ny int) {
	if nx >= len(*src) || nx < 0 {
		return
	}
	if ny >= len((*src)[nx]) || ny < 0 {
		return
	}
	if (*src)[nx][ny] == 0 {
		return
	}
	(*src)[nx][ny] = 0
	dfs(src, nx+1, ny)
	dfs(src, nx-1, ny)
	dfs(src, nx, ny+1)
	dfs(src, nx, ny-1)
}

func main() {
	fmt.Println(numIslands([][]byte{
		[]byte{1, 1, 1, 1, 0},
		[]byte{1, 1, 0, 1, 0},
		[]byte{1, 1, 0, 0, 0},
		[]byte{0, 0, 0, 0, 1},
	}))
}
