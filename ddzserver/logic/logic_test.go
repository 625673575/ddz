package logic

import (
	"testing"
	"fmt"
)

func TestDeadCard(t *testing.T) {
	DealCards()
	for k,v :=range everyPlayCardsInfo {
		fmt.Println(k,v)
	}
fmt.Println(extraLordCardsInfo)
}
func TestGetRandomLord(t *testing.T) {
	t.Log(GetRandomLord(3))
}
