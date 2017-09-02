package logic

import (
	"ddzserver/message"
	"math/rand"
	"time"
)

const SINGLE_CARDS_COUNT = 54
const DDOUBLE_CARDS_COUNT = SINGLE_CARDS_COUNT * 2

var library []message.CardInfo
var everyPlayCardsInfo map[int32]message.PlayCardInfo //每个玩家手中的牌
var extraLordCardsInfo message.PlayCardInfo
/// 发牌
func DealTo() {
}

//一开始发牌
func DealCards() {
	CreateCards(1)
	Shuffle()
	AllocCards(3, 17)
}
func Deal() {

}
func shuffle(array []message.CardInfo, source rand.Source) {
	random := rand.New(source)
	for i := len(array) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}
func Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())

	shuffle(library, source) // [c b a]
}
func CreateCards(count int32) {
	for i := int32(0); i < count; i++ {
		for color := 0; color < 4; color++ {
			for value := 0; value < 13; value++ {
				w := (message.Weight )(value)
				s := (message.Suits)(color)
				card := message.CardInfo{CardIndex: i, Weight: w, Color: s}
				library = append(library, card)
			}
		}

		//创建大小joker
		smallJoker := message.CardInfo{CardIndex: i, Weight: message.Weight_SJoker, Color: message.Suits_NoneSuits}
		largeJoker := message.CardInfo{CardIndex: i, Weight: message.Weight_LJoker, Color: message.Suits_NoneSuits}
		library = append(library, smallJoker)
		library = append(library, largeJoker)
	}
}

//正常分配到每个玩家手中牌的策略
func AllocCards(playerCount int32, everyPlayerCardsCount int32) {
	everyPlayCardsInfo = make(map[int32]message.PlayCardInfo)

	for j := int32(0); j < playerCount; j++ {
		info := make([]*message.CardInfo, everyPlayerCardsCount)
		for k := int32(0); k < everyPlayerCardsCount; k++ {
			index := j*everyPlayerCardsCount + k
			v := library[index]
			info[k] = &(v)
		}
		everyPlayCardsInfo[j] = message.PlayCardInfo{CardsInfo: info, User: j}
	}

	extraStartIndex := playerCount * everyPlayerCardsCount
	libraryLen := int32(len(library))

	extraLordCardsInfo = message.PlayCardInfo{}
	extraLordCardsInfo.CardsInfo = make([]*message.CardInfo, libraryLen-extraStartIndex)
	for i := int32(0); i < libraryLen-extraStartIndex; i++ {
		index := extraStartIndex + i
		v := library[index]
		extraLordCardsInfo.CardsInfo[i] = &v
	}
	library = nil
}

func GetRandomLord(playerCount int32) int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31n(playerCount)
}
