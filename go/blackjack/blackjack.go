package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king":
		return 10
	case "queen":
		return 10
	case "jack":
		return 10
	case "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	if card1Value == 11 && card2Value == 11 {
		return "P"
	}

	cardsValue := card1Value + card2Value

	if (cardsValue >= 21) &&
		dealerCardValue < 10 {
		return "W"
	}

	if cardsValue >= 17 ||
		(cardsValue >= 12 && cardsValue <= 16 && dealerCardValue < 7) {
		return "S"
	}

	return "H"
}
