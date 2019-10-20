package piscine

func FindNextPrime(nb int) int {
	counter := 0
	for i := 1; i <= nb; i++ {
		if nb%i == 0 {
			counter++
			if counter > 2 {
				for j := nb+1; j < 1000000; j++ {
					counter_next := 0
					for k := 1; k <= j; k++ {
						if j%k == 0 {
							counter_next++
							if counter_next > 2 {
								continue
							}
						}
					}
					return j
				}
			}
		}
	}
	return nb
}
