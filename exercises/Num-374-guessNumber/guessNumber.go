package Num_374_GuessNumber

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

var guessNum int

func guess(num int) int {
	if num == guessNum {
		return 0
	} else if num > guessNum {
		return -1
	} else {
		return 1
	}
}

func guessNumber(n int) int {
	m, target := 0, n/2
	for ; m <= n; target = (m + n) / 2 {
		t := guess(target)
		if t == 0 {
			return target
		} else if t < 0 {
			n = target - 1
		} else {
			m = target + 1
		}
	}
	return target
}
