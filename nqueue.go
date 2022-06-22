package main

func TotalNQueens(n int) int {
	// write your code here
	count := 0
	return dfs(n, 0, 0, 0, 0, &count)
}

func dfs(n, row, col, pie, na int, count *int) int {
	if row >= n {
		*count += 1
		return *count
	}
	bits := (^(col | pie | na)) & ((1 << n) - 1)
	for bits > 0 {
		p := bits & -bits
		dfs(n, row+1, col|p, (pie|p)<<1, (na|p)>>1, count)
		bits &= bits - 1
	}
	return *count
}
