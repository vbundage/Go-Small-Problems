// https://www.codewars.com/kata/562f91ff6a8b77dfe900006e/go

package kata

import "math"

func Movie(card, ticket int, perc float64) int {
	var systemA int
	var count int
	prevTicketPrice := float64(ticket) * perc
	systemB := float64(card) + prevTicketPrice
	prevTicketPrice = float64(prevTicketPrice) * perc

	for int(math.Ceil(systemB)) >= systemA {
		if count != 0 {
			systemB += prevTicketPrice
			prevTicketPrice = float64(prevTicketPrice) * perc
		}
		systemA += ticket
		count++
	}
	return count
}
