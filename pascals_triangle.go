package pascal

var triangle [][]int

func Triangle(n int) [][]int {
	triangle = make([][]int, n)
	for i := 0; i < n; i++ {
		triangle[i] = Row(i)
	}
	return triangle
}

func Row(n int) []int {
	lenRow := n + 1
	row := make([]int, lenRow)
	row[0] = 1
	row[lenRow-1] = 1
	var rowBefore []int
	if n >= 1 {
		rowBefore = triangle[n-1]
	}
	for i := 1; i < lenRow-1; i++ {
		row[i] = rowBefore[i-1] + rowBefore[i]
	}
	return row
}
