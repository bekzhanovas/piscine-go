package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	counter := 0
	for i := 1; i <= nb; i++ {
		if nb%i == 0 {
			counter++
			if counter > 2 {
				return false
			}
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	} else {
		for i := 1; i <= nb; i++ {
			if IsPrime(nb) == false {
				nb++
			} else {
				break
			}
		}
	}
	return nb
}
