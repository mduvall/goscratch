package algorithms

// Number of ways to do n choose m
func BinomialCoefficients(n int, m int) int {
	tab := [20][20]int{}

	for i := 0; i < 20; i++ {
		tab[i][0] = 1
		tab[i][i] = 1
	}

	for i := 1; i < 20; i++ {
		for j := 1; j < i; j++ {
			tab[i][j] = tab[i-1][j-1] + tab[i-1][j]
		}
	}
	return tab[n][m]
}
