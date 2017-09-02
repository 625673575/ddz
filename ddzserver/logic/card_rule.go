package logic

import (
	"ddzserver/message"
)

func CheckIsSingle(cards []message.CardInfo) bool {
	return len(cards) == 1
}
func CheckIsDouble(cards []message.CardInfo) bool {
	return len(cards) == 2 && cards[0].Weight == cards[1].Weight
}
func CheckIsStraight(cards []message.CardInfo) bool {
	l := len(cards)
	if l < 5 || l > 12 {
		return false
	}

	for i := 0; i < l-1; i++ {
		w := cards[i].Weight
		if cards[i+1].Weight-w != 1 {
			return false
		}

		//不能超过A
		if w > message.Weight_One || cards[i+1].Weight > message.Weight_One {
			return false
		}
	}
	return true
}
func CheckIsDoubleStraight(cards []message.CardInfo) bool {
	l := len(cards)
	if l < 6 || l%2 != 0 {
		return false
	}
	for i := 0; i < l; i += 2 {
		if cards[i+1].Weight != cards[i].Weight {
			return false
		}
		if i < l-2 {
			if (cards[i+2].Weight-cards[i].Weight != 1) {
				return false
			}
			//不能超过A
			if cards[i].Weight > message.Weight_One || cards[i+2].Weight > message.Weight_One {
				return false
			}
		}
	}

	return true
}
func CheckIsTripleStraight(cards []message.CardInfo) bool {
	l := len(cards)
	if l < 6 || l%3 != 0 {
		return false
	}

	for i := 0; i < l; i += 3 {
		if cards[i+1].Weight != cards[i].Weight {
			return false
		}

		if cards[i+2].Weight != cards[i].Weight {
			return false
		}

		if cards[i+1].Weight != cards[i+2].Weight {
			return false
		}

		if i < l-3 && cards[i+3].Weight-cards[i].Weight != 1 {
			return false
		}

		if cards[i].Weight > message.Weight_One || cards[i+3].Weight > message.Weight_One {
			return false
		}

	}
	return true
}

func CheckIsOnlyThree(cards []message.CardInfo) bool {
	l := len(cards)
	if l%3 != 0 {
		return false
	}

	if cards[0].Weight != cards[1].Weight {
		return false
	}
	if cards[1].Weight != cards[2].Weight {
		return false
	}
	if cards[0].Weight != cards[2].Weight {
		return false
	}

	return true
}
func CheckIsThreeAndOne(cards []message.CardInfo) bool {
	l := len(cards)
	if l != 4 {
		return false
	}

	if cards[0].Weight == cards[1].Weight && cards[1].Weight == cards[2].Weight {
		return true
	} else if cards[1].Weight == cards[2].Weight && cards[2].Weight == cards[3].Weight {
		return true
	}
	return false
}
func CheckIsThreeAndTwo(cards []message.CardInfo) bool {
	l := len(cards)
	if l != 5 {
		return false
	}

	if cards[0].Weight == cards[1].Weight && cards[1].Weight == cards[2].Weight {
		if cards[3].Weight == cards[4].Weight {
			return true
		}
	} else if cards[2].Weight == cards[3].Weight && cards[3].Weight == cards[4].Weight {
		if cards[0].Weight == cards[1].Weight {
			return true
		}
	}
	return false
}
func CheckIsBoom(cards []message.CardInfo) bool {
	l := len(cards)
	if l != 4 {
		return false
	}

	if cards[0].Weight != cards[1].Weight {
		return false
	}
	if cards[1].Weight != cards[2].Weight {
		return false
	}
	if cards[2].Weight != cards[3].Weight {
		return false
	}

	return true
}
func CheckIsJokerBoom(cards []message.CardInfo) bool {
	l := len(cards)
	if l != 2 {
		return false
	}
	if cards[0].Weight == message.Weight_SJoker {
		if cards[1].Weight == message.Weight_LJoker {
			return true
		}
		return false
	} else if cards[0].Weight == message.Weight_LJoker {
		if cards[1].Weight == message.Weight_SJoker {
			return true
		}
		return false
	}

	return false
}
func PopEnable(cards []message.CardInfo) (bool, message.CardsType) {
	cardType := message.CardsType_UnknowCard
	isRule := false
	switch len(cards) {
	case 1:
		isRule = true
		cardType = message.CardsType_Single

	case 2:
		if CheckIsDouble(cards) {
			isRule = true
			cardType = message.CardsType_Double
		} else if CheckIsJokerBoom(cards) {
			isRule = true
			cardType = message.CardsType_JokerBoom
		}
	case 3:
		if CheckIsOnlyThree(cards) {
			isRule = true
			cardType = message.CardsType_OnlyThree
		}
	case 4:
		if CheckIsBoom(cards) {
			isRule = true
			cardType = message.CardsType_Boom
		} else if CheckIsThreeAndOne(cards) {
			isRule = true
			cardType = message.CardsType_ThreeAndOne
		}
	case 5:
		if CheckIsStraight(cards) {
			isRule = true
			cardType = message.CardsType_Straight
		} else if CheckIsThreeAndTwo(cards) {
			isRule = true
			cardType = message.CardsType_ThreeAndTwo
		}
	case 6:
		if CheckIsStraight(cards) {
			isRule = true
			cardType = message.CardsType_Straight
		} else if CheckIsTripleStraight(cards) {
			isRule = true;
			cardType = message.CardsType_TripleStraight
		} else if CheckIsDoubleStraight(cards) {
			isRule = true
			cardType = message.CardsType_DoubleStraight
		}
	case 7:
		if CheckIsStraight(cards) {
			isRule = true
			cardType = message.CardsType_Straight
		}
	case 8:
		if CheckIsStraight(cards) {
			isRule = true
			cardType = message.CardsType_Straight
		} else if CheckIsDoubleStraight(cards) {
			isRule = true
			cardType = message.CardsType_DoubleStraight
		}
	case 9:
		if CheckIsStraight(cards) {
			isRule = true
			cardType = message.CardsType_Straight
		} else if (CheckIsOnlyThree(cards)) {
			isRule = true
			cardType = message.CardsType_OnlyThree
		}

	case 10:
		if CheckIsStraight(cards) {
			isRule = true
			cardType = message.CardsType_Straight
		} else if CheckIsDoubleStraight(cards) {
			isRule = true
			cardType = message.CardsType_DoubleStraight
		}

	case 11:
		if CheckIsStraight(cards) {
			isRule = true
			cardType = message.CardsType_Straight
		}

	case 12:
		if CheckIsStraight(cards) {
			isRule = true
			cardType = message.CardsType_Straight
		} else if CheckIsDoubleStraight(cards) {
			isRule = true
			cardType = message.CardsType_DoubleStraight
		} else if CheckIsOnlyThree(cards) {
			isRule = true
			cardType = message.CardsType_OnlyThree
		}
	case 13:
	case 14:
		if CheckIsDoubleStraight(cards) {
			isRule = true
			cardType = message.CardsType_DoubleStraight
		}

	case 15:
		if CheckIsOnlyThree(cards) {
			isRule = true
			cardType = message.CardsType_OnlyThree
		}
	case 16:
		if CheckIsDoubleStraight(cards) {
			isRule = true
			cardType = message.CardsType_DoubleStraight
		}
	case 17:
	case 18:
		if CheckIsDoubleStraight(cards) {
			isRule = true
			cardType = message.CardsType_DoubleStraight
		} else if CheckIsOnlyThree(cards) {
			isRule = true
			cardType = message.CardsType_OnlyThree
		}
	case 19:
	case 20:
		if CheckIsDoubleStraight(cards) {
			isRule = true
			cardType = message.CardsType_DoubleStraight
		}
	default:
	}
	return isRule, cardType
}
