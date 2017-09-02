package logic

import (
	"ddzserver/message"
	"reflect"
)

type CardRule struct {
	CardCount             int32
	PlayerCount           int32
	EveryFarmerCardsCount int32
}
type PlayRule struct {
	CardRule    CardRule
	playerIndex int32
	lordIndex   int32
}

func (rule *CardRule) DealCards() {
	CreateCards(rule.CardCount)
	Shuffle()
	AllocCards(rule.PlayerCount, rule.EveryFarmerCardsCount)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (rule *PlayRule) SetLord(i int32) {
	rule.lordIndex = i
}
func (rule *PlayRule) IsLord() bool {
	return rule.lordIndex == rule.playerIndex
}
func (rule *PlayRule) GetCards()[]*message.CardInfo {
	 return everyPlayCardsInfo[rule.playerIndex].CardsInfo
}
func (rule *PlayRule) Play(card []*message.CardInfo,cardsType message.CardsType) bool {
	for  i,v :=range rule.GetCards(){
		if reflect.DeepEqual(v,card){
			return true
		}
	}
	return false
}
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var SingleCardThreePlayer = CardRule{CardCount: 1, PlayerCount: 3, EveryFarmerCardsCount: 17}
var DoubleCardFourPlayer = CardRule{CardCount: 2, PlayerCount: 4, EveryFarmerCardsCount: 26}
var DouDiZhuSingleCardRule = PlayRule{CardRule: SingleCardThreePlayer}
var DouDiZhuDoubleCardRule = PlayRule{CardRule: DoubleCardFourPlayer}
