package Num_278_FirstBadVersion

import "sort"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var num int

func isBadVersion(i int) bool {
	return i == num
}

func firstBadVersion(n int) int {
	return sort.Search(n, func(i int) bool {
		return isBadVersion(i)
	})
}
