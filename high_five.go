package main

import (
    "fmt"
	"math"
    "sort"
)

func averageScore(gradeMap map[int][]int, id int) int {
	sum := 0.0
	grades, ok := gradeMap[id]
	if !ok {
		return -1
	}
    sort.Sort(sort.Reverse(sort.IntSlice(grades)))
    grades = grades[:5]
	for _, grade := range grades {
		sum += float64(grade)
    }
	gradeAmt := float64(len(grades))
    result := int(math.Floor(sum / gradeAmt))
    fmt.Printf("%d: sum=%f, %+v/%f=%d\n", id, sum, grades, gradeAmt, result)
    return result
}

func containsId(ids []int, id int) bool {
  for _, x := range ids {
    if x == id {
      return true
    }
  }
  return false
}

func highFive(items [][]int) [][]int {
  ids := make([]int, 0)
	gradeMap := make(map[int][]int)
	for _, score := range items {
		id := score[0]
		grade := score[1]
        if !containsId(ids, id) {
          ids = append(ids, id)
        }
		if _, ok := gradeMap[id]; !ok {
			gradeMap[id] = make([]int, 0)
		}
		gradeMap[id] = append(gradeMap[id], grade)
	}
	// reconstruction
	result := make([][]int, 0)
    for _, id := range ids {
		average := averageScore(gradeMap, id)
		result = append(result, []int{id, average})
	}
	return result
}

