// https://www.codewars.com/kata/514b92a657cdc65150000006/go

package kata

func Multiple3And5(number int) (sum int) {
	for count := 3; count < number; count++ {
		if count%3 == 0 || count%5 == 0 {
			sum += count
		}
	}
	return
}
