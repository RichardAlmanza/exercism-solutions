package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var cardValue int

	switch card {
	case "ace":
		cardValue = 11
	case "two":
		cardValue = 2
	case "three":
		cardValue = 3
	case "four":
		cardValue = 4
	case "five":
		cardValue = 5
	case "six":
		cardValue = 6
	case "seven":
		cardValue = 7
	case "eight":
		cardValue = 8
	case "nine":
		cardValue = 9
	case "ten":
		cardValue = 10
	case "jack":
		cardValue = 10
	case "queen":
		cardValue = 10
	case "king":
		cardValue = 10
	default:
		cardValue = 0
	}

	return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var decision string

	ownCardsValue := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	switch {
	case card1 == card2 && card1 == "ace":
		decision = "P"
	case ownCardsValue == 21 && dealerCardValue < 10:
		decision = "W"
	case ownCardsValue == 21:
		decision = "S"
	case ownCardsValue >= 17 && ownCardsValue <= 20:
		decision = "S"
	case ownCardsValue >= 12 && ownCardsValue <= 16 && dealerCardValue >= 7:
		decision = "H"
	case ownCardsValue >= 12 && ownCardsValue <= 16:
		decision = "S"
	case ownCardsValue <= 11:
		decision = "H"
	}

	return decision
}
