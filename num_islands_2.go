package main

func makeBlankIsland(m, n int) [][]byte {
  islands := make([][]byte, m)
  for mm := 0; m > mm; mm += 1 {
    islands[mm] = make([]byte, n)
    for nn := 0; n > nn; nn += 1 {
      islands[mm][nn] = '0'
    }
  }
  return islands
}

// optimization, save previous num island results to
// prevent deduplicating & doing dfs again
func numIslands2(m, n int, positions [][]int) []int {
  table := make([]int, len(positions))
  islands := makeBlankIsland(m, n)
  for idx, pos := range positions {
    islands[pos[0]][pos[1]] = '1'
    duplicate := make([][]byte, len(islands))
    for i := range islands {
      duplicate[i] = make([]byte, len(islands[i]))
      copy(duplicate[i], islands[i])
    }
    result := numIslands(duplicate)
    table[idx] = result
  }
  return table
}

func numIslands(islands [][]byte) int {
  totalIslands := 0
  for x := 0; x < len(islands); x += 1 {
    for y := 0; y < len(islands[x]); y += 1 {
      if islands[x][y] == '1' {
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
  if (*src)[nx][ny] == '0' {
    return
  }
  (*src)[nx][ny] = '0'
  dfs(src, nx+1, ny)
  dfs(src, nx-1, ny)
  dfs(src, nx, ny+1)
  dfs(src, nx, ny-1)
}
